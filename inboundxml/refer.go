// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

type Refer struct {
	Value          string `xml:",chardata"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	Timeout        int    `xml:"timeout,attr,omitempty"`
	CallbackUrl    string `xml:"callbackUrl,attr,omitempty"`
	CallbackMethod string `xml:"callbackMethod,attr,omitempty"`
	Sip            *Sip   `xml:"Sip,omitempty"`
}
