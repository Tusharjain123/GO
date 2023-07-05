package routers

import (
	controller "mondoSetup/controllers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", controller.GetAllUserData).Methods("GET")
	router.HandleFunc("/api/user", controller.CreatingUser).Methods("POST")
	router.HandleFunc("/api/user/{id}", controller.ChangeUser).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/deletedata", controller.DeleteAllUser).Methods("DELETE")
	return router
}
