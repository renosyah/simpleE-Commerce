package router

import (
	"fmt"
	"html/template"
	"net/http"

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

}

func CustomerLogin(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles(viper.GetString("template.customer_login"), viper.GetString("template.customer_header"), viper.GetString("template.customer_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}
	template.Execute(res, nil)

}

func HandleCustomerLogin(res http.ResponseWriter, req *http.Request) {

}
