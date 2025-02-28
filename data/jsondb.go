package data

type Database struct {
	FilePath string
}

func NewDatabase(filePath string) *Database {
	return &Database{FilePath: filePath}
}

func (db *Database) GetAll() {

}

func (db *Database) Save(dto ToDoTaskDTO) {
}

// func (db *Database) GetById(id int) (*ToDoTask, error) {

// }

// func (db *Database) RemoveById(id int) error {

// }
