package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *sql.DB
var mongoDB *mongo.Client

func SetupConnect() (*sql.DB, error) {

	var connection = fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		DBUsername, DBPassword, DBName, DBHost, DBPort, DBSSLMode)
	fmt.Println("Connection Info            : ", DBConnection, connection)

	db, err := sql.Open(DBConnection, connection)
	if err != nil {
		return db, err
	}
	fmt.Println("berhasil koneksi")

	return db, nil
}

func SetConnectionDB() {
	var err error
	db, err = SetupConnect()

	if err != nil {
		fmt.Println("Gagal Konek Database")
	}
}

func CloseConnectionDB() {
	db.Close()
}

func DbConn() *sql.DB {
	return db
}
func ConnectMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(MongoDBUri+":"+MongoDBPort))
	return client
}

func GetMongoDB() *mongo.Client {
	return mongoDB
}

func SetConnectionsMongo() {
	mongoDB = ConnectMongo()
}

func CloseConnectionsMongo() {
	var ctx = context.TODO()
	mongoDB.Disconnect(ctx)
}
