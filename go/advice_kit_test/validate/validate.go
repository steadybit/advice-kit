package validate

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/steadybit/advice-kit/go/advice_kit_test/client"
)

func ValidateEndpointReferences(path string, restyClient *resty.Client) error {
	c := client.NewAdviceClient(path, restyClient)
	var allErr error

	list, err := c.ListAdvice()
	allErr = errors.Join(allErr, err)

	for _, advice := range list.Advice {
		_, err := c.DescribeAdvice(advice)
		allErr = errors.Join(allErr, err)
	}

	return allErr
}
