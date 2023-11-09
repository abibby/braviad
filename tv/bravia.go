package tv

import "github.com/abibby/braviad/bravia"

type Bravia struct {
	client *bravia.Client
}

func (b *Bravia) inUse() (bool, error) {
	return false, nil
}

func (b *Bravia) PowerOn() error {
	inUse, err := b.inUse()
	if err != nil {
		return err
	}
	if inUse {
		return nil
	}
	err = b.client.System.SetPowerStatus(true)
	if err != nil {
		return err
	}
	err = b.client.AVContent.SetPlayContent("extInput:hdmi?port=4")
	if err != nil {
		return err
	}
	return nil
}
func (b *Bravia) PowerOff() error {

	return nil
}
func (b *Bravia) DisplayOn() error {

	return nil
}
func (b *Bravia) DisplayOff() error {

	return nil
}
