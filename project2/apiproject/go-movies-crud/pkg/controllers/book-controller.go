package routes

import (
	"github.com/Akkishiva/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes=func(router *mux.Router){
	router.HandelFunc("/book/",controllers.CreateBook).Method("POST")
	router.HandelFunc("/book/",controllers.GetBook).Method("GET")
	router.HandelFunc("/book/{bookId}",controllers.GetBookById).Method("GET")
	router.HandelFunc("/book/{bookId}",controllers.UpdateBook).Method("PUT")
	router.HandelFunc("/book/{bookId}",controllers.DeleteBook).Method("Delete")

}


