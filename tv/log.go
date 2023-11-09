package tv

import "log"

type Log struct{}

func (l *Log) PowerOn() error {
	log.Println("PowerOn")
	return nil
}
func (l *Log) PowerOff() error {
	log.Println("PowerOff")
	return nil
}
func (l *Log) DisplayOn() error {
	log.Println("DisplayOn")
	return nil
}
func (l *Log) DisplayOff() error {
	log.Println("DisplayOff")
	return nil
}
