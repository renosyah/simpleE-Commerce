package model

import (
	"database/sql"
	"fmt"
)

type Product struct {
	IdProduct   int64            `json:"id_product"`
	Category    *ProductCategory `json:"product_category"`
	Details     []ProductDetail  `json:"detail_product"`
	Images      []ProductImage   `json:"image_product"`
	ProductName string           `json:"product_name"`
	Stock       int32            `json:"stock"`
	Price       int32            `json:"price"`
	Curency     string           `json:"curency"`
}

func (p *Product) GetOneProduct(db *sql.DB) (*Product, error) {
	product := &Product{
		Category: &ProductCategory{},
		Details:  []ProductDetail{},
		Images:   []ProductImage{},
	}
	query := `SELECT id_product,id_product_category,product_name,stock,price,curency FROM ecommerce.product WHERE id_product=$1`
	err := db.QueryRow(fmt.Sprintf(query), p.IdProduct).Scan(&product.IdProduct,
		&product.Category.IdProductCategory,
		&product.ProductName,
		&product.Stock,
		&product.Price,
		&product.Curency,
	)
	product.Category, err = product.Category.GetOneProductCategory(db)
	if err != nil {
		return product, err
	}
	product.Details, err = product.GetAllProductDetail(db)
	if err != nil {
		return product, err
	}
	product.Images, err = product.GetAllProductImages(db)
	if err != nil {
		return product, err
	}
	return product, err
}

func (p *Product) GetAllProductByCategoryAndName(db *sql.DB) ([]Product, error) {
	var rows *sql.Rows
	var err error

	products := []Product{}

	if p.Category.IdProductCategory == 0 && p.ProductName != "" {

		query := `SELECT id_product,id_product_category,product_name,stock,price,curency FROM ecommerce.product WHERE product_name LIKE $1`
		rows, err = db.Query(fmt.Sprintf(query), "%"+p.ProductName+"%")

	} else if p.ProductName == "" && p.Category.IdProductCategory != 0 {

		query := `SELECT id_product,id_product_category,product_name,stock,price,curency FROM ecommerce.product WHERE id_product_category=$1`
		rows, err = db.Query(fmt.Sprintf(query), p.Category.IdProductCategory)

	} else if p.ProductName != "" && p.Category.IdProductCategory != 0 {

		query := `SELECT id_product,id_product_category,product_name,stock,price,curency FROM ecommerce.product WHERE id_product_category=$1 AND product_name LIKE $2`
		rows, err = db.Query(fmt.Sprintf(query), p.Category.IdProductCategory, "%"+p.ProductName+"%")

	} else {
		query := `SELECT id_product,id_product_category,product_name,stock,price,curency FROM ecommerce.product`
		rows, err = db.Query(fmt.Sprintf(query))

	}

	if err != nil {
		return products, err
	}
	for rows.Next() {

		product := Product{
			Category: &ProductCategory{},
			Details:  []ProductDetail{},
			Images:   []ProductImage{},
		}
		rows.Scan(&product.IdProduct,
			&product.Category.IdProductCategory,
			&product.ProductName,
			&product.Stock,
			&product.Price,
			&product.Curency,
		)
		product.Category, err = product.Category.GetOneProductCategory(db)
		if err != nil {
			return products, err
		}
		product.Details, err = product.GetAllProductDetail(db)
		if err != nil {
			return products, err
		}
		product.Images, err = product.GetAllProductImages(db)
		if err != nil {
			return products, err
		}

		products = append(products, product)
	}
	return products, nil
}

type ProductCategory struct {
	IdProductCategory int64  `json:"id_product_category"`
	CategoryName      string `json:"category_name"`
}

func (c *ProductCategory) GetOneProductCategory(db *sql.DB) (*ProductCategory, error) {
	result := &ProductCategory{}
	query := `SELECT id_product_category,category_name FROM ecommerce.product_category WHERE id_product_category=$1`
	err := db.QueryRow(fmt.Sprintf(query), c.IdProductCategory).Scan(&result.IdProductCategory, &result.CategoryName)
	return result, err
}
func (c *ProductCategory) GetAllCategory(db *sql.DB) ([]ProductCategory, error) {
	categories := []ProductCategory{}
	query := `SELECT id_product_category,category_name FROM ecommerce.product_category`
	rows, err := db.Query(fmt.Sprintf(query))
	if err != nil {
		return categories, err
	}
	for rows.Next() {
		category := ProductCategory{}
		rows.Scan(&category.IdProductCategory, &category.CategoryName)
		categories = append(categories, category)
	}
	return categories, nil
}

type ProductDetail struct {
	IdProductDetail int64  `json:"id_product_detail"`
	IdProduct       int64  `json:"id_product"`
	Description     string `json:"description"`
}

func (p *Product) GetAllProductDetail(db *sql.DB) ([]ProductDetail, error) {
	productDetails := []ProductDetail{}
	query := `SELECT id_product_detail,id_product,description FROM ecommerce.product_detail WHERE id_product=$1`
	rows, err := db.Query(fmt.Sprintf(query), p.IdProduct)
	if err != nil {
		return productDetails, err
	}
	for rows.Next() {
		productDetail := ProductDetail{}
		rows.Scan(&productDetail.IdProductDetail,
			&productDetail.IdProduct,
			&productDetail.Description,
		)
		productDetails = append(productDetails, productDetail)

	}
	return productDetails, nil
}

type ProductImage struct {
	IdProductImage int64  `json:"id_product_image"`
	IdProduct      int64  `json:"id_product"`
	UrlImage       string `json:"url_image"`
}

func (p *Product) GetAllProductImages(db *sql.DB) ([]ProductImage, error) {
	productImages := []ProductImage{}
	query := `SELECT id_product_image,id_product,url_image FROM ecommerce.product_image WHERE id_product=$1`
	rows, err := db.Query(fmt.Sprintf(query), p.IdProduct)
	if err != nil {
		return productImages, err
	}
	for rows.Next() {
		productImage := ProductImage{}
		rows.Scan(&productImage.IdProductImage,
			&productImage.IdProduct,
			&productImage.UrlImage,
		)
		productImages = append(productImages, productImage)
	}

	return productImages, nil
}
