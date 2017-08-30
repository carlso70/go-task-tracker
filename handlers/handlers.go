package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/carlso70/go-todo/repo"
	"github.com/gorilla/mux"
)

// Index welcome page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

// TodoIndex index todo
func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(repo.GetTodos()); err != nil {
		panic(err)
	}
}

// TodoShow show todo
func TodoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["todoId"])
	if err != nil {
		panic(err)
	} else {
		name := repo.RepoFindTodo(int(todoId)).Name
		fmt.Fprintf(w, "TodoID:%d Todo:%s\n", todoId, name)
	}
}

// TodoCreate create new todo
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo repo.Todo
	// LimitReader is good for preventing someone from sending 50 GBs of json
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // Unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := repo.RepoCreateTodo(todo)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated) // return 201 status code
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
