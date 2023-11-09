package tv

type TV interface {
	PowerOn() error
	PowerOff() error
	DisplayOn() error
	DisplayOff() error
}
