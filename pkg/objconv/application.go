package objconv

import (
	"reflect"

	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
)

func FormatApplicationDomain(in interface{}) *admission.Application {
	value := reflect.ValueOf(in).Elem()
	return &admission.Application{
		Identification: admission.Identification{
			UUID: value.FieldByName("UUID").String(),
		},
		School:  FormatSchoolDomain(value.FieldByName("School")),
		Student: FormatStudentDomain(value.FieldByName("Student")),
	}
}

func FormatApplicationNetwork(in interface{}) *unicampus_api_admission_v1alpha1.Application {
	value := reflect.ValueOf(in).Elem()
	return &unicampus_api_admission_v1alpha1.Application{
		UUID:    value.FieldByName("UUID").String(),
		School:  FormatSchoolNetwork(value.FieldByName("School")),
		Student: FormatStudentNetwork(value.FieldByName("Student")),
	}
}
