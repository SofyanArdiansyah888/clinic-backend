package user

import (
    "backend/models"
    "gorm.io/gorm"
)

type IUserRepository interface {
    FindAll() ([]models.User, error)
    FindByID(id uint) (*models.User, error)
    Create(user *models.User) error
    Update(user *models.User) error
    Delete(id uint) error
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
    return &userRepository{db: db}
}

func (r *userRepository) FindAll() ([]models.User, error) {
    var users []models.User
    err := r.db.Find(&users).Error
    return users, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
    var user models.User
    err := r.db.First(&user, id).Error
    return &user, err
}

func (r *userRepository) Create(user *models.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) Update(user *models.User) error {
    return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
    return r.db.Delete(&models.User{}, id).Error
}