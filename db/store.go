package db

import (
	"fmt"
	"github.com/matt9mg/rawflix-api/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StoreConnector interface {
	Connect() (*gorm.DB, error)
}

type Store struct {
	config *StoreConfig
}

type StoreConfig struct {
	Username string
	Password string
	Name     string
	Host     string
	Port     string
	SSLMode  string
	TimeZone string
}

func NewStore(config *StoreConfig) StoreConnector {
	return &Store{
		config: config,
	}
}

func (s *Store) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		s.config.Host,
		s.config.Username,
		s.config.Password,
		s.config.Name,
		s.config.Port,
		s.config.SSLMode,
		s.config.TimeZone,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
