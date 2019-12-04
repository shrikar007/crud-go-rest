package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"google.golang.org/genproto/googleapis/type/date"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Expense struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Amount      float64   `json:"amount"`
	CreatedOn   date.Date `json:"created_on" `
	UpdatedOn   date.Date `json:"updated_on"`
}

type Expenses []Expense

var expenses Expenses

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/expenses", func(r chi.Router) {
		r.Post("/", CreateExpense)
		r.Get("/", ListAllExpense)
		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", ListOneExpense)
			r.Put("/", UpdateExpense)
			r.Delete("/", DeleteExpense)
		})
	})

	log.Fatal(http.ListenAndServe(":8082", r))
}
var  index int
func CreateExpense(writer http.ResponseWriter, request *http.Request) {

	b, err := ioutil.ReadAll(request.Body)
	if err != nil {
		http.Error(writer, "unable to read request body", 500)
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		http.Error(writer, "unable to parse json request body", 422)
	}
	expense := new(Expense)
	if val, ok := data["description"].(string); ok {
		expense.Description = val
	}

	if val, ok := data["type"].(string); ok {
		expense.Type = val
	}
	if val, ok := data["amount"].(float64); ok {
		expense.Amount = val
	}
	index=index+1
	expense.Id=index
	expenses = append(expenses, *expense)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)

	_, _ = fmt.Fprintln(writer, `{"success": true}`)
}

func ListOneExpense(writer http.ResponseWriter, request *http.Request) {
	flag:=false

	for _,expense:=range expenses{
		if (strconv.Itoa(expense.Id))==(chi.URLParam(request, "id")){
			encoder := json.NewEncoder(writer)
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			encoder.Encode(expense)
			flag=true
			return
		}
	}
	if flag==false{
	    writer.WriteHeader(http.StatusNotFound)
	}
}

func ListAllExpense(writer http.ResponseWriter, request *http.Request) {

		encoder := json.NewEncoder(writer)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	encoder.Encode(expenses)
	fmt.Println(expenses)
}

func UpdateExpense(writer http.ResponseWriter, request *http.Request) {

	flag:=false

	reqBody, _ := ioutil.ReadAll(request.Body)
	for updateindex,expense:=range expenses{
		if (strconv.Itoa(expense.Id))==(chi.URLParam(request, "id")) {
			json.Unmarshal(reqBody, &expense)
			expenses[updateindex] = expense
			json.NewEncoder(writer).Encode(expense)
			_, _ = fmt.Fprintln(writer, `{"success": true}`)
			flag=true
		}
	}
	if flag==false{
		writer.WriteHeader(http.StatusNotFound)
	}
}

func DeleteExpense(writer http.ResponseWriter, request *http.Request) {


	flag:=false
	for deleteindex,expense:=range expenses{
		if strconv.Itoa(expense.Id)==chi.URLParam(request, "id"){
			expenses=append(expenses[:deleteindex],expenses[deleteindex+1:]...)
			_, _ = fmt.Fprintln(writer, `{"success": true}`)
			flag=true
			return
		}
	}
	if flag==false{
		writer.WriteHeader(http.StatusNotFound)
	}

}
