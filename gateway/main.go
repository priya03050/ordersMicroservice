package main

import (
	"log"
	"net/http"
	_ "github.com/joho/godotenv/autoload"
	common "github.com/priya03050/commons"
)

var(
	httpAddr=common.EnvString("HTTP_ADDR",":3000")
)

func main(){
	mux:= http.NewServeMux()
	handler:= NewHandler()

	handler.registerRoutes(mux)
	log.Printf("server is running %s",httpAddr)
	if err:=http.ListenAndServe(httpAddr,mux);err!=nil{
		log.Fatal("failed to start http server")
	}
	
}