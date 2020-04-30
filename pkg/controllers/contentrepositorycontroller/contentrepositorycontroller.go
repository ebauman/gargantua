package contentrepositorycontroller

import (
	"crypto/sha256"
	"encoding/base32"
	"fmt"
	"github.com/golang/glog"
	contentRepov1 "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	hfClientset "github.com/hobbyfarm/gargantua/pkg/client/clientset/versioned"
	hfInformers "github.com/hobbyfarm/gargantua/pkg/client/informers/externalversions"
	hfListers "github.com/hobbyfarm/gargantua/pkg/client/listers/hobbyfarm.io/v1"
	"github.com/rancher/gitwatcher/pkg/apis/gitwatcher.cattle.io/v1"
	gwClientset "github.com/rancher/gitwatcher/pkg/generated/clientset/versioned"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	"sort"
	"strings"
	"time"
)

const (
	ContentRepositoryLabel = "label.hobbyfarm.io/contentrepository"
	ContentRepositoryVersionLabel = "label.hobbyfarm.io/contentrepositoryversion"
	Namespace = "hobbyfarm" // probably variable this out somehow in the future
)

type ContentRepositoryController struct {
	hfClientset *hfClientset.Clientset
	gwClientset *gwClientset.Clientset

	crWorkqueue workqueue.RateLimitingInterface

	crSynced cache.InformerSynced

	crLister hfListers.ContentRepositoryLister
}

func NewContentRepositoryController(hfClientSet *hfClientset.Clientset,
	hfInformerFactory hfInformers.SharedInformerFactory,
	gwClientSet *gwClientset.Clientset) (*ContentRepositoryController, error) {
	crc := ContentRepositoryController{}

	crc.hfClientset = hfClientSet
	crc.gwClientset = gwClientSet
	crc.crSynced = hfInformerFactory.Hobbyfarm().V1().ContentRepositories().Informer().HasSynced

	crc.crWorkqueue = workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "ContentRepository")

	crc.crLister = hfInformerFactory.Hobbyfarm().V1().ContentRepositories().Lister()
	crInformer := hfInformerFactory.Hobbyfarm().V1().ContentRepositories().Informer()
	crInformer.AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
		AddFunc: crc.handleCR,
		UpdateFunc: func(old, new interface{}) {
			crc.handleCR(new)
		},
		DeleteFunc: crc.handleCR,
	}, time.Second*30)

	return &crc, nil
}

func (crc *ContentRepositoryController) Run(stopCh <-chan struct{}) error {
	defer crc.crWorkqueue.ShutDown()

	glog.V(4).Infof("starting contentrepository controller")
	glog.Info("waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, crc.crSynced); !ok {
		return fmt.Errorf("failed to wait for contentrepository cache to sync")
	}
	glog.Info("starting contentrepository controller workers")
	go wait.Until(crc.runCrWorker, time.Second, stopCh)
	glog.Info("started contentrepository controller workers")
	<-stopCh
	return nil
}

func (crc *ContentRepositoryController) enqueueCR(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		return
	}

	glog.V(8).Infof("Enqueing ContentRepository %s", key)
	crc.crWorkqueue.AddRateLimited(key)
}

func (crc *ContentRepositoryController) runCrWorker() {
	glog.V(6).Infof("Starting ContentRepository worker")
	for crc.processNextRepository() {

	}
}

func (crc *ContentRepositoryController) handleCR(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		// DeletedFinalStateUnknown occurs when an object has been deleted
		// while our reconciliation loop was not running or not receiving object updates.
		// Thus the object may have been gone for some time, but the controller may still need to
		// process the deletion. So DeleteFinalStateUnknown has a tombstone of the deleted object
		// that can be used to reconstruct the deleted object.
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			// this error occurs if the tombstone is not able to be decoded
			glog.Errorf("Error decoding object, invalid type")
			return
		}

		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			// this error occurs if the tombstone has been decoded, but can't
			// be turned into a Kubernetes object (k8s.io/apimachinery/pkg/apis/meta/v1 Object)
			glog.Errorf("Error decoding object tombstone, invalid type")
			return
		}
		// at this point, we have successfully recovered a metav1.Object. this does _not_
		// mean it is a ContentRepository, though. Just that the object is in fact a
		// Kubernetes metav1.Object.
		klog.V(4).Infof("Recovered deleted object %s from tombstone", object.GetName())
	}
	klog.V(4).Infof("Processing object %s", object.GetName())
	crc.enqueueCR(object)
}

func (crc *ContentRepositoryController) processNextRepository() bool {
	obj, shutdown := crc.crWorkqueue.Get()
	if shutdown {
		return false
	}

	err := func() error {
		defer crc.crWorkqueue.Done(obj)

		glog.V(4).Infof("Processing ContentRepository %v", obj)
		_, objName, err := cache.SplitMetaNamespaceKey(obj.(string))
		if err != nil {
			glog.Errorf("Error while splitting meta namespace key for ContentRepository %v", err)
			return nil
		}

		err = crc.reconcileContentRepository(objName)

		if err != nil {
			glog.Error(err)
		}
		crc.crWorkqueue.Forget(obj)
		glog.V(4).Infof("ContentRepository %v processed by controller", objName)

		return nil
	}()

	if err != nil {
		return true
	}

	return true
}

// if a CR change occurs, something is either being created, updated, or deleted
// we need to ensure that in the event of creation or update, the corresponding
// GitWatcher is updated or created. In the event of deletion, make sure we remove
// the GitWatcher & related GitCommits (if GW doesn't handle that), but leave
// the Scenarios and Courses alone.
func (crc *ContentRepositoryController) reconcileContentRepository(crId string) error {
	namespace, name, err := cache.SplitMetaNamespaceKey(crId)
	if err != nil {
		return fmt.Errorf("invalid resource key: %s", crId)
	}


	contentRepo, err := crc.crLister.Get(crId)
	if err != nil {
		// This probably means that the CR was deleted.
		if errors.IsNotFound(err) {
			// This definitely means that the CR was deleted.
			// In that case, we need to remove the corresponding GitWatcher.
			err = crc.processRemoval(namespace, name)

			if err != nil {
				return fmt.Errorf("error while removing ContentRepository: %v", err)
			}
			return nil // we have handled the deletion, so return nil
		}

		return err // failing the errors.IsNotFound check means that this isn't an "object doesn't exist" error, so return it
	}

	gwList, err := crc.gwClientset.GitwatcherV1().GitWatchers(Namespace).List(metav1.ListOptions{
		LabelSelector: fmt.Sprintf("%s=%s", ContentRepositoryLabel, crId), // this sucks, you _should_ be able to query for ownership but you can't
			})
	if err != nil {
		return fmt.Errorf("error retrieving GitWatchers for ContentRepository %s: %v", crId, err)
	}

	if len(gwList.Items) < 1 {
		// there are no gitwatchers associated with this contentrepo, so create one
		gw := v1.GitWatcher{}
		setVals(&gw, contentRepo)

		_, err = crc.gwClientset.GitwatcherV1().GitWatchers(Namespace).Create(&gw)
		if err != nil {
			return fmt.Errorf("error while creating GitWatcher: %v", err)
		}
	}

	if len(gwList.Items) == 1 {
		gw := gwList.Items[0]

		if contentRepo.ResourceVersion == gw.Labels[ContentRepositoryVersionLabel] {
			return nil
		}

		// the GitWatcher is out of sync with the latest version of the ContentRepository
		setVals(&gw, contentRepo)

		_, err := crc.gwClientset.GitwatcherV1().GitWatchers(Namespace).Update(&gw)
		if err != nil {
			return fmt.Errorf("error while updating GitWatcher: %v", err)
		}
	}

	if len(gwList.Items) > 1 {
		// there are too many gitwatchers. remove some.
		// deterministic: always keep the oldest one as it was presumably first
		// is this reasonable logic? I don't know but we're doing it
		sort.Slice(gwList.Items, func(i, j int) bool {
			return gwList.Items[i].CreationTimestamp.Unix() < gwList.Items[j].CreationTimestamp.Unix()
		})

		for _, gw := range gwList.Items[1:] {
			err = crc.gwClientset.GitwatcherV1().GitWatchers(Namespace).Delete(gw.Name, &metav1.DeleteOptions{})
			if err != nil {
				return fmt.Errorf("error while removing superfluous GitWatchers: %v", err)
			}
		}

		return nil
	}

	return nil
}

func (crc *ContentRepositoryController) processRemoval(namespace string, name string) error {
	// To remove a ContentRepository, we need to just remove all the
	// gitwatchers that are owned

	label := fmt.Sprintf("%s=%s", ContentRepositoryLabel, name)

	// first, look up all the gitwatchers
	gwList, err := crc.gwClientset.GitwatcherV1().GitWatchers(namespace).List(metav1.ListOptions{LabelSelector: label})
	if err != nil {
		return fmt.Errorf("error listing GitWatchers: %v", err)
	}

	if len(gwList.Items) == 0 {
		return nil // not deleting things that don't exist
	}

	// if we get to this point, there are gitwatchers that exist with the label selector used above
	// we need to delete them.
	// We specify metav1.DeletePropagationBackground to tell Kubernetes to garbage collect the GitCommit
	// objects that would otherwise be potentially orphaned when the GitWatcher is deleted
	deletionPolicy := metav1.DeletePropagationBackground
	for _, gw := range gwList.Items {
		err = crc.gwClientset.GitwatcherV1().GitWatchers(namespace).Delete(gw.Name, &metav1.DeleteOptions{PropagationPolicy: &deletionPolicy})
		if err != nil {
			return fmt.Errorf("error deleting GitWatcher %s: %v", gw.Name, err)
		}
	}

	// If all gitwatchers have been deleted, return happily!
	return nil
}

func makeName() string {
	hasher := sha256.New()
	hasher.Write([]byte(time.Now().String())) // gross
	sha := base32.StdEncoding.WithPadding(-1).EncodeToString(hasher.Sum(nil))[:8]
	return strings.ToLower(sha)
}

func setVals(gw *v1.GitWatcher, cr *contentRepov1.ContentRepository) {
	gw.Name = fmt.Sprintf("gw-%s-cr-%s", makeName(), cr.Name)
	gw.Spec.RepositoryURL = cr.Spec.RepositoryURL
	gw.Spec.Enabled = cr.Spec.Enabled
	gw.Spec.Branch = cr.Spec.Branch
	gw.Spec.Provider = "polling" // we should implement webhook, etc. in the future
	var tru = true // hack to get ptr to a bool
	ownerRef := metav1.OwnerReference{
		APIVersion: cr.APIVersion,
		Kind: cr.Kind,
		Name: cr.Name,
		Controller: &tru, // yuck
		UID: cr.UID,
	}
	labels := map[string]string{
		ContentRepositoryLabel: cr.Name,
		ContentRepositoryVersionLabel: cr.ResourceVersion,
	}

	gw.SetLabels(labels)
	gw.SetOwnerReferences([]metav1.OwnerReference{ownerRef})
}