package bravia

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	client *http.Client
	ip     string
	psk    string
	id     int

	System    *System
	AVContent *AVContent
}

func NewClient(ip, psk string) *Client {
	c := &Client{
		client: http.DefaultClient,
		ip:     ip,
		psk:    psk,
	}
	c.System = &System{Client: c}
	c.AVContent = &AVContent{Client: c}
	return c
}

func (c *Client) nextID() int {
	c.id++
	return c.id
}

type rpcRequest struct {
	Method  string `json:"method"`
	ID      int    `json:"id"`
	Params  []any  `json:"params"`
	Version string `json:"version"`
}

func (c *Client) rpc(service, method, version string, params, result any) error {

	rpcBody := &rpcRequest{
		Method:  method,
		ID:      c.nextID(),
		Params:  []any{params},
		Version: version,
	}

	body := &bytes.Buffer{}
	err := json.NewEncoder(body).Encode(rpcBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/sony/%s", c.ip, service), body)
	if err != nil {
		return err
	}

	req.Header.Add("X-Auth-PSK", c.psk)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("request failed with %s", resp.Status)
	}

	if result == nil {
		// io.Copy(os.Stdout, resp.Body)
		// os.Stdout.Write([]byte{'\n'})
		return nil
	}
	return json.NewDecoder(req.Body).Decode(result)
}
