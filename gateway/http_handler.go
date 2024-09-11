package main

import (
	"log"
	"net/http"
)


type handler struct{
//gateway
}

func NewHandler() *handler{
	return &handler{}
}

func (h*handler) registerRoutes(mux *http.ServeMux){
	mux.HandleFunc("POST /api/customers/{customerId}/orders", h.HandleCreateOrder)
}

func (h* handler) HandleCreateOrder(w http.ResponseWriter,r  *http.Request){
	log.Println("hello")
}