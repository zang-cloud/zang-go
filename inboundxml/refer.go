// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

//Use the <Refer> element to transfer the caller back to the customer's in-house or legacy SIP infrastructure.
//Through initiating the <Refer> verb, CPaaS instructs the caller SIP device to initiate a new call to the external system and replace the CPaaS leg with that new call.
//Therefore, the support of "blind" call transfers in Avaya CPaaS. The <Refer> verb can be invoked on inbound or outbound SIP calls.

//Refer verb attributes.
//action: URL to fetch the set of instructions for further processing. It is executed when the transfer fails with a failure response or when the <Refer> verb is timed out.
//method: Method used to request the action URL.
//timeout: The number of seconds CPaaS should wait for <Refer> verb to conclude.
//callbackUrl: URL where the status of the <Refer> can be sent. Note: This URL only receives parameters containing information about the call. The call does not execute XML given as a callbackUrl.
//callbackMethod: Method used to request the callback URL. Default Value: POST. Allowed Value: POST or GET.

type Refer struct {
	Value          string `xml:",chardata"`
	Action         string `xml:"action,attr,omitempty"`
	Method         string `xml:"method,attr,omitempty"`
	Timeout        int    `xml:"timeout,attr,omitempty"`
	CallbackUrl    string `xml:"callbackUrl,attr,omitempty"`
	CallbackMethod string `xml:"callbackMethod,attr,omitempty"`
	Sip            *Sip   `xml:"Sip,omitempty"`
}
