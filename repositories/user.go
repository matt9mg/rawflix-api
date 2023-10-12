package repositories

import (
	"github.com/matt9mg/rawflix-api/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateInBatches(batch int, users ...*entities.User) error
	UsernameExists(username string) (bool, error)
	Create(user *entities.User) error
	IDAndTokenExists(id uint, token string) (bool, error)
	FindOneByUsername(username string) (*entities.User, error)
	Save(user *entities.User) error
	RemoveTokenFromByUserID(userID uint) error
	FindWhereNotSyncedWithRecombee() ([]*entities.User, error)
	MarkAsSyncedWithRecombee(user *entities.User) error
	FindAll() ([]*entities.User, error)
}

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) UserRepository {
	return &User{
		db: db,
	}
}

func (u *User) CreateInBatches(batch int, users ...*entities.User) error {
	return u.db.CreateInBatches(users, batch).Error
}

func (u *User) UsernameExists(username string) (bool, error) {
	var total int
	err := u.db.Model(&entities.User{}).Select("count(*) as total").Where("username = ?", username).Scan(&total).Error

	return total > 0, err
}

func (u *User) IDAndTokenExists(id uint, token string) (bool, error) {
	var total int
	err := u.db.Model(&entities.User{}).Select("count(*) as total").Where("id = ?", id).Where("token = ?", token).Scan(&total).Error

	return total > 0, err
}

func (u *User) Create(user *entities.User) error {
	return u.db.Create(user).Error
}

func (u *User) FindOneByUsername(username string) (*entities.User, error) {
	var user *entities.User

	err := u.db.Model(&entities.User{}).Where("username = ?", username).Scan(&user).Error

	return user, err
}

func (u *User) Save(user *entities.User) error {
	return u.db.Save(user).Error
}

func (u *User) RemoveTokenFromByUserID(userID uint) error {
	return u.db.Model(&entities.User{}).Where("id = ?", userID).Update("token", nil).Error
}

func (u *User) FindWhereNotSyncedWithRecombee() ([]*entities.User, error) {
	var users []*entities.User

	err := u.db.Model(&entities.User{}).Where("recombee != ?", true).Scan(&users).Error

	return users, err
}

func (u *User) MarkAsSyncedWithRecombee(user *entities.User) error {
	return u.db.Model(user).Update("recombee", true).Error
}

func (u *User) FindAll() ([]*entities.User, error) {
	var users []*entities.User

	err := u.db.Model(&entities.User{}).Find(&users).Error

	return users, err
}
