package accesscode

import (
	"context"
	"fmt"
	"github.com/hobbyfarm/gargantua/v4/pkg/apis/hobbyfarm.io/v4alpha1"
	labels2 "github.com/hobbyfarm/gargantua/v4/pkg/labels"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"log/slog"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (acc *accessCodeController) ensureRole(ctx context.Context, obj runtime.Object) (runtime.Object, error) {
	var courses, scenarios, scheduledEvents, machineSets []string
	var labelSelector, objName string
	var isOtac bool

	switch a := obj.(type) {
	case *v4alpha1.AccessCode:
		courses = a.Spec.Courses
		scenarios = a.Spec.Scenarios
		scheduledEvents = a.Spec.ScheduledEvents
		machineSets = a.Spec.MachineSets
		isOtac = false
		objName = a.GetName()
		labelSelector = labels2.AccessCodeLabel
	case *v4alpha1.OneTimeAccessCode:
		courses = a.Spec.Courses
		scenarios = a.Spec.Scenarios
		scheduledEvents = a.Spec.ScheduledEvents
		machineSets = a.Spec.MachineSets
		isOtac = true
		objName = a.GetName()
		labelSelector = labels2.OneTimeAccessCodeLabel
	}

	// retrieve role corresponding with (ot)ac

	selector := labels.SelectorFromSet(map[string]string{
		labelSelector: objName,
	})
	roleList := &v4alpha1.RoleList{}
	if err := acc.roleClient.List(ctx, roleList, &client.ListOptions{
		LabelSelector: selector,
	}); err != nil {
		return nil, err
	}

	var role *v4alpha1.Role
	// if there isn't a role, create it
	if len(roleList.Items) == 0 {
		slog.Debug("role does not exist for (ot)ac, creating it", "kind", obj.GetObjectKind().GroupVersionKind().Kind,
			"objectName", objName)
		role = &v4alpha1.Role{
			ObjectMeta: metav1.ObjectMeta{
				GenerateName: "coderole-",
				Labels: map[string]string{
					labelSelector: objName,
				},
			},
		}

		if err := acc.roleClient.Create(ctx, role); err != nil {
			return nil, err
		}
	} else if len(roleList.Items) == 1 {
		role = &roleList.Items[0]
	} else {
		// there should not be more than one role for a code
		return nil, fmt.Errorf("more than one role exists for %s %s", obj.GetObjectKind().GroupVersionKind().Kind, objName)
	}

	role.Rules = []v4alpha1.Rule{
		makeRule("scenarios", scenarios),
		makeRule("courses", courses),
		makeRule("scheduledEvents", scheduledEvents),
		makeRule("machineSets", machineSets),
	}

	if err := acc.roleClient.Update(ctx, role); err != nil {
		return nil, err
	}

	return role, nil
}

func makeRule(resources string, resourceNames []string) v4alpha1.Rule {
	return v4alpha1.Rule{
		APIGroups: []string{"hobbyfarm.io"},
		Resources: []string{resources},
		ResourceNames: resourceNames
	}
}