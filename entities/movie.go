package entities

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"time"
)

type Movie struct {
	gorm.Model
	Title     string
	Year      string
	Rated     string
	Released  string
	Runtime   string
	Genre     datatypes.JSON
	Directors datatypes.JSON
	Writers   datatypes.JSON
	Actors    datatypes.JSON
	Plot      string
	Language  datatypes.JSON
	Country   datatypes.JSON
	Awards    string
	Poster    string
	Added     time.Time
	Recombee  bool `gorm:"default=false"`
}
