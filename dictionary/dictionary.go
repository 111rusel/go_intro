package dictionary

import "errors"

func Search(s map[string]string, x string) string {
	return s[x]
}

type Dictionary map[string]string

func (d Dictionary) Search(s string) string {
	return d[s]
}

func (d Dictionary) Add(s, c string) error {
	_, ok := d[s]
	if ok {
		return errors.New("error")
	}
	d[s] = c
	return nil
}

func (d Dictionary) Update(s, c string) {
	d[s] = c
}

func (d Dictionary) Delete(s string) {
	delete(d, s)
}
