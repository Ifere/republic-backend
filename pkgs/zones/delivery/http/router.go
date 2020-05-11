package zoneDelivery

import "github.com/gorilla/mux"

func (b ZoneDelivery) Start(router *mux.Router) {
	userRoute := router.PathPrefix("/v1/users").Subrouter()
	userRoute.Handle("/", b.CreateZone).Methods("POST")
	userRoute.Handle("/", b.FetchZones).Methods("GET")
	userRoute.Handle("/{userID}", b.GetZoneById).Methods("GET")
	userRoute.Handle("/{userID}", b.UpdateZoneDetails).Methods("PUT")
	userRoute.Handle("/{userID}", b.DeleteZone).Methods("DELETE")
}