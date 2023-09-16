package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("word is not in the dictionary")

func (d Dictionary) Search(searchWord string) (string, error) {

	if _, ok := d[searchWord]; ok {
		return d[searchWord], nil
	} else {
		return "", ErrNotFound
	}
}
