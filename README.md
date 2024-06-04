# must
Must functions which panic on errors

## See the samples

Have a [look at samples](/samples) for samples us usage.

## Code example

```go
package samples

import (
	"github.com/spandigital/must"
	"net/url"
	"time"
)

var timeWhichMustParse = must.Must(time.Parse(time.RFC3339, "2006-01-02T15:04:05Z"))
var urlsWhichMustParse = must.Must(url.Parse("http://www.google.com"))
```