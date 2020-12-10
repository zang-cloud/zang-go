// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConnectElement(t *testing.T) {
	Convey("Connect should pass", t, func() {
		ixml, err := New(Response{Connect: &Connect{
			Action: "http://example.com/example-callback-url/say?example=simple.xml",
			Method: "GET",
			Agent: &Agent{
				Value: "Please enter your agent ID",
			},
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
