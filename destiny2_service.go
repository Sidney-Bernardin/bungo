package main

type Destiny2Service struct {
	s *Service
}

func NewDestiny2Service(s *Service) *Destiny2Service {
	return &Destiny2Service{s: s}
}
