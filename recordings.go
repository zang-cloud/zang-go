// Copyright 2017 The Zang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"

	"github.com/zang-cloud/zang-go/helpers"
)

// ListRecordings -
func (c *Client) ListRecordings(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Recordings", c.Config.AccountSid),
		params,
	)
	return
}

// GetRecording -
func (c *Client) GetRecording(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/Recordings/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteRecording -
func (c *Client) DeleteRecording(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/Recordings/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}
