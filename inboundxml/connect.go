// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Connect - The <Connect> element allows us to plug and play with any bot
type Connect struct {
	Action string `xml:"action,attr,omitempty"`
	Method string `xml:"method,attr,omitempty"`
	Agent  *Agent `xml:"Agent,omitempty"`
}
