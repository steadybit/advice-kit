package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-resty/resty/v2"
	"github.com/steadybit/weakspot-kit/go/weakspot_kit_api"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type WeakspotAPI interface {
	ListWeakspots() (weakspot_kit_api.WeakspotList, error)
	DescribeWeakspot(ref weakspot_kit_api.DescribingEndpointReference) (weakspot_kit_api.WeakspotDescription, error)
}

type clientImpl struct {
	client   *resty.Client
	rootPath string
	spec     *openapi3.T
}

func NewWeakspotClient(rootPath string, client *resty.Client) WeakspotAPI {
	spec, _ := weakspot_kit_api.GetSwagger()
	return &clientImpl{
		client:   client,
		rootPath: rootPath,
		spec:     spec,
	}
}

func (c *clientImpl) ListWeakspots() (weakspot_kit_api.WeakspotList, error) {
	var list weakspot_kit_api.WeakspotList
	err := c.executeAndValidate(weakspot_kit_api.DescribingEndpointReference{Path: c.rootPath}, &list, "WeakspotList")
	return list, err
}

func (c *clientImpl) DescribeWeakspot(ref weakspot_kit_api.DescribingEndpointReference) (weakspot_kit_api.WeakspotDescription, error) {
	var description weakspot_kit_api.WeakspotDescription
	err := c.executeAndValidate(ref, &description, "WeakspotDescription")
	return description, err
}

func (c *clientImpl) executeAndValidate(ref weakspot_kit_api.DescribingEndpointReference, result interface{}, schemaName string) error {
	method, path := getMethodAndPath(ref)
	res, err := c.client.R().SetResult(result).Execute(method, path)
	if err != nil {
		return fmt.Errorf("%s %s failed: %w", method, path, err)
	}
	if !res.IsSuccess() {
		return fmt.Errorf("%s %s failed: %d %s", method, path, res.StatusCode(), res.Body())
	}

	return c.validateResponseBody(schemaName, res.Body())
}

func (c *clientImpl) validateResponseBody(name string, body []byte) error {
	if c.spec == nil || name == "" {
		return nil
	}

	schema, ok := c.spec.Components.Schemas[name]
	if !ok {
		return fmt.Errorf("component schema '%s' not found", name)
	}

	var decoded interface{}
	dec := json.NewDecoder(bytes.NewReader(body))
	dec.UseNumber()
	err := dec.Decode(&decoded)
	if err != nil {
		return fmt.Errorf("error decoding response body: %w", err)
	}

	err = schema.Value.VisitJSON(decoded, openapi3.VisitAsResponse())
	if err != nil {
		return fmt.Errorf("error validating response body using schema '%s': %w", name, err)
	}
	return nil
}

func getMethodAndPath(ref weakspot_kit_api.DescribingEndpointReference) (string, string) {
	method := "GET"
	if len(ref.Method) > 0 {
		method = cases.Upper(language.English).String(string(ref.Method))
	}
	return method, ref.Path
}
