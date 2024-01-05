package accesscode

import (
	"context"
	"github.com/acorn-io/mink/pkg/stores"
	"github.com/acorn-io/mink/pkg/strategy"
	v1 "github.com/hobbyfarm/gargantua/v3/pkg/apis/hobbyfarm.io/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
)

type accessCodeValidator struct {
	scenarioGetter strategy.Getter
	coursesGetter  strategy.Getter
}

func NewAccessCodeStorage(
	accessCodeStrategy strategy.CompleteStrategy,
	scenarioGetter strategy.Getter,
	coursesGetter strategy.Getter) (rest.Storage, error) {

	acValidator := accessCodeValidator{
		scenarioGetter: scenarioGetter,
		coursesGetter:  coursesGetter,
	}

	return stores.NewBuilder(accessCodeStrategy.Scheme(), &v1.AccessCode{}).
		WithCompleteCRUD(accessCodeStrategy).
		WithValidateCreate(acValidator).
		Build(), nil
}

func (av accessCodeValidator) Validate(ctx context.Context, obj runtime.Object) (result field.ErrorList) {
	result = append(result, av.validateScenarios(ctx, obj)...)
	result = append(result, av.validateCourses(ctx, obj)...)

	return
}

func (av accessCodeValidator) validateScenarios(ctx context.Context, obj runtime.Object) (result field.ErrorList) {
	accessCode := obj.(*v1.AccessCode)

	for _, s := range accessCode.Spec.Scenarios {
		if _, err := av.scenarioGetter.Get(ctx, accessCode.GetNamespace(), s); err != nil {
			result = append(result, field.Invalid(field.NewPath("spec", "scenarios"), s, err.Error()))
			return
		}
	}

	return
}

func (av accessCodeValidator) validateCourses(ctx context.Context, obj runtime.Object) (result field.ErrorList) {
	accessCode := obj.(*v1.AccessCode)

	for _, c := range accessCode.Spec.Courses {
		if _, err := av.coursesGetter.Get(ctx, accessCode.GetNamespace(), c); err != nil {
			result = append(result, field.Invalid(field.NewPath("spec", "courses"), c, err.Error()))
			return
		}
	}

	return
}
