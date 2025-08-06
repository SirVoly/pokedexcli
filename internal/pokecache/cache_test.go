package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		t.Run(c.key, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf(`
Test Failed:
	key %v not in cache
`, c.key)
				return
			}
			if string(val) != string(c.val) {
				t.Errorf(`
Test Failed:
	key: %v
	expected: %v
	actual: %v
`, c.key, c.val, val)
				return
			}
		})
	}
}

func TestNoGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
		},
		{
			key: "https://example.com/path",
		},
	}

	for _, c := range cases {
		t.Run(c.key, func(t *testing.T) {
			cache := NewCache(interval)
			val, ok := cache.Get(c.key)
			if ok {
				t.Errorf(`
Test Failed:
	key %v in cache with value %v
`, c.key, val)
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf(`
Test Failed:
	Expected to find the key
`)
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf(`
Test Failed:
	Expected not to find the key
`)
		return
	}
}
