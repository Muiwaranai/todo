package endpoints

import (
	"fmt"
	"net/http"
	"todo/jsondatabase"
)

func Addtask(w http.ResponseWriter, r *http.Request) {
	//databse := jsondatabase.NewJsonDB("jsondatabase/db.json")
	dto := jsondatabase.ToDoTaskDTO{}
	//println(dto.Id)
	fmt.Fprintf(w, "hello from add")
}
