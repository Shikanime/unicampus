package objconv

import (
	"reflect"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
)

func FormatStudentDomain(in interface{}) *admission.Student {
	value := reflect.ValueOf(in).Elem()
	return &admission.Student{
		Identification: admission.Identification{
			UUID: value.FieldByName("UUID").String(),
		},
		FirstName: value.FieldByName("FirstName").String(),
		LastName:  value.FieldByName("LastName").String(),
	}
}

func FormatStudentNetwork(in interface{}) *unicampus_api_admission_v1alpha1.Student {
	value := reflect.ValueOf(in)
	return &unicampus_api_admission_v1alpha1.Student{
		UUID:      value.FieldByName("UUID").String(),
		FirstName: value.FieldByName("FirstName").String(),
		LastName:  value.FieldByName("LastName").String(),
	}
}
