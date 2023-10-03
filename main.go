package main

import (
	"encoding/json" //This library to render the data into json to send the data request
	"fmt"           //Basic formatting methods to get the stuff in the terminal
	"log"           //To log the data on the frontend of the application basically in the black screen
	"net/http"      //To import the webserver facility

	//To generate a random number
	//To convert the datatypes "from" and "to" to the string represenation

	"github.com/gorilla/mux"
)

type Director struct{
firstname string `json:"firstname"`
lastname string `json:"lastname"`
}

type Movie struct{
id string `json:"id"`
isbn string `json:"isbn"`
title string `json:"title"`
director *Director `json:"director"`
}

var movies []Movie//For now basically this is the storage we're using...and initially we're going to add some of the movies in it

func getMovie(){

}

func getMovies(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type", "application/json")
 json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
 w.Header().Set("Content-Type", "application/json")
 params:=mux.Vars(r)

 for index, item := range movies{
if item.id == params["id"]{
movies = append(movies[:index], movies[index+1:]...)
break
}
 }

}





func main(){
r:=mux.NewRouter() //returns a new router instance

movies = append(movies,Movie{id:"1", isbn:"1234", title:"Armageddon", director:&Director{firstname: "Vicky", lastname: "Tyagi"}})
movies = append(movies,Movie{id:"2", isbn:"123442", title:"Armageddon II", director:&Director{firstname: "John", lastname: "Doe"}})


r.HandleFunc("/movies", getMovies).Methods("GET")
r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
r.HandleFunc("/movies", createMovie).Methods("POST")
r.HandleFunc("/movies/{id}",updateMovie).Methods("UPDATE")
r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

fmt.Printf("Starting server at port:8000\n")
log.Fatal(http.ListenAndServe(":8000",r))

}
