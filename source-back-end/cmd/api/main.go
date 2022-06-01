package main

import (
	"backend/config"
	"backend/models"
	"context"
	"database/sql"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

/*
type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
	jwt struct {
		secret string
	}
}
*/

type AppStatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}

type application struct {
	logger *log.Logger
	models models.Models
}

// main문 실행되기 전 우선 실행
func init() {
	profile := initProfile()
	setRuntimeConfig(profile)
}

// 환경변수는 profile을 확인하기 위해 하나만 설정
func initProfile() string {
	var profile string
	profile = os.Getenv("GO_PROFILE")
	if len(profile) <= 0 {
		profile = "local"
	}
	fmt.Println("GOLANG_PROFILE: " + profile)
	return profile
}

// profile 기반으로 config 파일을 읽고 전역변수에 언마샬링
func setRuntimeConfig(profile string) {
	viper.AddConfigPath(".")
	// 환경변수에서 읽어온 profile 이름의 yaml 파일을 configPath로 설정
	viper.SetConfigName(profile)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// viper는 읽어온 설정파일의 정보를 가지고 있으니, 전연벽수에 언마샬링해서 사용
	err = viper.Unmarshal(&config.RuntimeConf)
	if err != nil {
		panic(err)
	}

	// 설정파일 변경시 재적용
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		var err error
		err = viper.ReadInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
		err = viper.Unmarshal(&config.RuntimeConf)
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	viper.WatchConfig()
}

func main() {
	/*
		var cfg config.RuntimeConfig

		log.Println(os.Getenv("CONFIG_VAR_TWO"))

		flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on")
		flag.StringVar(&cfg.env, "env", "development", "Application environment (development|production")
		flag.StringVar(&cfg.db.dsn, "dsn",
			fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				"0.0.0.0", "5432", "kyo", "1q2w3e4r", "postgres"),
			"Postgres connection string")
		flag.StringVar(&cfg.jwt.secret, "jwt-secret", "2dce505d96a53c5768052ee90f3df2055657518dad489160df9913f66042e160", "secret")
		flag.Parse()

		// read jwt secret from env
		//cfg.jwt.secret = os.Getenv("GO_MOVIES_JWT")
	*/

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(config.RuntimeConf.Datasource)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		models: models.NewModels(db),
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.RuntimeConf.Server.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", config.RuntimeConf.Server.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

func openDB(ds config.Datasource) (*sql.DB, error) {

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		ds.Host, ds.Port, ds.UserName, ds.Password, ds.DatabaseName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
