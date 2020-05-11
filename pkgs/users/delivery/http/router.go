package userDelivery

import (
	"github.com/gorilla/mux"
)


func (b UserDelivery) Start(router *mux.Router) {
	userRoute := router.PathPrefix("/v1/users").Subrouter()
	userRoute.Handle("/", b.CreateUser).Methods("POST")
	userRoute.Handle("/", b.FetchUsers).Methods("GET")
	userRoute.Handle("/{userID}", b.GetUserById).Methods("GET")
	userRoute.Handle("/{userID}", b.UpdateUserDetails).Methods("PUT")
	userRoute.Handle("/{userID}", b.DeleteUser).Methods("DELETE")
}
