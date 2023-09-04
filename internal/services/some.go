package services

import (
	"github.com/sabahtalateh/gic"
	"github.com/sabahtalateh/gicex/internal/repos"
)

func init() {
	gic.Add[*SomeService](
		gic.WithInit(func() *SomeService {
			return &SomeService{r: gic.MustGet[*repos.SomeRepo]()}
		}),
	)
}

type SomeRepo interface {
	GetSome() (string, error)
}

type SomeService struct {
	r SomeRepo
}

func (s *SomeService) GetSome() (string, error) {
	return s.r.GetSome()
}
