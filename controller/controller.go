package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"preloading/test/service"
	"preloading/test/student"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type StudentController struct {
	StudentService *service.StudentService
}

func New(service *service.StudentService) *StudentController {
	return &StudentController{
		StudentService: service,
	}
}

func (c *StudentController) HandleRequest() {
	log.Println("Starting development server at http://127.0.0.1:8080/")
	log.Println("Quit the server with CONTROL-C.")
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/student", c.CreateNewStudent).Methods("POST")
	myRouter.HandleFunc("/student", c.GetAllStudents).Methods("GET")
	myRouter.HandleFunc("/student/{id}", c.GetSingleStudent).Methods("GET")
	myRouter.HandleFunc("/student/{id}", c.UpdateSingleStudent).Methods("PUT")
	myRouter.HandleFunc("/student/{id}", c.DeleteStudent).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func (c *StudentController) CreateNewStudent(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student *student.Student
	json.Unmarshal(reqBody, &student)
	er := c.StudentService.Add(student)
	if er != nil {
		json.NewEncoder(w).Encode(er.Error())
		fmt.Println("error", er)
		return
	}
	fmt.Println("Endpoint Hit: Creating New Student")
	json.NewEncoder(w).Encode(student.ID)
}

func (c *StudentController) GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students := []*student.Student{}
	er := c.StudentService.GetAll(&students)
	if er != nil {
		json.NewEncoder(w).Encode(er)
		fmt.Println("error", er)
		return
	}
	fmt.Println("Endpoint Hit: returnAllStudent")
	json.NewEncoder(w).Encode(students)
}

func (c *StudentController) GetSingleStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	student := student.Student{}
	ID, _ := uuid.FromString(key)
	er := c.StudentService.Get(&student, ID)
	if er != nil {
		json.NewEncoder(w).Encode(er)
		fmt.Println("error", er)
		return
	}

	fmt.Println("Endpoint Hit: Single Student:", key)
	json.NewEncoder(w).Encode(student)

}
func (c *StudentController) UpdateSingleStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var student *student.Student
	json.Unmarshal(reqBody, &student)
	fmt.Println(student)
	ID, err := uuid.FromString(key)
	if err != nil {
		fmt.Println(err)
	}
	er := c.StudentService.Update(student, ID)
	if er != nil {
		json.NewEncoder(w).Encode(er)
		fmt.Println("error", er)
		return
	}

	fmt.Println("Endpoint Hit: Update Student:", key)
	json.NewEncoder(w).Encode(student)
}

func (c *StudentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	student := student.Student{}
	ID, err := uuid.FromString(key)
	student.ID = ID
	if err != nil {
		fmt.Println(err)
	}
	er := c.StudentService.Delete(&student, ID)
	if er != nil {
		json.NewEncoder(w).Encode(er)
		fmt.Println("error", er)
		return
	}

	fmt.Println("Endpoint Hit: Delete Student:", key)
	json.NewEncoder(w).Encode(student)
}
