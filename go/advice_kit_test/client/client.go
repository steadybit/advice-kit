package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-resty/resty/v2"
	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type AdviceAPI interface {
	ListAdvice() (advice_kit_api.AdviceList, error)
	DescribeAdvice(ref advice_kit_api.DescribingEndpointReference) (advice_kit_api.AdviceDefinition, error)
}

type clientImpl struct {
	client   *resty.Client
	rootPath string
	spec     *openapi3.T
}

func NewAdviceClient(rootPath string, client *resty.Client) AdviceAPI {
	spec, _ := advice_kit_api.GetSwagger()
	return &clientImpl{
		client:   client,
		rootPath: rootPath,
		spec:     spec,
	}
}

func (c *clientImpl) ListAdvice() (advice_kit_api.AdviceList, error) {
	var list advice_kit_api.AdviceList
	err := c.executeAndValidate(advice_kit_api.DescribingEndpointReference{Path: c.rootPath}, &list, "AdviceList")
	return list, err
}

func (c *clientImpl) DescribeAdvice(ref advice_kit_api.DescribingEndpointReference) (advice_kit_api.AdviceDefinition, error) {
	var description advice_kit_api.AdviceDefinition
	err := c.executeAndValidate(ref, &description, "AdviceDefinition")
	return description, err
}

func (c *clientImpl) executeAndValidate(ref advice_kit_api.DescribingEndpointReference, result interface{}, schemaName string) error {
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

func getMethodAndPath(ref advice_kit_api.DescribingEndpointReference) (string, string) {
	method := "GET"
	if len(ref.Method) > 0 {
		method = cases.Upper(language.English).String(string(ref.Method))
	}
	return method, ref.Path
}
