// SPDX-License-Identifier: MIT
//go:build samples

package samples

import (
	. "github.com/spandigital/must"
	"net/url"
	"time"
)

var timeWhichMustParse = Must(time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"))
var urlsWhichMustParse = Must(url.Parse("http://www.google.com"))
