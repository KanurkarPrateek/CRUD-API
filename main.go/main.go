package main

import (
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	Id string `json:"Id"`
	Isbn string `json:"isbn"`
	title string`json:"tile"`
	Director *Director `json:"director"`
}

type Director struct{
	FirstName string `json:"firstname"`
	LastNmae string `json:"lastname"`

}

var movies []Movie

func main(){
	r :=mux.NewRouter()
	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/moviesId",getMovie).Methods("GET")
	r.HandleFunc("/movies",)
		
		}  {
		
	})
}