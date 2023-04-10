package student

import (
	"errors"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func Put(v map[string]*Student, value *Student) (int, error) {
	v[value.Name] = value
	if v[value.Name] == nil {
		return -1, errors.New("Ошибка добавления данных.\n")
	} else {
		return 0, nil
	}
}

func Get(v map[string]*Student, name string) (*Student, error) {
	if v[name] == nil {
		return nil, errors.New("Имя студента не найдено.\n")
	} else {
		return v[name], nil
	}
}
