package repositories

import (
	"github.com/matt9mg/rawflix-api/entities"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type MovieRepository interface {
	CreateInBatches(int, ...*entities.Movie) error
	Save(movies ...*entities.Movie) error
	Truncate() error
	FindAll() ([]*entities.Movie, error)
	FindWhereNotSyncedWithRecombee() ([]*entities.Movie, error)
	MarkAsSyncedWithRecombee(movie *entities.Movie) error
	FindMoviesWithGenre(genre string) ([]*entities.Movie, error)
	GetByRecommendation(ids []uint, userID uint, interactionType entities.InteractionType) ([]map[string]interface{}, error)
	FindByID(id uint) (*entities.Movie, error)
}

type Movie struct {
	db *gorm.DB
}

func NewMovie(db *gorm.DB) MovieRepository {
	return &Movie{
		db: db,
	}
}

func (m *Movie) CreateInBatches(batchSize int, movies ...*entities.Movie) error {
	return m.db.CreateInBatches(movies, batchSize).Error
}

func (m *Movie) Save(movies ...*entities.Movie) error {
	return m.db.Save(movies).Error
}

func (m *Movie) Truncate() error {
	return m.db.Exec("TRUNCATE movies RESTART IDENTITY CASCADE;").Error
}

func (m *Movie) FindAll() ([]*entities.Movie, error) {
	var movies []*entities.Movie

	err := m.db.Model(&entities.Movie{}).Scan(&movies).Error

	return movies, err
}

func (m *Movie) FindWhereNotSyncedWithRecombee() ([]*entities.Movie, error) {
	var movies []*entities.Movie

	err := m.db.Model(&entities.Movie{}).Where("recombee != ?", true).Scan(&movies).Error

	return movies, err
}

func (u *Movie) MarkAsSyncedWithRecombee(movie *entities.Movie) error {
	return u.db.Model(movie).Update("recombee", true).Error
}

func (u *Movie) FindMoviesWithGenre(genre string) ([]*entities.Movie, error) {
	var movies []*entities.Movie

	err := u.db.Model(&entities.Movie{}).Where(datatypes.JSONQuery("genre").HasKey(genre)).Scan(&movies).Error

	return movies, err
}

func (u *Movie) GetByRecommendation(ids []uint, userID uint, interactionType entities.InteractionType) ([]map[string]interface{}, error) {
	var movies []map[string]interface{}

	err := u.db.Raw("select movies.poster, movies.title, movies.id, movies.genre, movies.runtime, movies.plot, interactions.id as interaction_id "+
		"FROM movies left join interactions on movies.id = interactions.movie_id AND interactions.type = ? and interactions.user_id = ? "+
		"WHERE movies.id IN (?)", interactionType, userID, ids).Scan(&movies).Error

	return movies, err
}

func (u *Movie) FindByID(id uint) (*entities.Movie, error) {
	var movie *entities.Movie

	err := u.db.Model(&entities.Movie{}).Find(&movie, id).Error

	return movie, err
}
