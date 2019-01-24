package objconv

import (
	"reflect"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
)

func FormatSchoolDomain(in interface{}) *admission.School {
	inValue := reflect.ValueOf(in)
	return &admission.School{
		UUID:        inValue.Elem().FieldByName("UUID").String(),
		Name:        inValue.Elem().FieldByName("Name").String(),
		Description: inValue.Elem().FieldByName("Description").String(),
	}
}

func FormatSchoolNetwork(in interface{}) *unicampus_api_admission_v1alpha1.School {
	inValue := reflect.ValueOf(in)
	return &unicampus_api_admission_v1alpha1.School{
		UUID:        inValue.Elem().FieldByName("UUID").String(),
		Name:        inValue.Elem().FieldByName("Name").String(),
		Description: inValue.Elem().FieldByName("Description").String(),
	}
}
