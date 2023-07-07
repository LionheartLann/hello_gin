package cache

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

func GetFromCache(c Cache, key string) (interface{}, bool) {
	if value, ok := c.Get(key); ok {
		return value, true
	}
	return nil, false
}
