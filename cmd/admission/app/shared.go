package app

import (
	unicampus_api_admission_v1alpha1 "gitlab.com/deva-hub/unicampus/api/admission/v1alpha1"
	"gitlab.com/deva-hub/unicampus/pkg/admission"
)

func formatRegionDomain(in unicampus_api_admission_v1alpha1.Region) *admission.Region {
	return &admission.Region{
		City:    in.City,
		Country: in.Country,
		State:   in.State,
		Zipcode: in.Zipcode,
	}
}

func formatLocationDomain(in unicampus_api_admission_v1alpha1.Location) *admission.Location {
	return &admission.Location{
		Address:   in.Address,
		Latitude:  in.Latitude,
		Longitude: in.Longitude,
	}
}
