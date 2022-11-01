package data

import (
	"database/sql"
	"errors"
)

type Service struct {
	ServiceId    int64   `json:"service_id"`
	ServiceName  string  `json:"service_name"`
	ServicePrice float64 `json:"service_price"`
}
type ServiceModel struct {
	DB *sql.DB
}

func (s ServiceModel) Get(id int64) (*Service, error) {
	if id < 1 {
		return nil, errors.New("incorrect id")
	}

	query := `
			SELECT *
			FROM service
			WHERE service_id = $1
		`

	var service Service
	err := s.DB.QueryRow(query, id).Scan(
		&service.ServiceId,
		&service.ServiceName,
		&service.ServicePrice,
	)

	if err != nil {
		return nil, errors.New("service not found")
	}

	return &service, nil
}

func (s ServiceModel) Create(serviceName string, servicePrice float64) {
	query := `
		INSERT INTO service (service_name, service_price)
		VALUES ($1, $2)`

	s.DB.QueryRow(query, serviceName, servicePrice).Scan()
}
