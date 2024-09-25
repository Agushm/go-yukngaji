package order

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) Create(c *gin.Context, order *Order) (*Order, error) {
	err := r.db.Create(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *repository) Show(c *gin.Context, id string) (*Order, error) {
	order := &Order{}
	tx := r.db
	query := tx.Model(&order)
	err := query.Preload("Event").Preload("PaymentMethod").Where("id = ?", id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *repository) ShowByPublicID(c *gin.Context, public_id string, user_id *string) (*Order, error) {
	order := &Order{}
	tx := r.db
	query := tx.Model(&order)
	if user_id != nil {
		query.Where("user_id = ?", user_id)
	}
	err := query.Preload("Event").Preload("PaymentMethod").Where("public_id = ?", public_id).First(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *repository) Index(c *gin.Context) ([]*Order, error) {
	var order []*Order
	err := r.db.Preload("Event").Preload("PaymentMethod").Find(&order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (r *repository) Update(c *gin.Context, order *Order) (*Order, error) {
	err := r.db.Where("id = ?", order.ID).Save(order).Error
	if err != nil {
		return nil, err
	}
	return order, nil
}