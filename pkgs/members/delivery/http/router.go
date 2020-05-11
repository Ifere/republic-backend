package memberDelivery

import "github.com/gorilla/mux"

func (b MemberDelivery) Start(router *mux.Router) {
	userRoute := router.PathPrefix("/v1/users").Subrouter()
	userRoute.Handle("/", b.CreateMember).Methods("POST")
	userRoute.Handle("/", b.FetchMembers).Methods("GET")
	userRoute.Handle("/{userID}", b.GetMemberById).Methods("GET")
	userRoute.Handle("/{userID}", b.UpdateMemberDetails).Methods("PUT")
	userRoute.Handle("/{userID}", b.DeleteMember).Methods("DELETE")
}
