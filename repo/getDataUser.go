package repo

import (
	"context"
	"log"
	"sejutacita/config"
	"sejutacita/model"
)

func CheckDataUserMongo(filter interface{}) (model.ResponseCheckDataUser, string, string) {
	var res model.ResponseCheckDataUser
	var ctx = context.TODO()
	db := config.GetMongoDB()
	collection := db.Database("dbTest").Collection("colUser")
	//check data
	// filter := bson.M{"username": username}
	err := collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return res, "81", "data not found"
	}
	return res, "00", "Success"
}
func CheckUserMongo(filter interface{}) (model.ResponseLogin, string, string) {
	// pwd := helper.EncryptSHA1(u.Password)
	// u.Password = pwd
	var res model.ResponseLogin
	var ctx = context.TODO()
	db := config.GetMongoDB()
	collection := db.Database("dbTest").Collection("colUser")
	//check data
	// filter := bson.M{"username": u.Username, "password": u.Password}
	err := collection.FindOne(ctx, filter).Decode(&res)
	if err != nil {
		return res, "81", "data not found"
	}
	// //set jwt
	// token, code, desc := helper.TokenJWT(res.Role, res.Email, u.Username)
	// if code != "00" {
	// 	return res, code, desc
	// }
	// res.Token = token
	return res, "00", "Success"
}

func InsertDataUserMongo(data model.DataRegister, filterUsername interface{}, filterEmail interface{}) (string, string, string) {
	var res model.DataRegister
	var ctx = context.TODO()
	db := config.GetMongoDB()
	collection := db.Database("dbTest").Collection("colUser")
	//check data
	err := collection.FindOne(ctx, filterUsername).Decode(&res)
	if err == nil {
		return "", "81", "Username atau email sudah terdaftar"
	}
	err = collection.FindOne(ctx, filterEmail).Decode(&res)
	if err == nil {
		return "", "81", "Username atau email sudah terdaftar"
	}

	//insert
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err.Error())
	}

	return "Registered", "00", "Success"
}

func UpdateDataUserMongo(filter interface{}, update interface{}) (string, string, string) {

	var ctx = context.TODO()
	db := config.GetMongoDB()
	collection := db.Database("dbTest").Collection("colUser")
	//check data
	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return "", "81", "Update Gagal ::" + err.Error()
	}
	return "Success", "00", "Data Updated"
}
func DeleteDataUserMongo(delete interface{}) (string, string, string) {
	var ctx = context.TODO()
	db := config.GetMongoDB()
	collection := db.Database("dbTest").Collection("colUser")
	// delete := bson.M{"username": username}
	_, err := collection.DeleteMany(ctx, delete)
	if err != nil {
		return "", "81", "Update Gagal ::" + err.Error()
	}
	return "Success", "00", "Data Deleted"
}
