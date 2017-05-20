package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"framework/models"
)

type CountryRepository struct {
	C string
}

func (r *CountryRepository) Create(country *models.MS_Country) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	obj_id := bson.NewObjectId()
	country.Id = obj_id
	country.CreatedOn = time.Now()
	country.Status = "Created"

	context.User = country.CreatedBy
	err := col.Insert(&country)
	return err
}

func (r *CountryRepository) Update(country *models.MS_Country) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	// partial update on MogoDB
	err := col.Update(bson.M{"_id": country.Id},
		bson.M{"$set": bson.M{
			"name":        country.Name,
			"description": country.Description,
			"due":         country.Due,
			"status":      country.Status,
			"tags":        country.Tags,
		}})
	return err
}
func (r *CountryRepository) Delete(id string) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *CountryRepository) GetAll() []models.MS_Country {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	var countries []models.MS_Country
	iter := col.Find(nil).Iter()
	result := models.MS_Country{}
	for iter.Next(&result) {
		countries = append(countries, result)
	}
	return countries
}
func (r *CountryRepository) GetById(id string) (country models.MS_Country, err error) {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	err = col.FindId(bson.ObjectIdHex(id)).One(&country)
	return
}
func (r *CountryRepository) GetByUser(user string) []models.MS_Country {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	var countries []models.MS_Country
	iter := col.Find(bson.M{"createdby": user}).Iter()
	result := models.MS_Country{}
	for iter.Next(&result) {
		countries = append(countries, result)
	}
	return countries
}
