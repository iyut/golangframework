package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	MS_Home struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Address     string        `json:"address"`
		RT					string				`json:"rt"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_RT struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		RT					string				`json:"rt"`
		RW					string				`json:"rw"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_RW struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		RW					string				`json:"rw"`
		Villages		string				`json:"villages"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_Villages struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Villages		string				`json:"villages"`
		Regency			string				`json:"regency"`
		PostCode		string				`json:"postcode"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_Regency struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Regency			string				`json:"regency"`
		RegencyType string				`json:"regencytype"`
		Province		string				`json:"province"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_Regency_Type struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		RegencyType	string				`json:"regencytype"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_Regency struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Regency			string				`json:"regency"`
		RegencyType string				`json:"regencytype"`
		Province		string				`json:"province"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_Province struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Province		string				`json:"province"`
		Country			string				`json:"country"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	MS_Country struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Country			string				`json:"country"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Status      string        `json:"status,omitempty"`
	}
	Task struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
	TaskNote struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId      bson.ObjectId `json:"taskid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
)
