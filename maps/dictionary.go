package maps

import (
	"errors"
)

type (
	Dictionary    map[string]string
	DictionaryErr string
)

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrWordExist    = DictionaryErr("cannot add word because it already exists")
	ErrWordNotExist = DictionaryErr("cannot update word because it does not exist")
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

func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
