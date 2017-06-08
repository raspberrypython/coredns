package cache

import "testing"

func TestCacheInsertAndLookup(t *testing.T) {
	c := New(4)
	c.Insert("test", 1)

	if _, found := c.Lookup("test"); !found {
		t.Fatal("Failed to find inserted record")
	}
}

func TestCacheLen(t *testing.T) {
	c := New(4)

	c.Insert("test", 1)
	if l := c.Len(); l != 1 {
		t.Fatalf("Cache size should %d, got %d", 1, l)
	}

	c.Insert("test", 1)
	if l := c.Len(); l != 1 {
		t.Fatalf("Cache size should %d, got %d", 1, l)
	}

	c.Insert("test2", 2)
	if l := c.Len(); l != 2 {
		t.Fatalf("Cache size should %d, got %d", 2, l)
	}
}

func TestCacheEvict(t *testing.T) {
	c := New(1)
	c.Insert("test1", 1)
	c.Insert("test2", 2)
	// test1 should be gone

	if _, found := c.Lookup("test1"); found {
		t.Fatal("Found item that should have been evicted")
	}
}

func TestCacheLenEvict(t *testing.T) {
	c := New(4)
	c.Insert("test1", 1)
	c.Insert("test2", 1)
	c.Insert("test3", 1)
	c.Insert("test4", 1)

	if l := c.Len(); l != 4 {
		t.Fatalf("Cache size should %d, got %d", 4, l)
	}

	// This should evict one element
	c.Insert("test5", 1)
	if l := c.Len(); l != 4 {
		t.Fatalf("Cache size should %d, got %d", 4, l)
	}
}
