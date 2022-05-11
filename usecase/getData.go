package usecase

import (
	"sejutacita/helper"
	"sejutacita/model"
	"sejutacita/repo"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Login(u model.DataUser) model.ResponseGlobal {
	var result model.ResponseGlobal
	t := time.Now()
	//cek data mongo
	pwd := helper.EncryptSHA1(u.Password)
	u.Password = pwd
	filter := bson.M{"username": u.Username, "password": u.Password}
	res, code, desc := repo.CheckUserMongo(filter)
	if code != "00" {
		result.Status = code
		result.StatusDateTime = t
		result.StatusDesc = desc
		result.Result = res
		return result
	}
	//set jwt
	token, code, desc := helper.TokenJWT(res.Role, res.Email, u.Username)
	if code != "00" {
		// return res, code, desc
		result.Status = code
		result.StatusDateTime = t
		result.StatusDesc = desc
		result.Result = res
		return result
	}
	res.Token = token
	result.Status = code
	result.StatusDateTime = t
	result.StatusDesc = desc
	result.Result = res
	return result

}
func InsertData(u model.DataRegister) model.ResponseGlobal {
	var result model.ResponseGlobal
	t := time.Now()
	pwd := helper.EncryptSHA1(u.Password)
	u.Password = pwd
	filterUsername := bson.M{"username": u.Username}
	filterEmail := bson.M{"email": u.Email}

	res, code, desc := repo.InsertDataUserMongo(u, filterUsername, filterEmail)
	if code != "00" {
		result.Status = code
		result.StatusDateTime = t
		result.StatusDesc = desc
		result.Result = res
		return result
	}
	result.Status = code
	result.StatusDateTime = t
	result.StatusDesc = desc
	result.Result = res
	return result
}
func UpdateDataUser(u model.DataUpdate, username string, role string) model.ResponseGlobal {
	var result model.ResponseGlobal
	t := time.Now()
	if role == "admin" {
		if u.UsernameTarget != "" {
			username = u.UsernameTarget
		}
	}
	//cek data eksist
	filter := bson.M{"username": u.Username}
	resp, code, desc := repo.CheckDataUserMongo(filter)
	if code == "00" {
		result.Status = "41"
		result.StatusDateTime = t
		result.StatusDesc = "username sudah digunakan"
		result.Result = resp
		return result
	}
	if u.Email != "" {
		filter = bson.M{"email": u.Email}
		resp, code, desc = repo.CheckDataUserMongo(filter)
		if code == "00" {
			result.Status = "41"
			result.StatusDateTime = t
			result.StatusDesc = "email sudah digunakan"
			result.Result = resp
			return result
		}
	}
	//update
	filter = bson.M{"username": username}
	update := bson.M{"$set": bson.M{
		"username": u.Username,
		"email":    u.Email,
		"name":     u.Name,
	}}
	res, code, desc := repo.UpdateDataUserMongo(filter, update)
	if code != "00" {
		result.Status = code
		result.StatusDateTime = t
		result.StatusDesc = desc
		result.Result = res
		return result
	}
	result.Status = code
	result.StatusDateTime = t
	result.StatusDesc = desc
	result.Result = res
	return result
}
func DeleteDataUser(u model.DeleteUser, username string) model.ResponseGlobal {
	t := time.Now()
	var result model.ResponseGlobal
	if username == u.Username {
		result.Status = "31"
		result.StatusDateTime = t
		result.StatusDesc = "sorry, you are admin, cannot delet yourself"
		result.Result = ""
		return result
	}
	filter := bson.M{"username": u.Username}
	res, code, desc := repo.CheckDataUserMongo(filter)
	if code != "00" {
		result.Status = code
		result.StatusDateTime = t
		result.StatusDesc = desc
		result.Result = res
		return result
	}
	delete := bson.M{"username": u.Username}

	resp, code, desc := repo.DeleteDataUserMongo(delete)
	if code != "00" {
		result.Status = code
		result.StatusDateTime = t
		result.StatusDesc = desc
		result.Result = resp
		return result
	}
	result.Status = "00"
	result.StatusDateTime = t
	result.StatusDesc = "Success"
	result.Result = resp
	return result
}
