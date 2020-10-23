package structtype

type PropToTag map[string]string

func (p PropToTag) GetValue(key string) string {
	return p[key]
}
