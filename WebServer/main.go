package main

import (
    "fmt"
    "net/http"

)


func homeHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to my first Go web server!..")              
}
    

func aboutHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "This is about page.")
}

func main(){
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)

    fmt.Println("Server running on http://localhost:8080")
    // fmt.Println("Opens the last ")


    http.ListenAndServe(":8080", nil)
}