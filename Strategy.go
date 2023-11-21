package main

import "fmt"

type EvictionAlgo interface {
	evict(c *Cache)
}

type lfu struct{}

func (l *lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu")
}

type fifo struct{}

func (f *fifo) evict(c *Cache) {
	fmt.Println("Evicting by FIFO")
}

type Cache struct {
	eviction    EvictionAlgo
	capacity    int
	maxCapacity int
	storage     map[string]string
}

func initCache(algo EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		eviction:    algo,
		capacity:    0,
		maxCapacity: 2,
		storage:     storage,
	}
}

func (C *Cache) setCache(algo EvictionAlgo) {
	C.eviction = algo
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		// c.eviction.evict(c)
		// c.capacity--
		c.evict()
	}
	c.capacity++
	c.storage[key] = value

}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.eviction.evict(c)
	c.capacity--
}

func main() {
	fifo := &fifo{}
	cache := initCache(fifo)
	cache.add("a", "4")
	cache.add("b", "6")
	cache.setCache(&lfu{})

	cache.add("e", "5")
}
