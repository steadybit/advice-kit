# AdviceKit Go API

This module exposes Go types that you will find helpful when implementing an AdviceKit extension.

The types are generated automatically from the AdviceKit [OpenAPI specification](https://github.com/steadybit/advice-kit/tree/main/openapi).

## Installation

Add the following to your `go.mod` file:

```
go get github.com/steadybit/advice-kit/go/advice_kit_api@v0.1.0
```

## Usage

```go
import (
	"github.com/steadybit/advice-kit/go/advice_kit_api"
)

AdviceList := advice_kit_api.AdviceList{
    Advice: []advice_kit_api.DescribingEndpointReference{
        {
            "GET",
            "/advice/my-fancy-advice",
        },
    },
}
```