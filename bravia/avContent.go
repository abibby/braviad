package bravia

type setPlayContentOptions struct {
	URI string `json:"uri"`
}

type AVContent struct {
	*Client
}

func (c *AVContent) SetPlayContent(uri string) error {
	options := &setPlayContentOptions{URI: uri}
	return c.rpc("avContent", "setPlayContent", "1.0", options, nil)
}
