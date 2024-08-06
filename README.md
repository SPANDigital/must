# must
`Must` functions which panic on error, there is a `Must2` and `Must3` variant as well.

## See the samples

Have a [look at samples](/samples) for samples us usage.

## Code example

```go
package samples

import (
	. "github.com/spandigital/must"
	"net/url"
	"time"
)

var timeWhichMustParse = Must(time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"))
var urlsWhichMustParse = Must(url.Parse("http://www.google.com"))
```