package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/url"
	"time"
)

type _product struct {
	gorm.Model
	URL       string
	price     float64
	updatedAt time.Time
}

type productType struct {
	db *gorm.DB
}

type Product interface {
	Add(product *_product) error
	First() (*_product, error)
	Find(id int64) (*_product, error)
	GetAll() ([]_product, error)
	Close() error
}

func New() Product {

	dsn := url.URL{
		User:     url.UserPassword("postgres", "password"),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:     "postgres",
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open("postgres", dsn.String())

	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&_product{})

	return &productType{
		db: db,
	}
}

func (p *productType) Add(product *_product) error {
	p.db.Create(product)
	return nil
}

func (p *productType) First() (*_product, error) {
	var product _product
	p.db.First(&product, 1)
	return &product, nil
}

func (p *productType) Find(id int64) (*_product, error) {
	var product _product
	p.db.First(&product, "Id = ?", id)
	return &product, nil
}

func (p *productType) GetAll() ([]_product, error) {
	var products []_product
	p.db.Find(&products)

	return products, nil
}

func (p *productType) Close() error {
	err := p.db.Close()

	return err
}
