package repository

import (
	"errors"
	"kupon/features/coupon/repository"
	"kupon/features/users"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Coupons  []repository.Coupon `gorm:"foreignKey:UserID"`
}

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &userQuery{
		db: db,
	}
}

// Register implements users.Repository.
func (r *userQuery) Register(newUser users.User) (users.User, error) {
	// panic("unimplemented")
	inputUser := new(User)
	inputUser.Name = newUser.Name
	inputUser.Email = newUser.Email
	inputUser.Password = newUser.Password

	tx := r.db.Create(&inputUser)
	if tx.Error != nil {
		return users.User{}, tx.Error
	}
	newUser.ID = inputUser.ID

	return newUser, nil
}

// Login implements users.Repository.
func (r *userQuery) Login(email string) (users.User, error) {
	// panic("unimplemented")
	dataLogin := new(User)

	if err := r.db.Where("email = ?", email).First(dataLogin).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return users.User{}, errors.New("email tidak ditemukan")
		}
		return users.User{}, err
	}
	result := new(users.User)
	result.ID = dataLogin.ID
	result.Name = dataLogin.Name
	result.Email = dataLogin.Email
	result.Password = dataLogin.Password

	return *result, nil

}

// Update implements users.Repository.
func (r *userQuery) Update(id uint, updateUser users.User) (users.User, error) {
	// panic("unimplemented")
	var updateData User
	tx := r.db.First(&updateData, id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return users.User{}, errors.New("User id not found")
		}
		return users.User{}, tx.Error
	}
	updateData.Name = updateUser.Name
	updateData.Email = updateUser.Email
	// updateData.Password = updateUser.Password

	updateTx := r.db.Model(&User{}).Where("id=?", id).Updates(updateData)
	if updateTx.Error != nil {
		return users.User{}, updateTx.Error
	}
	return updateUser, nil
}

// GetUser implements users.Repository.
func (r *userQuery) GetUser() ([]users.User, error) {
	// panic("unimplemented")
	var userData []User
	tx := r.db.Find(&userData)
	if tx.Error != nil {
		return []users.User{}, tx.Error
	}
	var userList []users.User
	for _, user := range userData {
		userList = append(userList, users.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}
	return userList, nil
}
