package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type UserRepository struct {
	baseRepository
}

func NewUserRepository(baseRepository baseRepository) *UserRepository {
	return &UserRepository{baseRepository: baseRepository}
}

func (up *UserRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := up.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (up *UserRepository) Create(user *models.User) error {
	user.ID = uuid.New().String()
	err := up.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (up *UserRepository) Update(user *models.User) (*models.User, error) {
	err := up.db.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
