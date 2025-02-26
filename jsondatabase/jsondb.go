package jsondatabase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type jsondb struct {
	FilePath string
}

func NewJsonDB(filePath string) *jsondb {
	return &jsondb{FilePath: filePath}
}

func (db *jsondb) Save(dto ToDoTaskDTO) {
	tasksInDb := JsonToCollection()
	task := ToDoTask{Id: tasksInDb[len(tasksInDb)-1].Id + 1, Title: dto.Title, Desctiption: dto.Desctiption}

	tasksInDb = append(tasksInDb, task)
	CollectionToJson(tasksInDb)
}

// func (db *jsondb) GetById(id int) (*ToDoTask, error) {

// }

// func (db *jsondb) RemoveById(id int) error {

// }

func JsonToCollection() []ToDoTask {
	jsonFile, err := os.Open("jsondatabase/db.json")
	Check(err)
	defer jsonFile.Close()

	bytevalue, err := ioutil.ReadAll(jsonFile)
	Check(err)

	var tasks []ToDoTask
	json.Unmarshal(bytevalue, &tasks)

	return tasks
}

func CollectionToJson(tasks []ToDoTask) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	err = ioutil.WriteFile("jsondatabase/db.json", data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write to db.json: %w", err)
	}

	fmt.Println("Successfully updated db.json")
	return nil
}

func Check(err error) error {
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
