package tv

import "log"

type Log struct{}

func (l *Log) Activate() error {
	log.Println("Activate")
	return nil
}
func (l *Log) PowerOff() error {
	log.Println("PowerOff")
	return nil
}
func (l *Log) DisplayOff() error {
	log.Println("DisplayOff")
	return nil
}
