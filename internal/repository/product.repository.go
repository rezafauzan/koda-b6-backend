package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) (*ProductRepository, error) {
	return &ProductRepository{
		db: db,
	}, nil
}

func (u ProductRepository) CreateProduct(newProduct *models.Product) (models.Product, error) {
	sql := "INSERT INTO products (category_id, favorite_product, name, description, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, category_id, favorite_product, name, description, price, stock, created_at, updated_at"
	rows, err := u.db.Query(context.Background(), sql, newProduct.CategoryId, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Stock, time.Now(), time.Now())
	if err != nil {
		return models.Product{}, errors.New("Failed to create new product! : " + err.Error())
	}

	registeredProduct, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return models.Product{}, errors.New("Failed to create new product! : " + err.Error())
	}

	return registeredProduct, nil
}

func (u ProductRepository) GetAllProducts() ([]models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products`
	
	rows, err := u.db.Query(context.Background(), sql)

	if err != nil {
		return []models.Product{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])

	if err != nil {
		return []models.Product{}, errors.New("Failed to create response get all products! : " + err.Error())
	}

	return products, nil
}

func (u ProductRepository) GetProductById(productId int) (models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1`
	rows, err := u.db.Query(context.Background(), sql, productId)

	if err != nil {
		return models.Product{}, err
	}

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (u ProductRepository) GetAllProductsByName(productName string) ([]models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products WHERE name ILIKE '%' || $1 || '%'`
	rows, err := u.db.Query(context.Background(), sql, productName)

	if err != nil {
		return []models.Product{}, err
	}

	product, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return []models.Product{}, err
	}

	return product, nil
}

func (u ProductRepository) UpdateProduct(newData models.Product) (models.Product, error) {
	product, err := u.GetProductById(newData.Id)
	if err != nil {
		return models.Product{}, err
	}

	if newData.Name == "" {
		newData.Name = product.Name
	}

	if newData.Description == "" {
		newData.Description = product.Description
	}

	sql := `UPDATE products SET category_id = $1, favorite_product = $2, name = $3, description = $4, price = $5, campaign_id = $6, stock = $7, updated_at = $8 WHERE id = $9`

	_, err = u.db.Exec(context.Background(), sql, newData.CategoryId, newData.Name, newData.Description, newData.Price, newData.Stock, time.Now(), newData.Id)
	if err != nil {
		return models.Product{}, err
	}

	updatedProduct, err := u.GetProductById(newData.Id)
	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (u ProductRepository) DeleteProduct(id int) error {
	sql := `DELETE FROM products WHERE id = $1`
	_, err := u.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	return nil
}
