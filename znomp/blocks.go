package znomp

import (
	"net/http"
)

// BlocksResponse representation of /api/blocks response.
type BlocksResponse []Block

// Blocks returns the global blocks stats.
func (c Client) Blocks() (BlocksResponse, error) {
	var (
		err error
		req *http.Request
	)

	req, err = http.NewRequest("GET", c.endpoint, nil)
	if err != nil {
		return nil, err
	}

	err = c.authenticate(req)
	if err != nil {
		return nil, err
	}

	// TODO deserialize json
	res := make(BlocksResponse, 0)

	return res, nil
}
