package dictionary

import (
	lorem "github.com/drhodes/golorem"
	"github.com/google/uuid"
)

func NewID() string {
	return uuid.New().String()
}

func NewTitle() string {
	return lorem.Word(2, 7)
}

func NewContent() string {
	return lorem.Paragraph(5, 15)
}
