package environment

import (
	"context"
	"github.com/golang/glog"
	hfv1controllers "github.com/hobbyfarm/gargantua/pkg/generated/controllers/hobbyfarm.io/v1"
	hfv1types "github.com/hobbyfarm/gargantua/pkg/apis/hobbyfarm.io/v1"
	"github.com/hobbyfarm/gargantua/pkg/generated/informers/externalversions"
	"github.com/hobbyfarm/gargantua/pkg/generated/listers/hobbyfarm.io/v1"
	k8stypedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
	"time"
)

type Handler struct {
	envWorkqueue workqueue.RateLimitingInterface
	//vmWorkqueue  workqueue.RateLimitingInterface

	vmTemplateIndexer cache.Indexer

	vmLister  v1.VirtualMachineLister
	envLister v1.EnvironmentLister

	vmSynced  cache.InformerSynced
	envSynced cache.InformerSynced
}

const (
	controllerAgentName = "environment-controller"
	vmEnvironmentIndex = "vm.vmclaim.controllers.hobbyfarm.io/environment-index"
)

func Register(
	ctx context.Context,
	events k8stypedcorev1.EventInterface,
	environments hfv1controllers.EnvironmentController,
	exver externalversions.SharedInformerFactory,
	) {

		cx := &Handler {
			envWorkqueue:      workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Environment")
			vmTemplateIndexer: exver.Hobbyfarm().V1().VirtualMachineTemplates().Informer().GetIndexer(),
			vmLister: exver.Hobbyfarm().V1().VirtualMachines().Lister(),
			envLister: exver.Hobbyfarm().V1().Environments().Lister(),
			vmSynced: exver.Hobbyfarm().V1().VirtualMachines().Informer().HasSynced,
			envSynced: exver.Hobbyfarm().V1().Environments().Informer().HasSynced,
		}

		envInformer := exver.Hobbyfarm().V1().Environments().Informer()
		vmInformer := exver.Hobbyfarm().V1().VirtualMachines().Informer()

		envInformer.AddEventHandlerWithResyncPeriod(cache.ResourceEventHandlerFuncs{
			AddFunc: cx.handleVM,
			UpdateFunc: func(old, new interface{}) {
				cx.handleVM(new)
			},
			DeleteFunc: cx.handleVM,
		}, time.Second * 30)
}

func (h *Handler) handleVM(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			glog.Errorf("error decoding object, invalid type")
			return
		}

	}
}