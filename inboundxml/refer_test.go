// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestReferElement(t *testing.T) {
	Convey("Refer should pass", t, func() {

		ixml, err := New(Response{Refer: &Refer{
			Action:         "https://example.com/actionURL",
			Method:         "POST",
			Timeout:        180,
			CallbackUrl:    "https://example.com/callbackURL",
			CallbackMethod: "POST",
			Sip: &Sip{
				Password: "pass",
				Value:    "username@example.com",
				Username: "username",
			},
		}})

		So(ixml, ShouldNotBeNil)
		So(ixml.String(), ShouldEqual, "<Response><Refer action=\"https://example.com/actionURL\" method=\"POST\" timeout=\"180\" callbackUrl=\"https://example.com/callbackURL\" callbackMethod=\"POST\"><Sip username=\"username\" password=\"pass\">username@example.com</Sip></Refer></Response>")
		So(err, ShouldBeNil)
	})
}
