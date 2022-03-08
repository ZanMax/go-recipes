package main

import "fmt"

type Set struct {
	container map[string]struct{}
}

func newSet() *Set {
	return &Set{
		container: make(map[string]struct{}),
	}
}

func (c *Set) Add(key string) {
	c.container[key] = struct{}{}
}

func (c *Set) Remove(key string) error {
	_, exists := c.container[key]
	if !exists {
		return fmt.Errorf("Key not found")
	}
	delete(c.container, key)
	return nil
}

func (c *Set) Exists(key string) bool {
	_, exists := c.container[key]
	return exists
}

func (c *Set) Size() int {
	return len(c.container)
}
