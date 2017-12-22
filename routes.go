package main

import (
    "net/http"
    "to-go/controllers"
    "github.com/gorilla/mux"
)

type Route struct {
    Method      string
    Pattern     string
    Name        string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{ "GET", "/", "Index", controllers.Index, },

    Route{ "GET", "/todos", "TodoIndex", controllers.TodosIndex, },
    Route{ "GET", "/todos/{todoId}", "TodoShow", controllers.TodoShow, },
    Route{ "POST", "/todos", "TodoCreate", controllers.TodoCreate, },
    Route{ "PUT", "/todos/{todoId}", "TodoUpdate", controllers.TodoUpdate, },
    Route{ "DELETE", "/todos/{todoId}", "TodoDelete", controllers.TodoDelete, },

    Route{ "POST", "/login", "Login", controllers.Login, },
    // get this one working and validate email format and password length
    Route{ "POST", "/users", "UserCreate", controllers.UserCreate, },
    // secure these two endpoints with auth token
    Route{ "PUT", "/users/{userId}", "UserUpdate", controllers.UserUpdate, },
    Route{ "DELETE", "/users/{userId}", "UserDelete", controllers.UserDelete, },
}

