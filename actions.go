package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Conexión MongoDB
func getSession() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return session
}

func responseCuenta(w http.ResponseWriter, status int, results Cuenta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

func responseCuentas(w http.ResponseWriter, status int, results []Cuenta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}

var collection = getSession().DB("cuenta_go").C("cuentas")

func CuentaList(w http.ResponseWriter, r *http.Request) {
	var results []Cuenta
	err := collection.Find(nil).Sort("-_id").All(&results)

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Resultados:", results)
	}

	responseCuentas(w, 200, results)

}

//Mostrar Cuentas
func CuentaShow(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cuenta_id := params["id"]

	if !bson.IsObjectIdHex(cuenta_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(cuenta_id)

	results := Cuenta{}
	err := collection.FindId(oid).One(&results)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseCuenta(w, 200, results)
}

//Agregar Cuenta
func CuentaAdd(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var cuenta_data Cuenta
	err := decoder.Decode(&cuenta_data)

	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	//Guardar en BD
	err = collection.Insert(cuenta_data)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	responseCuenta(w, 200, cuenta_data)
}

func CuentaUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cuenta_id := params["id"]

	if !bson.IsObjectIdHex(cuenta_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(cuenta_id)
	decoder := json.NewDecoder(r.Body)

	var cuenta_data Cuenta
	err := decoder.Decode(&cuenta_data)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
		return
	}

	defer r.Body.Close()

	document := bson.M{"_id": oid}
	change := bson.M{"$set": cuenta_data}
	err = collection.Update(document, change)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	responseCuenta(w, 200, cuenta_data)
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func (this *Message) setStatus(data string) {
	this.Status = data
}

func (this *Message) setMessage(data string) {
	this.Message = data
}

func CuentaRemove(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cuenta_id := params["id"]

	if !bson.IsObjectIdHex(cuenta_id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(cuenta_id)

	err := collection.RemoveId(oid)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	//results := Message{"success", "La cuenta con ID "+cuenta_id+" ha sido borrada correctamente"}
	message := new(Message)

	message.setStatus("success")
	message.setMessage("La cuenta con ID " + cuenta_id + " ha sido borrada correctamente")

	results := message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(results)
}
func GetCuentaAPI(w http.ResponseWriter, r *http.Request) {

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	json.NewEncoder(w).Encode("OKOK")
}
