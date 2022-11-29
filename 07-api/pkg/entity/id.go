package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(idValue string) (ID, error) {
	id, err := uuid.Parse(idValue)
	return ID(id), err
}
