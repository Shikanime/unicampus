package objconv

import (
	"reflect"

	unicampus_api_admission_v1alpha1 "github.com/Shikanime/unicampus/api/admission/v1alpha1"
	"github.com/Shikanime/unicampus/pkg/admission"
)

func FormatStudentDomain(in interface{}) *admission.Student {
	inValue := reflect.ValueOf(in)
	return &admission.Student{
		UUID:      inValue.Elem().FieldByName("UUID").String(),
		FirstName: inValue.Elem().FieldByName("FirstName").String(),
		LastName:  inValue.Elem().FieldByName("LastName").String(),
	}
}

func FormatStudentNetwork(in interface{}) *unicampus_api_admission_v1alpha1.Student {
	inValue := reflect.ValueOf(in)
	return &unicampus_api_admission_v1alpha1.Student{
		UUID:      inValue.Elem().FieldByName("UUID").String(),
		FirstName: inValue.Elem().FieldByName("FirstName").String(),
		LastName:  inValue.Elem().FieldByName("LastName").String(),
	}
}
