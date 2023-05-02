package main

import (
	"encoding/json" //encode data while sending to postman
	"fmt"
	"log" //to log errors or print messages
	"math/rand"
	"net/http"
	"strconv" //convert string

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`

}
type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`


}

var movies []Movie

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies =append(movies[:index ],movies[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(movies)
	}

}
func getMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	params:=mux.Vars(r)
	for _,item:=range movies{
		if item.ID==params["id"]{
			json.NewEncoder(w).Encode(item)
			return

		}
		
	}

}
func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID=strconv.Itoa(rand.Intn(10000000))
	movies=append(movies,movie)
	json.NewEncoder(w).Encode(movie)


}

func updateMovie(w http.ResponseWriter,r *http.Request){
	// set json context type
	w.Header().Set("content-Type","application/json")
	// params
	params:=mux.Vars(r)
	// loop over the movie ,range
	// delete the movie with id that you send
	// add a new movie - the movie that we send body of position
	for index,item:=range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movie
			_=json.NewDecoder(r.Body).Decode(&movie)
			movie.ID=params["id"]
			movies=append(movies,movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func main(){
	r:=mux.NewRouter()
	movies=append(movies,Movie{ID: "1",Isbn: "438227",Title: "one",Director: &Director{Firstname: "john",Lastname: "m"}})
	movies=append(movies,Movie{ID: "2",Isbn: "458227",Title: "two",Director: &Director{Firstname: "michale",Lastname: "h"}})
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8080",r))
	
}