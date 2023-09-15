package validate

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/steadybit/weakspot-kit/go/weakspot_kit_test/client"
)

func ValidateEndpointReferences(path string, restyClient *resty.Client) error {
	c := client.NewWeakspotClient(path, restyClient)
	var allErr error

	list, err := c.ListWeakspots()
	allErr = errors.Join(allErr, err)

	for _, weakspot := range list.Weakspots {
		_, err := c.DescribeWeakspot(weakspot)
		allErr = errors.Join(allErr, err)
	}

	return allErr
}
