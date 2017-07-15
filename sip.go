// Copyright 2017 The Zang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package zang

import (
	"fmt"

	"github.com/zang-cloud/zang-go/helpers"
)

// ListSipDomains -
func (c *Client) ListSipDomains(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains", c.Config.AccountSid),
		params,
	)
	return

}

// GetSipDomain -
func (c *Client) GetSipDomain(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// CreateSipDomain -
func (c *Client) CreateSipDomain(params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/Domains", c.Config.AccountSid),
		params,
	)
	return
}

// UpdateSipDomain -
func (c *Client) UpdateSipDomain(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Update(
		fmt.Sprintf("%s/SIP/Domains/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteSipDomain -
func (c *Client) DeleteSipDomain(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/Domains/%s", c.Config.AccountSid, sid),
		params,
	)
	return
}

// ListSipDomainCredentialMappings -
func (c *Client) ListSipDomainCredentialMappings(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains/%s/CredentialListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// MapSipDomainCredentialList -
func (c *Client) MapSipDomainCredentialList(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/Domains/%s/CredentialListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteSipDomainCredentialListMapping -
func (c *Client) DeleteSipDomainCredentialListMapping(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/Domains/%s/CredentialListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// ListSipDomainIpAclMappings -
func (c *Client) ListSipDomainIpAclMappings(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Get(
		fmt.Sprintf("%s/SIP/Domains/%s/IpAccessControlListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// MapSipDomainIpAcl -
func (c *Client) MapSipDomainIpAcl(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Create(
		fmt.Sprintf("%s/SIP/Domains/%s/IpAccessControlListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}

// DeleteSipDomainIpAclMapping -
func (c *Client) DeleteSipDomainIpAclMapping(sid string, params map[string]string) (resp *Response, err error) {
	if cerr := c.Config.Validate(); cerr != nil {
		return nil, cerr
	}

	if verr := helpers.ValidateSid(sid); err != nil {
		return nil, verr
	}

	resp, err = c.Request.Delete(
		fmt.Sprintf("%s/SIP/Domains/%s/IpAccessControlListMappings", c.Config.AccountSid, sid),
		params,
	)
	return
}
