package services

import (
	db "github.com/hamza-sharif/homeopathic-doctor-assistant/db"
)

// Service initializes our database instance.
type Service struct {
	db db.Client
}

func NewService(db db.Client) *Service {
	m := &Service{
		db: db,
	}
	return m
}
