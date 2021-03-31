package maps

type Dictionary map[string]string

const (
	ErrWordNotFound     = DictionaryError("this word is not in the dictionary")
	ErrWordExists       = DictionaryError("word already exists")
	ErrWordDoesNotExist = DictionaryError("word does not exist")
)

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrWordNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil

}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {
	//lol
}
