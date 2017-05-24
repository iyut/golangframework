package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"framework/models"
)

type HomeRepository struct {
	C string
}

func (r *HomeRepository) Create(home *models.MS_Home) error {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)		
	
	obj_id := bson.NewObjectId()
	home.Id = obj_id
	home.CreatedOn = time.Now()
   	home.Status = "Created"	
	
	context.User = home.CreatedBy 	
	err := col.Insert(&home)
	return err
}

func (r *HomeRepository) Update(home *models.MS_Home) error {
	// partial update on MogoDB
	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)			
	
	err := col.Update(bson.M{"_id": home.Id},
		bson.M{"$set": bson.M{
			"address":      home.Address,
			"rt":		home.RT,
			"description":  home.Description,
			"createdon":    home.CreatedOn,
			"status":      	home.Status,
			"createdby":	home.CreatedBy,
		}})
	return err
}
func (r *HomeRepository) Delete(id string) error {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)			

	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *HomeRepository) GetAll() []models.MS_Home {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)		
	

	var homes []models.MS_Home
	iter := col.Find(nil).Iter()
	result := models.MS_Home{}
	for iter.Next(&result) {
		homes = append(homes, result)
	}
	return homes
}
func (r *HomeRepository) GetById(id string) (home models.MS_Home, err error) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)		
	
	err = col.FindId(bson.ObjectIdHex(id)).One(&home)
	return
}

func (r *HomeRepository) GetByUser(user string) []models.MS_Home {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	var homes []models.MS_Home
	iter := col.Find(bson.M{"createdby": user}).Iter()
	result := models.MS_Home{}
	for iter.Next(&result) {
		homes = append(homes, result)
	}
	return homes
}
