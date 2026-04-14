package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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
	sql := "INSERT INTO products (category_id, name, description, price, stock, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, category_id, name, description, price, stock, created_at, updated_at"
	rows, err := p.db.Query(context.Background(), sql, newProduct.CategoryId, newProduct.Name, newProduct.Description, newProduct.Price, newProduct.Stock, time.Now(), time.Now())
	if err != nil {
		return models.Product{}, errors.New("Failed to create new product! : " + err.Error())
	}

	registeredProduct, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return models.Product{}, errors.New("Failed to create new product! : " + err.Error())
	}

	if p.rdb != nil {
		cacheKey := "products"
		err := p.rdb.Del(context.Background(), cacheKey).Err()
		if err != nil {
			return models.Product{}, errors.New("Redis Error : " + err.Error())
		}
	}

	return registeredProduct, nil
}

func (p ProductRepository) GetAllProducts() ([]models.Product, error) {
	cacheKey := "products"

	if p.rdb == nil {
		return []models.Product{}, errors.New("Redis connection not established!")
	}

	valueCache, err := p.rdb.Get(context.Background(), cacheKey).Result()

	if err == nil {
		var products []models.Product
		if err := json.Unmarshal([]byte(valueCache), &products); err != nil {
			return nil, err
		}
		return products, nil
	}
	fmt.Println("Cache miss")
	if err != redis.Nil {
		return nil, err
	}

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
		data, err := json.Marshal(products)
		if err != nil {
			fmt.Println("Failed to marshal cache")
		} else {
			err := p.rdb.Set(context.Background(), cacheKey, data, time.Hour).Err()
			if err != nil {
				fmt.Println("Failed to set cache")
			}
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

func (p ProductRepository) GetProductsByCategoryId(categoryId int) ([]models.Product, error) {
	sql := `SELECT id, category_id, name, description, price, stock, created_at, updated_at FROM products WHERE category_id = $1`
	rows, err := p.db.Query(context.Background(), sql, categoryId)
	if err != nil {
		return []models.Product{}, err
	}

	products, err := pgx.CollectRows(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return []models.Product{}, err
	}

	return products, nil
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

	if newData.Price == 0 {
		newData.Price = product.Price
	}

	if newData.Stock == 0 {
		newData.Stock = product.Stock
	}

	if newData.CategoryId == 0 {
		newData.CategoryId = product.CategoryId
	}

	sql := "UPDATE products SET category_id = $1, name = $2, description = $3, price = $4, stock = $5, updated_at = $6 WHERE id = $7 RETURNING id, category_id, name, description, price, stock, created_at, updated_at"

	rows, err := p.db.Query(context.Background(), sql, newData.CategoryId, newData.Name, newData.Description, newData.Price, newData.Stock, time.Now(), newData.Id)
	if err != nil {
		return models.Product{}, errors.New("Failed to update product! : " + err.Error())
	}

	updatedProduct, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[models.Product])
	if err != nil {
		return models.Product{}, errors.New("Failed to update product! : " + err.Error())
	}

	if p.rdb != nil {
		cacheKey := "products"
		err := p.rdb.Del(context.Background(), cacheKey).Err()
		if err != nil {
			return models.Product{}, errors.New("Redis Error : " + err.Error())
		}
	}

	return updatedProduct, nil
}

func (p ProductRepository) DeleteProduct(id int) (models.Product, error) {
	product, err := p.GetProductById(id)
	if err != nil {
		return models.Product{}, err
	}

	sql := `DELETE FROM products WHERE id = $1`
	_, err = p.db.Exec(context.Background(), sql, id)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}
