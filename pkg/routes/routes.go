package routes

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"user-authentication-with-go/pkg/controllers"
	"user-authentication-with-go/pkg/middleware"
)

func SetupRoutes(db *gorm.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/register", controllers.Register(db)).Methods("POST")
	r.HandleFunc("/login", controllers.Login(db)).Methods("POST")

	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.IsAuthorized("admin"))
	adminRouter.HandleFunc("/dashboard", controllers.AdminDashboard())
	adminRouter.HandleFunc("/books", controllers.CreateBook(db)).Methods("POST")
	adminRouter.HandleFunc("/books/{id}", controllers.UpdateBook(db)).Methods("PUT")
	adminRouter.HandleFunc("/books/{id}", controllers.DeleteBook(db)).Methods("DELETE")
	adminRouter.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	adminRouter.HandleFunc("/books/{id}", controllers.GetBook(db)).Methods("GET")

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.Use(middleware.IsAuthorized("normal"))
	userRouter.HandleFunc("/profile", controllers.UserProfile())
	userRouter.HandleFunc("/books", controllers.GetBooks(db)).Methods("GET")
	userRouter.HandleFunc("/books/{id}", controllers.GetBook(db)).Methods("GET")

	return r
}
