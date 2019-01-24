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
		Region: admission.Region{
			City:    value.FieldByName("City").String(),
			Country: value.FieldByName("Country").String(),
			State:   value.FieldByName("State").String(),
			Zipcode: value.FieldByName("Zipcode").String(),
		},
		Location: admission.Location{
			Address:   value.FieldByName("Addres").String(),
			Latitude:  value.FieldByName("Latitude").Float(),
			Longitude: value.FieldByName("Longitude").Float(),
		},
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
