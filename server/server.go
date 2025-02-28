package server

import "todo/data"

type Server struct {
	db data.Database
}

func NewServerL(db data.Database) *Server {
	return &Server{
		db: db,
	}
}
