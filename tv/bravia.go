package tv

import (
	"log"

	"github.com/abibby/braviad/bravia"
)

type Bravia struct {
	client *bravia.Client
}

type BraviaConfig struct {
	IP   string `json:"ip"`
	PSK  string `json:"psk"`
	Port int    `json:"port"`
}

func NewBravia(cfg *BraviaConfig) *Bravia {
	return &Bravia{
		client: bravia.NewClient(cfg.IP, cfg.PSK),
	}
}

func (b *Bravia) inUse() error {
	powerStatus, err := b.client.System.GetPowerStatus()
	if err != nil {
		return err
	}

	content, err := b.client.AVContent.GetPlayingContentInfo()
	if err != nil {
		return err
	}
	log.Print(content)
	if powerStatus == "active" && content.URI != "extInput:hdmi?port=4" {
		return ErrInUse
	}
	return nil
}

func (b *Bravia) Activate() error {
	err := b.inUse()
	if err != nil {
		return err
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
	err := b.inUse()
	if err != nil {
		return err
	}
	err = b.client.System.SetPowerStatus(false)
	if err != nil {
		return err
	}
	return nil
}
func (b *Bravia) DisplayOff() error {
	err := b.inUse()
	if err != nil {
		return err
	}
	err = b.client.System.SetPowerStatus(false)
	if err != nil {
		return err
	}
	return nil
}
