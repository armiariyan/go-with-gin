package repo

import (
	"intern_golang/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	All(userID string) ([]entities.User, error)
	// InsertProduct(product entity.Product) (entity.Product, error)
	// UpdateProduct(product entity.Product) (entity.Product, error)
	// DeleteProduct(productID string) error
	// FindOneProductByID(ID string) (entity.Product, error)
	// FindAllProduct(userID string) ([]entity.Product, error)
}

type userRepo struct {
	connection *gorm.DB
}

func NewProductRepo(connection *gorm.DB) UserRepository {
	return &userRepo{
		connection: connection,
	}
}

func (c *userRepo) All(userID string) ([]entities.User, error) {
	users := []entities.User{}
	// c.connection.Preload("User").Where("user_id = ?", userID).Find(&users)
	c.connection.Find(&users)
	return users, nil
}

// func (c *productRepo) InsertProduct(product entity.Product) (entity.Product, error) {
// 	c.connection.Save(&product)
// 	c.connection.Preload("User").Find(&product)
// 	return product, nil
// }

// func (c *productRepo) UpdateProduct(product entity.Product) (entity.Product, error) {
// 	c.connection.Save(&product)
// 	c.connection.Preload("User").Find(&product)
// 	return product, nil
// }

// func (c *productRepo) FindOneProductByID(productID string) (entity.Product, error) {
// 	var product entity.Product
// 	res := c.connection.Preload("User").Where("id = ?", productID).Take(&product)
// 	if res.Error != nil {
// 		return product, res.Error
// 	}
// 	return product, nil
// }

// func (c *productRepo) FindAllProduct(userID string) ([]entity.Product, error) {
// 	products := []entity.Product{}
// 	c.connection.Where("user_id = ?", userID).Find(&products)
// 	return products, nil
// }

// func (c *productRepo) DeleteProduct(productID string) error {
// 	var product entity.Product
// 	res := c.connection.Preload("User").Where("id = ?", productID).Take(&product)
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	c.connection.Delete(&product)
// 	return nil
// }
