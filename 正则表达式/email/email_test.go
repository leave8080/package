package email

import (
	"gofer/pkg/log"
	"testing"
)

func TestVerifyEmailFormat(t *testing.T) {
	log.Info(VerifyEmailFormat("wunaiwei8080@gmail.com"))
}

type Rout struct {
	pa *Rout
	en *En
}

type En struct {
	*Rout
	r []*Rout
}
