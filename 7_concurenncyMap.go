package main

import "sync"

/*
Реализовать конкурентную запись данных в map
*/

/*
Для данной задачи можно использовать sync.Map из стандартной библиотеки
или реализовать конкурентную запись вручную
*/

type Cash struct {
	sync.RWMutex
	m map[string]int
}

// Доступ к данным по ключу
func (c *Cash) Load(key string) (int, bool) {
	c.RLock()
	defer c.RUnlock()
	val, ok := c.m[key]
	return val, ok
}

// Запись данных в кэш
func (c *Cash) Store(key string, value int) {
	c.Lock()
	defer c.Unlock()
	c.m[key] = value
}

// Удаление данных по ключу
func (c *Cash) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.m, key)
}
