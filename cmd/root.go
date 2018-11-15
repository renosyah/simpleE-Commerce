package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AuthScureDevelopment/lib-arjuna/cache"
	"github.com/AuthScureDevelopment/lib-arjuna/db"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	homedir "github.com/mitchellh/go-homedir"

	"github.com/renosyah/simpleE-Commerce/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	dbPool    *sql.DB
	cachePool *redis.Pool
)

var rootCmd = &cobra.Command{
	Use: "simple Ecommerce",
	Run: func(cmd *cobra.Command, args []string) {

		initDB()

		router.Init(dbPool, cachePool)

		r := mux.NewRouter()
		http.Handle("/", r)

		r.HandleFunc("/", router.CustomerHome)
		r.HandleFunc("/detail_product", router.ProductDetail)
		r.HandleFunc("/login", router.CustomerLogin)
		r.HandleFunc("/register", router.CustomerRegister)

		r.HandleFunc("/admin/login", router.AdminLoginPage)
		r.HandleFunc("/admin/handle_login", router.HandleAdminLogin)

		r.HandleFunc("/admin/home", router.AdminHomePage)

		http.Handle("/data/", http.StripPrefix("/data/", http.FileServer(http.Dir("data"))))
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
		http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
		http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir("fonts"))))

		http.ListenAndServe(fmt.Sprintf(":%d", viper.GetInt("app.port")), nil)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig, initCache)
}
func initConfig() {
	viper.SetConfigType("toml")
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/simpleEcommerce")
		viper.SetConfigName(".simpleEcommerce")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initDB() {
	dbOptions := db.DBOptions{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.name"),
		SSLMode:  viper.GetString("database.sslmode"),
	}
	dbConn, err := db.Connect(dbOptions)
	if err != nil {
		log.Fatalln("Error connect to DB: %v\n", err)
		return
	}
	dbPool = dbConn
}

func initCache() {
	cacheOptions := cache.CacheOptions{
		Host:        viper.GetString("cache.host"),
		Port:        viper.GetInt("cache.port"),
		Password:    viper.GetString("cache.password"),
		MaxIdle:     viper.GetInt("cache.max_idle"),
		IdleTimeout: viper.GetInt("cache.idle_timeout"),
		Enabled:     viper.GetBool("cache.enabled"),
	}
	cachePool = cache.Connect(cacheOptions)

}
