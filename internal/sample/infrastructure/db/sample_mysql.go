package db

import (
	"database/sql"
	"fmt"
	"gin-api-sample/internal/sample/domain/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type SampleRepository interface {
	GetSamples() ([]model.Sample, error)
	CreateSample(sample *model.Sample) error
	UpdateSample(sample *model.Sample) error
	DeleteSample(id string) error
}

type sampleRepository struct {
	db *sql.DB
}

func NewSampleRepository() SampleRepository {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return &sampleRepository{db: db}
}

func (r *sampleRepository) GetSamples() ([]model.Sample, error) {
	rows, err := r.db.Query("SELECT id, name FROM samples")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var samples []model.Sample
	for rows.Next() {
		var sample model.Sample
		if err := rows.Scan(&sample.ID, &sample.Name); err != nil {
			return nil, err
		}
		samples = append(samples, sample)
	}
	return samples, nil
}

func (r *sampleRepository) CreateSample(sample *model.Sample) error {
	_, err := r.db.Exec("INSERT INTO samples (name) VALUES (?)", sample.Name)
	return err
}

func (r *sampleRepository) UpdateSample(sample *model.Sample) error {
	_, err := r.db.Exec("UPDATE samples SET name = ? WHERE id = ?", sample.Name, sample.ID)
	return err
}

func (r *sampleRepository) DeleteSample(id string) error {
	_, err := r.db.Exec("DELETE FROM samples WHERE id = ?", id)
	return err
}
