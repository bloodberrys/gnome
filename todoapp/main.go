package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Download gin by $go get -u github.com/gin-gonic/gin

// create a schema
type todo struct {
	Id          string `json:"id"`
	Item        string `json:"item"`
	IsCompleted bool   `json:"isCompleted"`
}

// create a list of object (not json)
var todos = []todo{
	{Id: "1", Item: "Clean Room", IsCompleted: false},
	{Id: "2", Item: "Read Book", IsCompleted: false},
	{Id: "3", Item: "Record Video", IsCompleted: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.POST("/todos", setTodo)
	router.Run("localhost:9090")
}

// accessing todo lists
func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func setTodo(context *gin.Context) {
	// Create a new Todo object
	var newTodo todo

	// receive the payload and error handling
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	// append the todos list with the new payload
	todos = append(todos, newTodo)

	// return status code 201 to requester (created)
	context.IndentedJSON(http.StatusCreated, newTodo)

}

// get list by id and error if any related to the schema
// with multiple type of return = *todo and error (will only result one of those)
func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.Id == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New("there are no data")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {

		// status not found is 404
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "the data not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)

}

// to change the isCompleted data
func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {

		// status not found is 404
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "the data not found"})
		return
	}

	todo.IsCompleted = !todo.IsCompleted

	context.IndentedJSON(http.StatusOK, todo)
}
