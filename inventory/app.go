package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
) 

const DbName = "inventory"
const DbUser = "root"
const DbPassword = "1234"

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func(app *App) Initalize() error{
	connectionstring:=fmt.Sprintf("%v:%v@tcp(127.0.0.1:3036)/%v",DbUser,DbPassword,DbName)
    var err error
	app.DB,err=sql.Open("mysql",connectionstring)
	if err!=nil{
		return err
	}
	app.Router=mux.NewRouter().StrictSlash(true)
	return nil

func(app *App) Run(address string){
	log.Fatal(http.ListenAndServe(address,app.Router))
}
func sendResponse(w http.ResponseWriter,statusCode int,payload interface{}){
	response,_:=json.Marshal(payload)
	w.Header.Set("content-type","application.json")
	w.WriteHeader(statusCode)
	w.Write(response)

}
func sendError(w http.ResponseWriter,statusCode int,err string){
	error_message:=map[string]{"error":err}
	sendResponse(w,statusCode,error_message)
}

func(app *App) handleroutes(){
	app.Router.HandleFunc("/products",getProducts).Methods("GET")
}

}	

}