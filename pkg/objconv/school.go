package objconv

import (
	"reflect"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
)

func FormatSchoolDomain(in interface{}) *admission.School {
	value := reflect.ValueOf(in).Elem()
	return &admission.School{
		Identification: admission.Identification{
			UUID: value.FieldByName("UUID").String(),
		},
		Name:        value.FieldByName("Name").String(),
		Description: value.FieldByName("Description").String(),
	}
}

func FormatSchoolNetwork(in interface{}) *unicampus_api_admission_v1alpha1.School {
	value := reflect.ValueOf(in)
	return &unicampus_api_admission_v1alpha1.School{
		UUID:        value.FieldByName("UUID").String(),
		Name:        value.FieldByName("Name").String(),
		Description: value.FieldByName("Description").String(),
	}
}
