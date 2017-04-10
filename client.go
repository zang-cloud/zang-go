// Copyright 2017 The Zang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

// Client -
type Client struct {
	*Config
}

// NewClient -
func NewClient() (*Client, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, err
	}

	return &Client{Config: config}, nil
}
