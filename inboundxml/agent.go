// Copyright 2017 Avaya Cloud Inc. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package inboundxml

// Agent - the <Agent> element contains the Agent SID
type Agent struct {
	Value string `xml:",chardata"`
}
