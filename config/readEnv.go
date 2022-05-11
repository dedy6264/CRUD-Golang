package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	AppName       string
	AppKey        string
	ServerPort    string
	DBUsername    string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
	DBSSLMode     string
	DBConnection  string
	MongoDBUri    string
	MongoDBPort   string
	MongoDBName   string
	MongoDataName string
)

func SetEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	AppName = os.Getenv("APP_NAME")
	AppKey = os.Getenv("APP_KEY")
	ServerPort = os.Getenv("APP_PORT")
	//----POSTGRESQL
	DBUsername = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASS")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	DBSSLMode = os.Getenv("SSL_MODE")
	DBConnection = os.Getenv("DB_DRIVER")
	//----MONGODB
	MongoDBUri = os.Getenv("MONGO_HOST")
	MongoDBPort = os.Getenv("MONGO_PORT")
	MongoDBName = os.Getenv("MONGO_DB")
	MongoDataName = os.Getenv("MONGO_DATA_NAME")

	fmt.Println("AppName :: ", AppName)
	fmt.Println("________________________________")
	fmt.Println("ServerPort :: ", ServerPort)
	fmt.Println("________________________________")
	fmt.Println("DBUsername :: ", DBUsername)
	fmt.Println("________________________________")
	fmt.Println("DBPassword :: ", DBPassword)
	fmt.Println("________________________________")
	fmt.Println("DBName :: ", DBName)
	fmt.Println("________________________________")
	fmt.Println("DBHost :: ", DBHost)
	fmt.Println("________________________________")
	fmt.Println("DBPort :: ", DBPort)
	fmt.Println("________________________________")
	fmt.Println("DBSSLMode :: ", DBSSLMode)
	fmt.Println("________________________________")
	fmt.Println("DBConnection :: ", DBConnection)
	fmt.Println("________________________________")
	fmt.Println("MongoDBUri :: ", MongoDBUri)
	fmt.Println("________________________________")
	fmt.Println("MongoDBPort :: ", MongoDBPort)
	fmt.Println("________________________________")
	fmt.Println("MongoDBName :: ", MongoDBName)
	fmt.Println("________________________________")
	fmt.Println("MongoDataName :: ", MongoDataName)
	fmt.Println("________________________________")
}
