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

func (s *Valid) Rules() []camBase.RuleInterface {
	return []camBase.RuleInterface{
		cam.NewRule([]string{"Email", "MyEmail"}, cam.Rule.Email, cam.Rule.Length(0, 100)),
	}
}
