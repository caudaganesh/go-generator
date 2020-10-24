package structtype

type PropToCamel map[string]string

func (p PropToCamel) GetValue(key string) string {
	return p[key]
}
