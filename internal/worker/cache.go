package worker

type CacheLayer struct {
	data map[string]interface{}
}

var _dataCache = CacheLayer{data: make(map[string]interface{})}

func CacheIt(uuid string, data interface{}) {
	_dataCache.data[uuid] = data
}

func BringIt(uuid string) interface{} {
	return _dataCache.data[uuid]
}
