package remessa

type Params struct {
	params map[string]string
}

func NewParams() Params {
	p := make(map[string]string)
	return Params{params: p}
}

func (p Params) Add(key, value string) Params {
	p.params[key] = value
	return p
}

func (p Params) Get(key string) string {
	return p.params[key]
}
