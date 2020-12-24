package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)

	}
	return router
}

var routes = Routes{
	Route{
		"CuentaList",
		"GET",
		"/cuentas",
		CuentaList,
	},
	Route{
		"CuentaShow",
		"GET",
		"/cuentas/{id}",
		CuentaShow,
	},
	Route{
		"CuentaAdd",
		"POST",
		"/cuenta",
		CuentaAdd,
	},
	Route{
		"CuentaUpdate",
		"PUt",
		"/cuenta/{id}",
		CuentaUpdate,
	},
	Route{
		"CuentaRemove",
		"DELETE",
		"/cuenta/{id}",
		CuentaRemove,
	},
}
