package main

import (
    "net/http"
    "to-go/controllers"
    "github.com/gorilla/handlers"
    "github.com/gorilla/mux"
)

type Route struct {
    Method      string
    Pattern     string
    Name        string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

//func NewRouter() *mux.Router {
func NewRouter() http.Handler {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),
        handlers.AllowedMethods([]string{"OPTIONS", "DELETE", "GET", "HEAD", "POST", "PUT"}),
        handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "X-AUTH-TOKEN"}),
    )(router)
}

var routes = Routes{
    Route{ "GET", "/", "Index", controllers.Index, },

    Route{ "GET", "/todos", "TodoIndex", controllers.TodosIndex, },
    Route{ "GET", "/todos/{todoId}", "TodoShow", controllers.TodoShow, },
    Route{ "POST", "/todos", "TodoCreate", controllers.TodoCreate, },
    Route{ "PUT", "/todos/{todoId}", "TodoUpdate", controllers.TodoUpdate, },
    Route{ "DELETE", "/todos/{todoId}", "TodoDelete", controllers.TodoDelete, },

    Route{ "POST", "/login", "Login", controllers.Login, },
    Route{ "POST", "/signup", "UserCreate", controllers.UserCreate, },
    Route{ "PUT", "/users/{userId}", "UserUpdate", controllers.UserUpdate, },
    Route{ "DELETE", "/users/{userId}", "UserDelete", controllers.UserDelete, },
}

