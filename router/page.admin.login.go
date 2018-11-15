package router

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/renosyah/simpleE-Commerce/model"
	"github.com/spf13/viper"
)

func AdminLoginPage(res http.ResponseWriter, req *http.Request) {
	template, err := template.ParseFiles(viper.GetString("template.admin_login"), viper.GetString("template.admin_header"), viper.GetString("template.admin_footer"))
	if err != nil {
		fmt.Println("%r", err)
	}

	template.Execute(res, nil)
}

func HandleAdminLogin(res http.ResponseWriter, req *http.Request) {
	admin := model.AdminStruct{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	result, err := admin.Login(dbPool)

	if err == nil {
		SetCookie("admin_session", fmt.Sprint("", result.IdAdmin), res)
	}

	if err != nil {
		fmt.Println("%r", err)
	}

	json.NewEncoder(res).Encode(result)

}
