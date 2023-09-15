# WeakspotKit Go API

This module exposes Go types that you will find helpful when implementing an WeakspotKit extension.

The types are generated automatically from the WeakspotKit [OpenAPI specification](https://github.com/steadybit/weakspot-kit/tree/main/openapi).

## Installation

Add the following to your `go.mod` file:

```
go get github.com/steadybit/weakspot-kit/go/weakspot_kit_api@v0.1.0
```

## Usage

```go
import (
	"github.com/steadybit/weakspot-kit/go/weakspot_kit_api"
)

WeakspotList := weakspot_kit_api.WeakspotList{
    Weakspots: []weakspot_kit_api.DescribingEndpointReference{
        {
            "GET",
            "/weakspots/my-fancy-weakspot",
        },
    },
}
```