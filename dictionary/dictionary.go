package dictionary

func Search(s map[string]string, x string) string {
	return s[x]
}

type Dictionary map[string]string

func (d Dictionary) Search(s string) string {
	return d[s]
}
