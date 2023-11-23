package tv

import "errors"

var ErrInUse = errors.New("the tv is in use")

type TV interface {
	Activate() error
	PowerOff() error
	DisplayOff() error
}
