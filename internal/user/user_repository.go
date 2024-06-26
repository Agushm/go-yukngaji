package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) CreateUser(c *gin.Context, user *User) (*User, error) {
	err := r.db.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetUserByEmail(c *gin.Context, email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).Where("deleted_at IS NULL").First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) DeleteByID(c *gin.Context, id string) error {
	var user User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}
	var now = time.Now()
	user.DeletedAt = &now
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *repository) Show(c *gin.Context, id string) (*User, error) {
	user := &User{}
	err := r.db.Where("id = ?", id).Where("deleted_at IS NULL").First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) UpdateByAdmin(c *gin.Context, id string, user *User) (*User, error) {
	err := r.db.Where("id = ?", id).Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
