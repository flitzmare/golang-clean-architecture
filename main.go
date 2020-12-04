package main

import (
	"golang-clean-architecture/infrastructure/datastore"
	"golang-clean-architecture/infrastructure/router"
	"golang-clean-architecture/registry"
)

func main() {
	reg := registry.NewRegistry(datastore.NewDB())
	r := router.NewRouter(reg.NewAppController())
	r.Run(":1234")
}

//func setupRouter() *gin.Engine {
//	r := gin.Default()
//	main := NewMainInstance()
//	r.POST("/users", main.user_post)
//	return r
//}

//type Main struct {
//	datastore gorm.DB
//}
//
//func NewMainInstance() Main {
//	return Main{}
//}

//func NewMainInstance(datastore gorm.DB) Main {
//	return Main{datastore: datastore}
//}

//func initdb() gorm.DB {
//	dsn := "user=riksasuviana dbname=example port=5432"
//	datastore, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//
//	datastore.AutoMigrate(&User{})
//	return *datastore
//}

//func (m Main) user_post(context *gin.Context) {
//	user := &RegisterRequest{
//		Name:     context.PostForm("Name"),
//		Username: context.PostForm("Username"),
//		Password: context.PostForm("Password"),
//	}
//
//	v := validator.New()
//	err := v.Struct(user)
//	if err != nil {
//		context.JSON(422, err.Error())
//		return
//	}
//
//	//tx := m.datastore.Create(&user)
//	//if tx.Error != nil {
//	//	panic(tx.Error)
//	//}
//	context.JSON(200, "OK")
//}

//type RegisterRequest struct {
//	Name      string `json:"Name" validate:"required"`
//	Username  string `json:"Username" validate:"required"`
//	Password  string
//}

//type Book struct {
//	ID        uint   `gorm:"primaryKey"`
//	BookName      string `validate:"required"`
//	CreatedAt time.Time
//	UpdatedAt time.Time
//	DeletedAt gorm.DeletedAt
//}
