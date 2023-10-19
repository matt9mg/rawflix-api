package repositories

import (
	"github.com/matt9mg/rawflix-api/entities"
	"gorm.io/gorm"
)

type InteractionRepository interface {
	CreateInBatches(batch int, interactions ...*entities.Interaction) error
	FindAllNotSyncedWithRecombee() ([]*entities.Interaction, error)
	MarkAsSynced(interaction *entities.Interaction) error
	Create(interaction *entities.Interaction) error
}

type Interaction struct {
	db *gorm.DB
}

func NewInteraction(db *gorm.DB) InteractionRepository {
	return &Interaction{
		db: db,
	}
}

func (i *Interaction) CreateInBatches(batch int, interactions ...*entities.Interaction) error {
	return i.db.CreateInBatches(interactions, batch).Error
}

func (i *Interaction) FindAllNotSyncedWithRecombee() ([]*entities.Interaction, error) {
	var interactions []*entities.Interaction

	err := i.db.Model(&entities.Interaction{}).Where("recombee = false").Scan(&interactions).Error

	return interactions, err
}

func (i *Interaction) MarkAsSynced(interaction *entities.Interaction) error {
	return i.db.Model(interaction).Update("recombee", true).Error
}

func (i *Interaction) Create(interaction *entities.Interaction) error {
	return i.db.Create(interaction).Error
}
