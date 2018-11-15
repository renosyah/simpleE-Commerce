package router

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/renosyah/simpleE-Commerce/model"
	"github.com/spf13/viper"
)

type CustomerHomeStruct struct {
	Costumer   *model.Costumer         `json:"costumer"`
	Categories []model.ProductCategory `json:"categories"`
	Products   []model.Product         `json:"products"`
}

func CustomerHome(res http.ResponseWriter, req *http.Request) {

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

	id_costumer, err := strconv.ParseInt(GetCookie(costumer_session, req), 10, 64)
	if err != nil {
		//fmt.Println("%r", err)
	}

	u := &model.Costumer{IdCostumer: id_costumer}
	costumer, err := u.GetOneCostumer(dbPool)
	if err != nil {
		//fmt.Println("%r", err)
	}

	template, err := template.ParseFiles(viper.GetString("template.customer_home"), viper.GetString("template.customer_header"), viper.GetString("template.customer_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}
	template.Execute(res, &CustomerHomeStruct{Categories: categories, Products: products, Costumer: costumer})
}

type ProductDetailStruct struct {
	Costumer *model.Costumer `json:"costumer"`
	Product  *model.Product  `json:"product"`
}

func ProductDetail(res http.ResponseWriter, req *http.Request) {

	id := req.FormValue("id_product")
	id_int64, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Println("%r", err)
	}

	template, err := template.ParseFiles(viper.GetString("template.customer_product_detail"), viper.GetString("template.customer_header"), viper.GetString("template.customer_footer"))
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

	id_costumer, err := strconv.ParseInt(GetCookie(costumer_session, req), 10, 64)
	if err != nil {
		//fmt.Println("%r", err)
	}

	u := &model.Costumer{IdCostumer: id_costumer}
	costumer, err := u.GetOneCostumer(dbPool)
	if err != nil {
		//fmt.Println("%r", err)
	}
	template.Execute(res, &ProductDetailStruct{Product: product, Costumer: costumer})
}
