package infra

type KeyValueStore interface {
	Set(key string, timeWindow, requestCount int64)
	Get(key string) (int64, int64)
}
