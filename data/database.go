package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var filePath string = "data/database.json"

type Database struct {
	FilePath string
}

func NewDatabase() *Database {
	return &Database{FilePath: filePath}
}

func (db *Database) GetAll() []ToDoTask {
	return jsonToCollection()
}

func (db *Database) Save(dto ToDoTaskDTO) {
	tasks := jsonToCollection()
	newId := tasks[len(tasks)-1].Id + 1
	newTask := ToDoTask{Id: newId, Title: dto.Title, Desctiption: dto.Desctiption}

	tasks = append(tasks, newTask)
	collectionToJson(tasks)
}

func (db *Database) GetById(id int) (*ToDoTask, error) {
	tasks := jsonToCollection()

	for _, task := range tasks {
		if task.Id == id {
			return &task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (db *Database) RemoveById(id int) error {
	tasks := jsonToCollection()
	var index int = -1
	for i, task := range tasks {
		if task.Id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("task not found")
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	collectionToJson(tasks)
	return nil
}

func jsonToCollection() []ToDoTask {
	jsonData, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonData.Close()

	byteData, _ := ioutil.ReadAll(jsonData)
	var tasks []ToDoTask
	json.Unmarshal(byteData, &tasks)

	return tasks
}

func collectionToJson(tasks []ToDoTask) error {
	byteData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to convert to JSON: %w", err)
	}

	err = ioutil.WriteFile(filePath, byteData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
