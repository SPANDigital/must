# must
`Must` functions which panic on error, there is a `Must2` and `Must3` variant as well.

[![Open in Dev Containers](https://img.shields.io/static/v1?label=Dev%20Containers&message=Open&color=blue&logo=visualstudiocode)](https://vscode.dev/redirect?url=vscode://ms-vscode-remote.remote-containers/cloneInVolume?url=https://github.com/SPANDigital/must)
![Develop Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/must/go.yml?branch=develop&label=develop)
![Main Go Action Workflow Status](https://img.shields.io/github/actions/workflow/status/spandigital/must/go.yml?branch=main&label=main)
![Tag status](https://img.shields.io/github/v/tag/SPANDigital/must)

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