package structs

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camBase"
)

type Valid struct {
	cam.ValidInterface
	Email   string
	MyEmail Email
}

type Email string

func (s *Valid) Rules() []camBase.ValidRuleInterface {
	return []camBase.ValidRuleInterface{
		cam.NewRule([]string{"Email", "MyEmail"}, []camBase.ValidHandler{cam.Rule.Email}),
	}
}
