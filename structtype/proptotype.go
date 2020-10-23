package structtype

type PropToType map[string]string

func (p PropToType) GetValue(key string) string {
	return p[key]
}
