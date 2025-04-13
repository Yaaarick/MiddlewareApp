package service

import (
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int {
	d := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)

	dur := time.Until(d)

	return int(dur.Hours() / 24)
}
