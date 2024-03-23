package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

type Service interface {
	Strat(id int) bool
	Stop() bool
}

type ServiceManager struct {
	Services []Service
}

type Service1 struct {
	id int
}

type Service2 struct {
	id int
}

func (s *Service1) Strat(id int) bool {
	s.id = id
	fmt.Printf("Start servise ID=%d", s.id)
	return true
}

func (s *Service2) Strat(id int) bool {
	s.id = id
	fmt.Printf("Start servise ID=%d", s.id)
	return true
}

func (s Service1) Stop() bool {
	return true

}

func (s Service2) Stop() bool {
	return true
}

func NewServiceManager() *ServiceManager {
	manager := new(ServiceManager)
	manager.Services = append(manager.Services, &Service1{}, &Service2{})
	return manager
}

func (sm ServiceManager) StartAll() {
	for i, serv := range sm.Services {
		serv.Strat(i)
	}
}
