package pokecache
import (
	"fmt"
	"sync"
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

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
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
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
func TestOverwriteKey(t *testing.T) {
	cache := NewCache(1 * time.Second)

	cache.Add("key1", []byte("first"))
	cache.Add("key1", []byte("second"))

	val, ok := cache.Get("key1")
	if !ok {
		t.Fatalf("expected key to exist")
	}

	if string(val) != "second" {
		t.Fatalf("expected value to be overwritten")
	}
}
func TestGetMissingKey(t *testing.T) {
	cache := NewCache(1 * time.Second)

	_, ok := cache.Get("does-not-exist")
	if ok {
		t.Fatalf("expected key to not exist")
	}
}
func TestMultipleKeys(t *testing.T) {
	cache := NewCache(1 * time.Second)

	cache.Add("a", []byte("1"))
	cache.Add("b", []byte("2"))
	cache.Add("c", []byte("3"))

	tests := map[string]string{
		"a": "1",
		"b": "2",
		"c": "3",
	}

	for k, expected := range tests {
		val, ok := cache.Get(k)
		if !ok {
			t.Fatalf("expected key %s to exist", k)
		}
		if string(val) != expected {
			t.Fatalf("unexpected value for key %s", k)
		}
	}
}
func TestPartialExpiration(t *testing.T) {
	interval := 20 * time.Millisecond
	cache := NewCache(interval)

	cache.Add("old", []byte("data"))

	time.Sleep(interval + 5*time.Millisecond)

	cache.Add("new", []byte("data2"))

	time.Sleep(interval / 2)

	_, okOld := cache.Get("old")
	if okOld {
		t.Fatalf("expected old key to be expired")
	}

	_, okNew := cache.Get("new")
	if !okNew {
		t.Fatalf("expected new key to still exist")
	}
}
func TestConcurrentAccess(t *testing.T) {
	cache := NewCache(1 * time.Second)

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i)
			cache.Add(key, []byte("value"))
			cache.Get(key)
		}(i)
	}

	wg.Wait()
}
