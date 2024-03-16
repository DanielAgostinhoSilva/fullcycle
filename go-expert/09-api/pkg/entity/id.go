package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewId() ID {
	return ID(uuid.New())
}

func ParseID(id string) (ID, error) {
	uuid, err := uuid.Parse(id)
	return ID(uuid), err
}
