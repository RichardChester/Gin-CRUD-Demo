package model

import (
	"gopkg.in/mgo.v2/bson"
	"go-crud-demo/db"
	"go-crud-demo/forms"
)

type User struct {
	Id 					bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Username 		string  `json:"username"`
	Password		string  `json:"password"`
	Age 				int	    `json:"age"`
	Sex			    string	`json:"sex"`
}
type UserModel struct{}

var (
	dbConnect = db.NewConnection("localhost")
	dbName = "go-test"
	collectionName = "user"
)

//创建一个用户
func (m *UserModel) Create (data forms.CreateUserCommand) error {
	collection := dbConnect.Use(dbName,collectionName)
	err := collection.Insert(bson.M{
		"username":data.Username,
		"password":data.Password,
		"age":data.Age,
		"sex":data.Sex,
	})
	return err
}
//获取所有用户
func (m *UserModel) GetAll () (list []User,err error){
	collection := dbConnect.Use(dbName,collectionName)
	err = collection.Find(bson.M{}).All(&list)
	return
}
//查找某一个用户
func (m *UserModel) GetOne (id string) (user User,err error){
	collection := dbConnect.Use(dbName,collectionName)
	err = collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return
}
//更新一个用户
func (m *UserModel) UpdateOne (id string, data forms.UpdateUserCommand) (err error){
	collection := dbConnect.Use(dbName,collectionName)
	err = collection.UpdateId(bson.ObjectIdHex(id),data)

	return
}
//删除一个用户
func (m *UserModel) DeleteOne (id string) (err error) {
	collection := dbConnect.Use(dbName,collectionName)
	err = collection.RemoveId(id)

	return
}