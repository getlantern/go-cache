package cache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewCache()

	val, found := cache.Get("a")
	if found {
		t.Error("Getting a non-existent key should have returned not found")
	}
	if val != nil {
		t.Error("Getting a non-existent key should have returned a nil value")
	}

	cache.Set("a", "Item A", 100*time.Millisecond)
	val, _ = cache.Get("a")
	if val != "Item A" {
		t.Errorf("Expected 'Item A', got %s", val)
	}

	cache.Set("b", "Item B", 100*time.Millisecond)
	val, _ = cache.Get("b")
	if val != "Item B" {
		t.Errorf("Expected 'Item B', got %s", val)
	}

	cache.Set("c", nil, 100*time.Millisecond)
	val, found = cache.Get("c")
	if val != nil {
		t.Errorf("Expected nil, got %s", val)
	}
	if !found {
		t.Errorf("A nil value should still have returned found")
	}

	cache.Set("a", "Item A Prime", 100*time.Millisecond)
	val, _ = cache.Get("a")
	if val != "Item A Prime" {
		t.Errorf("Expected 'Item A Prime', got %s", val)
	}

	time.Sleep(100 * time.Millisecond)

	val, found = cache.Get("a")
	if found {
		t.Error("Getting expired key 'a' should have returned not found")
	}
	if val != nil {
		t.Error("Getting expired key 'a' should have returned a nil value")
	}

	val, found = cache.Get("b")
	if found {
		t.Error("Getting expired key 'b' should have returned not found")
	}
	if val != nil {
		t.Error("Getting expired key 'b' should have returned a nil value")
	}

	val, found = cache.Get("c")
	if found {
		t.Error("Getting expired key 'c' should have returned not found")
	}
	if val != nil {
		t.Error("Getting expired key 'c' should have returned a nil value")
	}
}
