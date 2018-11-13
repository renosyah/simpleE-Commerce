package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/renosyah/simpleE-Commerce/model"
	"github.com/spf13/viper"
)

type UserHomeStruct struct {
	Categories []model.ProductCategory `json:"categories"`
	Products   []model.Product         `json:"products"`
}

func UserHome(res http.ResponseWriter, req *http.Request) {

	search_name := req.FormValue("search_product")
	category := req.FormValue("category")

	id_category, err := strconv.ParseInt(category, 10, 64)
	if err != nil {
		id_category = 0
	}

	p := &model.Product{
		ProductName: search_name,
		Category: &model.ProductCategory{
			IdProductCategory: id_category,
		},
	}
	products, err := p.GetAllProductByCategoryAndName(dbPool)
	if err != nil {
		fmt.Println("%r", err)
	}
	c := &model.ProductCategory{}
	categories, err := c.GetAllCategory(dbPool)
	if err != nil {
		fmt.Println("%r", err)
	}

	template, err := template.ParseFiles(viper.GetString("template.user_home"), viper.GetString("template.user_header"), viper.GetString("template.user_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}
	template.Execute(res, &UserHomeStruct{Categories: categories, Products: products})
}

type ProductDetailStruct struct {
	Product *model.Product
}

func ProductDetail(res http.ResponseWriter, req *http.Request) {

	id := req.FormValue("id_product")
	id_int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("%r", err)
	}

	template, err := template.ParseFiles(viper.GetString("template.user_product_detail"), viper.GetString("template.user_header"), viper.GetString("template.user_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}
	p := &model.Product{
		IdProduct: id_int64,
	}
	product, err := p.GetOneProduct(dbPool)
	if err != nil {
		fmt.Println("%r", err)
	}

	template.Execute(res, &ProductDetailStruct{Product: product})
}
