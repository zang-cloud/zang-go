// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAgentElement(t *testing.T) {
	Convey("Agent should pass", t, func() {
		ixml, err := New(Response{Connect: &Connect{
			Agent: &Agent{
				Value: "1234",
			},
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
