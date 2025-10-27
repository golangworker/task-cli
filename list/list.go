package list

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"task-cli/task"
)

// json файл с сохраненными задачами
const FILE_NAME = "file.json"

type tasks map[string]task.Task

// структура для хранения списка задач
type list struct {
	Tasks tasks
}

// функция инициализации и загрузка задач из файла
func Init() (list, error) {

	newTasks := make(tasks)
	_, err := os.Stat(FILE_NAME)

	// если ошибка, не связанна с существованием файла
	if err == nil {
		// если файл существует, загружаем все данные в мапу
		data, err := os.ReadFile(FILE_NAME)
		if err != nil {
			return list{}, err
		}

		if err := json.Unmarshal(data, &newTasks); err != nil {
			return list{}, err
		}

	} else if !errors.Is(err, os.ErrNotExist) {
		return list{}, err

	}
	return list{
		Tasks: newTasks,
	}, nil
}

// метод для добавления задачи в список
func (l *list) Add(t task.Task) {
	l.Tasks[t.ID] = t
	fmt.Println("Task added successfully!")
}

// метод для удаления задачи из списка
func (l *list) Delete(ID string) error {

	if isTask(l.Tasks, ID) {
		delete(l.Tasks, ID)
		return nil
	} else {
		return fmt.Errorf("task with ID %s not found", ID)
	}
}

// метод для сохранения списка задач в файл
func (l *list) Save() error {

	bytes, err := json.MarshalIndent(l.Tasks, "", "  ")
	if err != nil {
		return err
	}
	os.WriteFile(FILE_NAME, bytes, 0644)
	fmt.Println("Tasks saved successfully!")
	return nil
}

// локальный метод для проверки наличия задачи в списке
func isTask(t tasks, ID string) bool {
	if _, ok := t[ID]; ok {
		return true
	}
	return false
}
