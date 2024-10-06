package subdomain

import (
	"errors"
	"strconv"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Subdomain {
	return allEntities
}

func (s *Service) Get(idx int) (*Subdomain, error) {
	return &allEntities[idx], nil
}

func (s *Service) Add(Title string) {
	allEntities = append(allEntities, Subdomain{Title: Title})
}

func (s *Service) Edit(idx int, Title string) (err error) {

	if len(allEntities) < idx {
		return errors.New("Not find produkt with idx:" + strconv.Itoa(idx))
	}

	product := &allEntities[idx]
	product.Title = Title

	return nil
}
