package maps

import (
	"errors"
)

type (
	Dictionary    map[string]string
	DictionaryErr string
)

const (
	ErrNotFound  = DictionaryErr("could not find the word you were looking for")
	ErrWordExist = DictionaryErr("cannot add word because it already exists")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExist
	default:
		return err
	}
	return nil
}
