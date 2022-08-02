package dictionary

import "errors"

var (
	ErrNotFound     = errors.New("could not find the word you were looking for")
	ErrAlreadyExist = errors.New("the word is already exist")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAlreadyExist
	default:
		return err
	}

	return nil
}
