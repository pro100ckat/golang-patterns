package singleton

import (
	"fmt"
	"sync"
)

type singleInstance struct{}

var (
	// once sync.Once -- или так
	mu       = &sync.Mutex{}
	instance *singleInstance
)

// GetInstance Одиночка определяет статический метод GetInstance, который возвращает единственный экземпляр своего класса.
//  Конструктор одиночки должен быть скрыт от клиентов.
//  Вызов метода GetInstance должен стать единственным способом получить объект этого класса.
func GetInstance() *singleInstance {
	/*
		once.Do(func() {
				instance = &singleInstance{}
			})
	*/
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		fmt.Println("Creating singleton instance now.")
		instance = &singleInstance{}
	} else {
		fmt.Println("Single instance already created.")
	}
	return instance
}
