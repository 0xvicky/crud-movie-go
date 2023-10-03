package main

import(
	"fmt" //Basic formatting methods to get the stuff in the terminal
	"net/http" //To import the webserver facility
	"math/rand" //To generate a random number
    "strconv"  //To convert the datatypes "from" and "to" to the string represenation
	"log" //To log the data on the frontend of the application basically in the black screen
	"encoding/json" //This library to render the data into json to send the data request
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


var movies []Movie

func main(){
r:=mux.NewRouter() //returns a new router instance
}
