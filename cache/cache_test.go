package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	val := cache.Get("a")
	if val != nil {
		t.Error("Getting a non-existent key should have returned an empty val")
	}

	cache.Set("a", "Item A", 100*time.Millisecond)
	val = cache.Get("a")
	if val != "Item A" {
		t.Errorf("Expected 'Item A', got %s", val)
	}

	cache.Set("b", "Item B", 100*time.Millisecond)
	val = cache.Get("b")
	if val != "Item B" {
		t.Errorf("Expected 'Item B', got %s", val)
	}

	cache.Set("a", "Item A Prime", 100*time.Millisecond)
	val = cache.Get("a")
	if val != "Item A Prime" {
		t.Errorf("Expected 'Item A Prime', got %s", val)
	}

	time.Sleep(100 * time.Millisecond)

	val = cache.Get("a")
	if val != nil {
		t.Error("Getting expired key 'a' should have returned an empty val")
	}

	val = cache.Get("b")
	if val != nil {
		t.Error("Getting expired key 'b' should have returned an empty val")
	}
}
