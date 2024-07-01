package user

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        string     `json:"id" `
	Name      string     `json:"name" binding:"required"`
	Username  string     `json:"username" binding:"required"`
	Gender    string     `json:"gender" binding:"required"`
	Age       int        `json:"age" binding:"required"`
	Phone     string     `json:"phone" binding:"required"`
	Email     *string    `json:"email" binding:"required"`
	Address   string     `json:"address" binding:"required"`
	Password  *string    `json:"-" binding:"required"`
	Role      string     `json:"role"`
	Activity  *string    `json:"activity"`
	CreatedAt time.Time  `json:"created_at" `
	UpdatedAt time.Time  `json:"-" `
	DeletedAt *time.Time `json:"-" `
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	Name     string  `json:"name" binding:"required"`
	Gender   string  `json:"gender"`
	Age      string  `json:"age"`
	Phone    string  `json:"phone"`
	Email    *string `json:"email"`
	Username string  `json:"username"`
	Address  string  `json:"address"`
	Password *string `json:"password"`
	Activity string  `json:"activity" binding:"required"`
}

func CreateUserToUser(u CreateUser) (res *User, err error) {
	user := User{}
	user.Name = u.Name
	user.Gender = u.Gender

	user.Age, err = strconv.Atoi(u.Age)
	if err != nil {
		return nil, err
	}

	user.Phone = u.Phone
	user.Username = u.Username
	user.Address = u.Address
	user.Activity = &u.Activity
	user.Email = u.Email
	user.Password = u.Password
	return &user, nil
}

type Repository interface {
	CreateUser(c *gin.Context, user *User) (*User, error)
	GetUserByEmail(c *gin.Context, email string) (*User, error)
	DeleteByID(c *gin.Context, id string) error
	Show(c *gin.Context, id string) (*User, error)
	UpdateByAdmin(c *gin.Context, id string, user *User) (*User, error)
}

type Service interface {
	Register(c *gin.Context, user *CreateUser) (*User, error)
	Login(c *gin.Context, user *Login) (*User, error)
	GetUserByEmail(c *gin.Context, email string) (*User, error)
	Show(c *gin.Context, id string) (*User, error)
	Presence(c *gin.Context, user *CreateUser) (*User, error)
	DeleteByID(c *gin.Context, id string) error
	CreateRanger(c *gin.Context, user *CreateUser) (*User, error)
	UpdateByAdmin(c *gin.Context, id string, user *CreateUser) (*User, error)
}

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	UpdateByAdmin(c *gin.Context)
	UpdateAuth(c *gin.Context)
	Show(c *gin.Context)
}
