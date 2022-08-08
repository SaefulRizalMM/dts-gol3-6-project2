package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", showTasks).Methods("GET")
	router.HandleFunc("/api/task/{id}", showTaskById).Methods("GET")
	router.HandleFunc("/api/task", insertTask).Methods("POST")
	router.HandleFunc("/api/update/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/api/change-status/{id}", changeStatusTask).Methods("PUT")
	router.HandleFunc("/api/delete/{id}", deleteTask).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	// log.Fatal(http.ListenAndServe(":1234", router))
	log.Fatal(http.ListenAndServe(":1234",
		handlers.CORS(handlers.AllowedHeaders([]string{
			"X-Requested-With",
			"Content-Type",
			"Authorization"}),
			handlers.AllowedMethods([]string{
				"GET",
				"POST",
				"PUT",
				"DELETE",
				"HEAD",
				"OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(router)))

}

func showTasks(w http.ResponseWriter, r *http.Request) {
	var tasks Tasks
	var arr_task []Tasks
	// var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select * from task")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&tasks.Id, &tasks.Description, &tasks.Deadline, &tasks.Employee, &tasks.Status); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_task = append(arr_task, tasks)
		}
	}

	// response.Status = 1
	// response.Message = "Success"
	// response.Data = arr_task

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(arr_task)

}

func showTaskById(w http.ResponseWriter, r *http.Request) {
	var tasks Tasks
	var arr_task []Tasks
	var response Response
	params := mux.Vars(r)

	db := connect()
	defer db.Close()

	rows, err := db.Query("Select * from task WHERE id = ?", params["id"])
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&tasks.Id, &tasks.Description, &tasks.Deadline, &tasks.Employee, &tasks.Status); err != nil {
			log.Fatal(err.Error())

		} else {
			arr_task = append(arr_task, tasks)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arr_task

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func insertTask(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	description := keyVal["description"]
	employee := keyVal["employee"]
	deadline := keyVal["deadline"]
	status := keyVal["status"]

	_, err = db.Exec("INSERT INTO task (description, employee, deadline, status) VALUES(?, ?, ?, ?)", description, employee, deadline, status)

	if err != nil {
		log.Print(err)
		return
	}
	response.Status = 200
	response.Message = "Insert data successfully"
	fmt.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var response Response
	params := mux.Vars(r)

	db := connect()
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	// id := keyVal["id"]
	description := keyVal["description"]
	employee := keyVal["employee"]
	deadline := keyVal["deadline"]
	// status := keyVal["status"]

	_, err = db.Exec("UPDATE task SET description=?, employee=?, deadline=? WHERE id=?", description, employee, deadline, params["id"])

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Update data successfully"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func changeStatusTask(w http.ResponseWriter, r *http.Request) {
	var response Response
	params := mux.Vars(r)

	db := connect()
	defer db.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	// id := keyVal["id"]
	status := keyVal["status"]

	_, err = db.Exec("UPDATE task SET status=? WHERE id=?", status, params["id"])

	if err != nil {
		log.Print(err)
	}

	response.Status = 200
	response.Message = "Change status data successfully"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	var response Response
	params := mux.Vars(r)

	db := connect()
	defer db.Close()

	if _, err := db.Exec("DELETE FROM task WHERE id=?", params["id"]); err != nil {
		panic(err)
	}

	response.Status = 200
	response.Message = "Delete data successfully"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
