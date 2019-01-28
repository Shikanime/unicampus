package postgres

import (
	"github.com/jinzhu/gorm"
	unicampus_api_education_v1alpha1 "gitlab.com/deva-hub/unicampus/api/education/v1alpha1"
	"gitlab.com/deva-hub/unicampus/internal/pkg/services"
)

func NewPostgresRepository(conn *services.PostgreSQLService) *PostgresRepository {
	return &PostgresRepository{
		conn: conn.Driver(),
	}
}

type PostgresRepository struct {
	conn *gorm.DB
}

func (r *PostgresRepository) GetSchool(school *unicampus_api_education_v1alpha1.School) (*unicampus_api_education_v1alpha1.School, error) {
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

	return formatSchoolGRPC(data), nil
}

func (r *PostgresRepository) ListSchools(schools []*unicampus_api_education_v1alpha1.School) ([]*unicampus_api_education_v1alpha1.School, error) {
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
		res = append(res, formatSchoolGRPC(data))
	}

	return res, nil
}

func (r *PostgresRepository) CreateSchool(school *unicampus_api_education_v1alpha1.School) error {
	if err := r.conn.Create(formatSchoolPostgres(school)).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) UpdateSchool(school *unicampus_api_education_v1alpha1.School) error {
	if err := r.conn.Update(formatSchoolPostgres(school)).Error; err != nil {
		return err
	}
	return nil
}

func (r *PostgresRepository) DeleteSchool(school *unicampus_api_education_v1alpha1.School) error {
	if err := r.conn.Delete(formatSchoolPostgres(school)).Error; err != nil {
		return err
	}
	return nil
}

func formatSchoolGRPC(school *School) *unicampus_api_education_v1alpha1.School {
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
	return &School{
		UUID:        in.UUID,
		Name:        in.Name,
		Description: in.Description,
		Phone:       in.Phone,
		Email:       in.Email,
	}
}
