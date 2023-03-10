package dictionary

import (
	"math/rand"
	"time"

	lorem "github.com/drhodes/golorem"
	"github.com/google/uuid"
)

func choiceRand(dict []string) string {
	if len(dict) < 1 {
		panic("нету смысла выполнять операция для одного слова...")
	}

	rand.Seed(time.Now().UnixNano())

	index := rand.Intn(len(dict))

	return dict[index]
}

func NewId() string {
	return uuid.New().String()
}

func NewTitle() string {
	titles := []string{
		"10 лучших способов", "Как стать профессионалом", "Секреты успеха", "7 шагов к цели",
		"Интересные факты о", "Что нового происходит в мире?", "Как быть продуктивным?",
		"Стоит ли учиться на программиста в 2023?", "Как нравиться девушкам?", "Почему я?",
		"Какие книги должен прочитать каждый программист", "Женщина в айти - к беде", "В чем смысл?",
	}

	return choiceRand(titles)
}

func NewContent() string {
	return lorem.Paragraph(5, 15)
}
