package task

import (
	"math/rand"
	"strconv"
	"time"
)

// структура Задачи
type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// setter для Задачи
func Init(desctiption string) Task {
	return Task{
		ID:          strconv.Itoa(rand.Int()),
		Description: desctiption,
		Status:      "in-progress",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}
}

// метод, который устанавливает статус Задачи как "done"
func (t *Task) MarkItDone() {
	t.Status = "done"
	t.UpdatedAt = time.Now().Format(time.RFC3339)
}
