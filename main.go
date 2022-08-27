package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/LakeCombs/mongo-golang/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {

	r := httprouter.New()

	uc := controllers.NewUserController(getSession())

	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		Fprintln(err)
		panic(err)

	}

	return s
}

func Fprintln(err error) {
	panic("unimplemented")
}
