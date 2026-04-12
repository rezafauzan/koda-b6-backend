package repository

import (
	"context"
	"errors"
	"rezafauzan/koda-b6-golang/internal/models"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ProductRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewProductRepository(db *pgxpool.Pool, rdb *redis.Client) (*ProductRepository, error) {
	return &ProductRepository{
		db:  db,
		rdb: rdb,
	}, nil
}

func (p ProductRepository) CreateProduct(newProduct *models.Product) (models.Product, error) {
	sql := "INSERT INTO products (category_id, favorite_product, name, description, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, category_id, favorite_product, name, description, price, stock, created_at, updated_at"
	rows, err := p.db.Query(context.Background(), sql, newProduct.CategoryId, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Stock, time.Now(), time.Now())
	if err != nil {
		return models.Product{}, errors.New("Failed to create new product! : " + err.Error())
	}

	registeredProduct, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return models.Product{}, errors.New("Failed to create new product! : " + err.Error())
	}

	return registeredProduct, nil
}

func (p ProductRepository) GetAllProducts() ([]models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products`

	rows, err := p.db.Query(context.Background(), sql)

	if err != nil {
		return []models.Product{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])

	if err != nil {
		return []models.Product{}, errors.New("Failed to create response get all products! : " + err.Error())
	}

	if p.rdb != nil {
		cacheKey := "products"
		err := p.rdb.Del(context.Background(), cacheKey).Err()
		if err != nil {
			return []models.Product{}, errors.New("Redis Error : " + err.Error())
		}
	}

	return products, nil
}

func (p ProductRepository) GetProductById(productId int) (models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products WHERE id = $1`
	rows, err := p.db.Query(context.Background(), sql, productId)

	if err != nil {
		return models.Product{}, err
	}

	product, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p ProductRepository) GetAllProductsByName(productName string) ([]models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products WHERE name ILIKE '%' || $1 || '%'`
	rows, err := p.db.Query(context.Background(), sql, productName)

	if err != nil {
		return []models.Product{}, err
	}

	product, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return []models.Product{}, err
	}

	return product, nil
}

func (p ProductRepository) UpdateProduct(newData models.Product) (models.Product, error) {
	product, err := p.GetProductById(newData.Id)
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

	_, err = p.db.Exec(context.Background(), sql, newData.CategoryId, newData.Name, newData.Description, newData.Price, newData.Stock, time.Now(), newData.Id)
	if err != nil {
		return models.Product{}, err
	}

	updatedProduct, err := p.GetProductById(newData.Id)
	if err != nil {
		return models.Product{}, err
	}

	return updatedProduct, nil
}

func (p ProductRepository) DeleteProduct(id int) error {
	sql := `DELETE FROM products WHERE id = $1`
	_, err := p.db.Exec(context.Background(), sql, id)
	if err != nil {
		return err
	}
	return nil
}
