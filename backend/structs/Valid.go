package structs

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/base/camStructs"
)

type Valid struct {
	cam.ValidInterface
	Email   string
	MyEmail Email
}

type Email string

func (s *Valid) Rules() []camStatics.RuleInterface {
	return []camStatics.RuleInterface{
		camStructs.NewRule([]string{"Email", "MyEmail"}, cam.Rule.Email, cam.Rule.Length(0, 100)),
	}
}
