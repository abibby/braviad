package bravia

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrIDMismatch = errors.New("id mismatch")

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

type rpcResult struct {
	ID     int               `json:"id"`
	Error  []any             `json:"error"`
	Result []json.RawMessage `json:"result"`
}

func (c *Client) rpc(service, method, version string, params, result any) error {
	id := c.nextID()
	rpcPrams := []any{}
	if params != nil {
		rpcPrams = []any{params}
	}
	rpcBody := &rpcRequest{
		Method:  method,
		ID:      id,
		Params:  rpcPrams,
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

	r := &rpcResult{}

	// if id > 1 {

	// 	io.Copy(os.Stdout, resp.Body)
	// }
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	if r.ID != id {
		return ErrIDMismatch
	}

	if r.Error != nil {
		return fmt.Errorf("error %v", r.Error)
	}

	if result == nil {
		return nil
	}

	if len(r.Result) == 0 {
		return fmt.Errorf("no results for %s.%s", service, method)
	}

	return json.Unmarshal(r.Result[0], result)
}
