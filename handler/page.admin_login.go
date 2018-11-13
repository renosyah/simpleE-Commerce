package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/securecookie"
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

var cookiehandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func HandleAdminLogin(res http.ResponseWriter, req *http.Request) {
	admin := model.AdminStruct{
		Username: req.FormValue("username"),
		Password: req.FormValue("password"),
	}

	result, err := admin.Login(dbPool)

	if err == nil {
		value := map[string]int64{
			"name": result.IdAdmin,
		}
		if encoded, err := cookiehandler.Encode("session", value); err == nil {
			cookie_ku := &http.Cookie{
				Name:  "session",
				Value: encoded,
				Path:  "/",
			}
			http.SetCookie(res, cookie_ku)
		}
	}

	if err != nil {
		fmt.Println("%r", err)
	}

	json.NewEncoder(res).Encode(result)

}
