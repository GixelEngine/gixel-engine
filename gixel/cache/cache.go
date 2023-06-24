package cache

import (
	"embed"
)

type CacheOptions struct {
	Key     string
	Unique  bool
	Persist bool
	NoCache bool
}

type cacheRecord[T any] struct {
	persist bool
	data    *T
}

type GxlCache[T any] struct {
	cache map[string]*cacheRecord[T]
	FS    *embed.FS
}

func NewCache[T any](fs *embed.FS) *GxlCache[T] {
	return &GxlCache[T]{
		cache: make(map[string]*cacheRecord[T]),
		FS:    fs,
	}
}

func (gc *GxlCache[T]) Add(data *T, key string, persist bool) *T {
	if data == nil || key == "" {
		return nil
	}

	gc.cache[key] = &cacheRecord[T]{persist: persist, data: data}

	return data
}

func (gc *GxlCache[T]) Get(key string) *T {
	record, ok := gc.cache[key]
	if !ok || record == nil {
		return nil
	}

	return record.data
}

func (gc *GxlCache[T]) Clear() {
	for _, record := range gc.cache {
		if !record.persist {
			record = nil
		}
	}
}
