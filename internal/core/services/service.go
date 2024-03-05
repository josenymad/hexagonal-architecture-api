package services

import "fmt"

type Service struct {
}

func NewService() Service {

	return Service{}
}

func (service *Service) SendData(data interface{}) {
	fmt.Println(data)
}
