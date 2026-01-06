package storage

import (
	"errors"
	"todolist/todo"
)

type Storage struct {
	tasks  []todo.Task
	lastID int
}

func New() *Storage {
	return &Storage{
		tasks:  make([]todo.Task, 0),
		lastID: 0,
	}
}

func (s *Storage) Add(title string) {
	s.lastID++

	task := todo.Task{
		ID:     s.lastID,
		Title:  title,
		Active: true,
	}

	s.tasks = append(s.tasks, task)
}

func (s *Storage) Delete(id int) error {
	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("задача с таким ID не найдена")
}

func (s *Storage) Complete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks[i].Active = false
			return nil
		}
	}
	return errors.New("задача с таким ID не найдена")
}

func (s *Storage) GetAll() []todo.Task {
	return s.tasks
}
