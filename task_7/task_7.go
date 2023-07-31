package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type SyncMap[K comparable, V any] struct {
	// Добавляем мютекс для защиты от гонки
	sync.Mutex
	m map[K]V
}

func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{m: make(map[K]V)}
}

func (m *SyncMap[K, V]) Get(key K) V {
	// Захватываем мьютекс
	m.Lock()
	// Отложенно освобождаем мьютекс
	defer m.Unlock()

	return m.m[key]
}

func (m *SyncMap[K, V]) Set(key K, value V) {
	// Захватываем мьютекс
	m.Lock()
	// Отложенно освобождаем мьютекс
	defer m.Unlock()

	m.m[key] = value
}

func (m *SyncMap[K, V]) Range(f func(key K, value V)) {
	// Захватываем мьютекс
	m.Lock()
	// Отложенно освобождаем мьютекс
	defer m.Unlock()

	for key, value := range m.m {
		f(key, value)
	}
}

func main() {
	syncMap := NewSyncMap[int, int]()

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)

		// Генерируем случайные числа и записывает их в карту
		go func(wg *sync.WaitGroup) {
			defer wg.Done()

			syncMap.Set(rand.Intn(100), rand.Int())
		}(&wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	// Выводим содержимое карты в консоль
	syncMap.Range(func(key int, value int) {
		fmt.Println(key, value)
	})
}
