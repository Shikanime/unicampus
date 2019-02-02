package postgres

import (
	postgresDriver "github.com/jinzhu/gorm"
	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func New(conn *services.PostgreSQLService) *Repository {
	return &Repository{
		conn: conn.Driver(),
	}
}

type Repository struct {
	conn *postgresDriver.DB
}

func (r *Repository) Init() error {
	if err := r.conn.AutoMigrate(
		&School{},
		&Sector{},
		&Location{},
		&GeoPoint{},
		&Region{},
		&Link{},
	).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetSchool(school *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {
	data := new(School)
	if err := r.conn.
		Preload("Links").
		Preload("Pictures").
		Preload("Location.GeoPoint").
		Preload("Location.Region").
		Preload("Location.Sector").
		Take(data, formatSchoolPostgres(school)).Error; err != nil {
		return nil, err
	}

	return formatSchoolDomain(data), nil
}

func (r *Repository) ListSchools(schools []*unicampus_api_education_v1alpha1.School) ([]*unicampus_api_education_v1alpha1.School, error) {
	req := make([]*School, len(schools))
	for _, school := range schools {
		req = append(req, formatSchoolPostgres(school))
	}

	datas := make([]*School, len(req))
	if err := r.conn.
		Preload("Links").
		Preload("Pictures").
		Preload("Location.GeoPoint").
		Preload("Location.Region").
		Preload("Location.Sector").
		Find(&datas, req).Error; err != nil {
		return nil, err
	}

	res := make([]*unicampus_api_education_v1alpha1.School, len(datas))
	for _, data := range datas {
		res = append(res, formatSchoolDomain(data))
	}

	return res, nil
}

func (r *Repository) CreateSchool(school *unicampus_api_education_v1alpha1.School) error {
	if err := r.conn.Create(formatSchoolPostgres(school)).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateSchool(school *unicampus_api_education_v1alpha1.School) error {
	if err := r.conn.Update(formatSchoolPostgres(school)).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteSchool(school *unicampus_api_education_v1alpha1.School) error {
	if err := r.conn.Delete(formatSchoolPostgres(school)).Error; err != nil {
		return err
	}
	return nil
}

func formatSchoolDomain(school *School) *unicampus_api_education_v1alpha1.School {
	var locations []*unicampus_api_education_v1alpha1.Location
	for _, location := range school.Locations {
		locations = append(locations, &unicampus_api_education_v1alpha1.Location{
			Address: location.Address,
			GeoPoint: &unicampus_api_education_v1alpha1.GeoPoint{
				Latitude:  location.GeoPoint.Latitude,
				Longitude: location.GeoPoint.Longitude,
			},
		})
	}

	return &unicampus_api_education_v1alpha1.School{
		UUID:        school.UUID,
		Name:        school.Name,
		Description: school.Description,
		Phone:       school.Phone,
		Email:       school.Email,
		Locations:   locations,
	}
}

func formatSchoolPostgres(in *unicampus_api_education_v1alpha1.School) *School {
	var locations []Location
	for _, location := range in.Locations {
		locations = append(locations, Location{
			Address: location.Address,
			GeoPoint: GeoPoint{
				Latitude:  location.GeoPoint.Latitude,
				Longitude: location.GeoPoint.Longitude,
			},
		})
	}

	return &School{
		UUID:        in.UUID,
		Name:        in.Name,
		Description: in.Description,
		Phone:       in.Phone,
		Email:       in.Email,
		Locations:   locations,
	}
}
