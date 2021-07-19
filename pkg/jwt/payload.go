package jwt

type Payload map[string]interface{}

func (i *Payload) Get(key string) interface{} {
	return (*i)[key]
}

func (i *Payload) Set(key string, value interface{}) {
	(*i)[key] = value
}
