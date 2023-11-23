package bravia

type System struct {
	*Client
}

type setPowerStatusOptions struct {
	Status bool `json:"status"`
}

// https://pro-bravia.sony.net/develop/integrate/rest-api/spec/service/system/v1_0/setPowerStatus/index.html
func (c *System) SetPowerStatus(status bool) error {
	options := &setPowerStatusOptions{Status: status}
	return c.rpc("system", "setPowerStatus", "1.0", options, nil)
}

type getPowerStatusResult struct {
	Status string `json:"status"`
}

// https://pro-bravia.sony.net/develop/integrate/rest-api/spec/service/system/v1_0/getPowerStatus/index.html
func (c *System) GetPowerStatus() (string, error) {
	result := &getPowerStatusResult{}
	err := c.rpc("system", "getPowerStatus", "1.0", nil, result)
	if err != nil {
		return "", err
	}
	return result.Status, nil
}
