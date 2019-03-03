package routes

import (
	"auth/controllers"
	"expense/utils/auth"
	"net/http"

	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {

	r := mux.NewRouter().StrictSlash(true)
	r.Use(commonMiddleware)
	r.Use(auth.JwtVerify)

	r.HandleFunc("/api", controllers.TestAPI).Methods("GET")
	r.HandleFunc("/api/register", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/api/login", controllers.Login).Methods("POST")
	r.HandleFunc("/api/user", controllers.FetchUsers).Methods("GET")

	r.HandleFunc("/api/user/{id}", controllers.GetUser).Methods("GET")
	r.HandleFunc("/api/user/{id}", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods("DELETE")
	return r
}

//set content-type
func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
		next.ServeHTTP(w, r)
	})
}
