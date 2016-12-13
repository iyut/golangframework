package data

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"framework/models"
)

type TaskRepository struct {
	C string
}

func (r *TaskRepository) Create(task *models.Task) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	obj_id := bson.NewObjectId()
	task.Id = obj_id
	task.CreatedOn = time.Now()
	task.Status = "Created"

	context.User = task.CreatedBy
	err := col.Insert(&task)
	return err
}

func (r *TaskRepository) Update(task *models.Task) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	// partial update on MogoDB
	err := col.Update(bson.M{"_id": task.Id},
		bson.M{"$set": bson.M{
			"name":        task.Name,
			"description": task.Description,
			"due":         task.Due,
			"status":      task.Status,
			"tags":        task.Tags,
		}})
	return err
}
func (r *TaskRepository) Delete(id string) error {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	err := col.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}
func (r *TaskRepository) GetAll() []models.Task {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	var tasks []models.Task
	iter := col.Find(nil).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
func (r *TaskRepository) GetById(id string) (task models.Task, err error) {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	err = col.FindId(bson.ObjectIdHex(id)).One(&task)
	return
}
func (r *TaskRepository) GetByUser(user string) []models.Task {

	context := NewContext()
	defer context.Close()
	col := context.DbCollection(r.C)

	var tasks []models.Task
	iter := col.Find(bson.M{"createdby": user}).Iter()
	result := models.Task{}
	for iter.Next(&result) {
		tasks = append(tasks, result)
	}
	return tasks
}
