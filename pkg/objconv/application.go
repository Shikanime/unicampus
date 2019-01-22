package objconv

import (
	"reflect"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/pkg/admission"
)

func FormatApplicationDomain(in interface{}) *admission.Application {
	inValue := reflect.ValueOf(in)
	return &admission.Application{
		UUID:    inValue.Elem().FieldByName("UUID").String(),
		School:  FormatSchoolDomain(inValue.Elem().FieldByName("School")),
		Student: FormatStudentDomain(inValue.Elem().FieldByName("Student")),
	}
}

func FormatApplicationNetwork(in interface{}) *unicampus_api_admission_v1alpha1.Application {
	inValue := reflect.ValueOf(in)
	return &unicampus_api_admission_v1alpha1.Application{
		UUID:    inValue.Elem().FieldByName("UUID").String(),
		School:  FormatSchoolNetwork(inValue.Elem().FieldByName("School")),
		Student: FormatStudentNetwork(inValue.Elem().FieldByName("Student")),
	}
}
