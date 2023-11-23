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

type GetPlayingContentInfoResult struct {
	Source string `json:"source"`
	Title  string `json:"title"`
	URI    string `json:"uri"`
}

// https://pro-bravia.sony.net/develop/integrate/rest-api/spec/service/avcontent/v1_0/getPlayingContentInfo/index.html
func (c *AVContent) GetPlayingContentInfo() (*GetPlayingContentInfoResult, error) {
	r := &GetPlayingContentInfoResult{}
	err := c.rpc("avContent", "getPlayingContentInfo", "1.0", nil, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
