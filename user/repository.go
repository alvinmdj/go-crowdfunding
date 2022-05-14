package user

import "gorm.io/gorm"

type Repository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

// User repository instance
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Repository to create a new user
func (r *repository) Create(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Repository to find user by email
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Repository to find user by ID
func (r *repository) FindById(id int) (User, error) {
	var user User

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// Repository to update user
func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
