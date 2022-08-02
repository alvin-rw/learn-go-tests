package dictionary

type Dictionary map[string]string

const (
	ErrNotFound     = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExist = DictionaryErr("the word is already exist")
)

type DictionaryErr string

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
