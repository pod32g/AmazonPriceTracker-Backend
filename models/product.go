package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/url"
	"time"
)

type ProductStruct struct {
	URL       string
	Price     float32
	Name      string
	UpdatedAt time.Time
}

type productType struct {
	db *gorm.DB
}

type Product interface {
	Add(product *ProductStruct) error
	First() (*ProductStruct, error)
	Find(id int64) (*ProductStruct, error)
	GetAll() ([]ProductStruct, error)
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

	db.AutoMigrate(&ProductStruct{})

	return &productType{
		db: db,
	}
}

func (p *productType) Add(product *ProductStruct) error {
	p.db.Create(product)
	return nil
}

func (p *productType) First() (*ProductStruct, error) {
	var product ProductStruct
	p.db.First(&product, 1)
	return &product, nil
}

func (p *productType) Find(id int64) (*ProductStruct, error) {
	var product ProductStruct
	p.db.First(&product, "Id = ?", id)
	return &product, nil
}

func (p *productType) GetAll() ([]ProductStruct, error) {
	var products []ProductStruct
	p.db.Find(&products)

	return products, nil
}

func (p *productType) Close() error {
	err := p.db.Close()

	return err
}
