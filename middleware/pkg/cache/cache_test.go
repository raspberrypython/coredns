package cache

import "testing"

func TestCacheAddAndGet(t *testing.T) {
	c := New(4)
	c.Add("test", 1)

	if _, found := c.Get("test"); !found {
		t.Fatal("Failed to find inserted record")
	}
}

func TestCacheLen(t *testing.T) {
	c := New(4)

	c.Add("test", 1)
	if l := c.Len(); l != 1 {
		t.Fatalf("Cache size should %d, got %d", 1, l)
	}

	c.Add("test", 1)
	if l := c.Len(); l != 1 {
		t.Fatalf("Cache size should %d, got %d", 1, l)
	}

	c.Add("test2", 2)
	if l := c.Len(); l != 2 {
		t.Fatalf("Cache size should %d, got %d", 2, l)
	}
}

func TestCacheEvict(t *testing.T) {
	c := New(1)
	c.Add("test1", 1)
	c.Add("test2", 2)
	// test1 should be gone

	if _, found := c.Get("test1"); found {
		t.Fatal("Found item that should have been evicted")
	}
}

func TestCacheLenEvict(t *testing.T) {
	c := New(4)
	c.Add("test1", 1)
	c.Add("test2", 1)
	c.Add("test3", 1)
	c.Add("test4", 1)

	if l := c.Len(); l != 4 {
		t.Fatalf("Cache size should %d, got %d", 4, l)
	}

	// This should evict one element
	c.Add("test5", 1)
	if l := c.Len(); l != 4 {
		t.Fatalf("Cache size should %d, got %d", 4, l)
	}
}
