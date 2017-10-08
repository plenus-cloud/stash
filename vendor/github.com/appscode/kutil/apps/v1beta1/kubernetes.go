package v1beta1

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/appscode/kutil"
	"k8s.io/apimachinery/pkg/runtime/schema"
	apps "k8s.io/client-go/pkg/apis/apps/v1beta1"
)

func GetGroupVersionKind(v interface{}) schema.GroupVersionKind {
	return apps.SchemeGroupVersion.WithKind(kutil.GetKind(v))
}

func AssignTypeKind(v interface{}) error {
	if reflect.ValueOf(v).Kind() != reflect.Ptr {
		return fmt.Errorf("%v must be a pointer", v)
	}

	switch u := v.(type) {
	case *apps.StatefulSet:
		u.APIVersion = apps.SchemeGroupVersion.String()
		u.Kind = kutil.GetKind(v)
		return nil
	case *apps.Deployment:
		u.APIVersion = apps.SchemeGroupVersion.String()
		u.Kind = kutil.GetKind(v)
		return nil
	}
	return errors.New("unknown api object type")
}
