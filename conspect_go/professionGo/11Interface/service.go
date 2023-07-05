//service.go

package main

type Service struct {
	description    string
	durationMonths int
	monthlyFee     float64
	features       []string
}

func (s Service) getName() string {
	return s.description
}

func (s Service) getCoast(_ bool) float64 {
	return s.monthlyFee * float64(s.durationMonths)
}
