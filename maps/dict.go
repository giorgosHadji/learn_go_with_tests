package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
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

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, new_val string) error {
	_, ok := d.Search(key)
	if ok == ErrNotFound {
		return ErrWordDoesNotExist
	}
	d[key] = new_val
	return nil
}
func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, key)
		return nil
	default:
		return nil
	}
}
