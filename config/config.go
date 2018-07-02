package config

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// AppConf is main app config
var AppConf = &appConfig{}

type appConfig struct {
	Postgres   *postgreConf
	Port       string
	UploadPath string
}

type postgreConf struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func (c *appConfig) Load(fileNames ...string) error {

	err := godotenv.Overload(fileNames...)
	if err != nil {
		log.Println(".env file not found, trying fetch environment variables")
	}

	if e, ok := os.LookupEnv("MSITE_UI_PORT"); ok {
		log.Printf("Msite port: %s\n", e)
		c.Port = ":" + e
	} else {
		log.Println("Msite port (default): 80")
		c.Port = ":80"
	}

	if e, ok := os.LookupEnv("MSITE_UPLOAD_PATH"); ok {
		log.Printf("Msite upload path: %s\n", e)
		c.UploadPath = e
	} else {
		log.Println("Msite upload path (default): /tmp")
		c.UploadPath = "/tmp"
	}

	c.Postgres, err = loadPostgreConfig()
	if err != nil {
		return err
	}

	return nil
}

func loadPostgreConfig() (*postgreConf, error) {
	Postgre := &postgreConf{}
	if e, ok := os.LookupEnv("POSTGRES_HOST"); ok {
		log.Printf("Setup Postgre host: %s\n", e)
		Postgre.Host = e
	} else {
		return nil, errors.New("POSTGRES_HOST not found")
	}

	if e, ok := os.LookupEnv("POSTGRES_PORT"); ok {
		log.Printf("Setup Postgre port: %s\n", e)
		Postgre.Port = e
	} else {
		log.Println("Setup default Postgres port: 5432")
		Postgre.Port = "5432"
	}

	if e, ok := os.LookupEnv("POSTGRES_DB"); ok {
		log.Printf("Setup Postgres db: %s\n", e)
		Postgre.DbName = e
	}

	if e, ok := os.LookupEnv("POSTGRES_USER"); ok {
		log.Printf("Setup Postgres user: %s\n", e)
		Postgre.User = e
	} else {
		return nil, errors.New("POSTGRES_USER not found")
	}

	if e, ok := os.LookupEnv("POSTGRES_PASSWORD"); ok {
		log.Printf("Setup Postgre password: %s\n", maskString(e, 0))
		Postgre.Password = e
	} else {
		return nil, errors.New("POSTGRES_PASSWORD not found")
	}

	return Postgre, nil
}

func maskString(s string, showLastSymbols int) string {
	if len(s) <= showLastSymbols {
		return s
	}
	return strings.Repeat("*", len(s)-showLastSymbols) + s[len(s)-showLastSymbols:]
}
