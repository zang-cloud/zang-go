// Copyright 2017 The Zang Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPauseElement(t *testing.T) {
	Convey("Pause should pass", t, func() {
		ixml, err := New(Response{Pause: &Pause{
			Length: 5,
		}})
		So(ixml, ShouldNotBeNil)
		So(err, ShouldBeNil)

	})
}
