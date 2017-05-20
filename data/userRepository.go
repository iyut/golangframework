package data

import (
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"

	"framework/models"
)

type UserRepository struct {
	C string
}

func (r *UserRepository) CreateUser(user *models.User) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	obj_id := bson.NewObjectId()
	user.Id = obj_id
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	//clear the incoming text password
	user.Password = ""
	err = col.Insert(&user)
	return err
}

func (r *UserRepository) Login(user models.User) (u models.User, err error) {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	err = col.Find(bson.M{"email": user.Email}).One(&u)
	if err != nil {
		return
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}

func (r *UserRepository) Delete(id string) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
