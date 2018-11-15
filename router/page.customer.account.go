package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/renosyah/simpleE-Commerce/model"
	"github.com/spf13/viper"
)

func CustomerRegister(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles(viper.GetString("template.customer_register"), viper.GetString("template.customer_header"), viper.GetString("template.customer_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}
	template.Execute(res, nil)
}

func HandleCustomerRegister(res http.ResponseWriter, req *http.Request) {
	costumer := &model.Costumer{
		Name:     req.FormValue("name"),
		UserName: req.FormValue("username"),
		Email:    req.FormValue("email"),
		Address:  req.FormValue("address"),
		Password: req.FormValue("password"),
	}
	value, err := costumer.AddCostumer(dbPool)
	if err == nil {

		SetCookie(costumer_session, fmt.Sprint(value), res)
		http.Redirect(res, req, "/", 302)

	} else {
		fmt.Println("%r", err)
		http.Redirect(res, req, "/login", 302)
	}
}

func CustomerLogin(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles(viper.GetString("template.customer_login"), viper.GetString("template.customer_header"), viper.GetString("template.customer_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}
	template.Execute(res, nil)

}

func HandleCustomerLogin(res http.ResponseWriter, req *http.Request) {
	costumer := &model.Costumer{
		UserName: req.FormValue("username"),
		Email:    req.FormValue("username"),
		Password: req.FormValue("password"),
	}
	result, err := costumer.GetLoginCostumer(dbPool)
	if err == nil && result.IdCostumer != 0 {

		SetCookie(costumer_session, fmt.Sprint(result.IdCostumer), res)
		http.Redirect(res, req, "/", 302)

	} else {
		http.Redirect(res, req, "/login", 302)
	}

}

func HandleCustomerLogout(res http.ResponseWriter, req *http.Request) {
	ClearCookie(costumer_session, res)
	http.Redirect(res, req, "/", 302)
}
