package entity

import "github.com/google/uuid"

// layer que pode ser compartilhada na aplicacao
type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error){
	id, err := uuid.Parse(s)
	return id, err
}