// Package advice_kit_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.13.4 DO NOT EDIT.
package advice_kit_api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
)

// Defines values for DescribingEndpointReferenceMethod.
const (
	GET DescribingEndpointReferenceMethod = "GET"
)

// AdviceDefinition Provides details about a advice
type AdviceDefinition struct {
	// AssessmentQueryActionNeeded A Assessment Target Query Addon that is used to identify targets with this advice in the target list of the assessmentQueryApplicable
	AssessmentQueryActionNeeded string `json:"assessmentQueryActionNeeded"`

	// AssessmentQueryApplicable A Assessment Target Base Query that is used identifies targets that could have this advice.
	AssessmentQueryApplicable string `json:"assessmentQueryApplicable"`

	// Description Provides details about a advice
	Description AdviceDefinitionDescription `json:"description"`

	// Experiments A list of experiment templates that are available for this advice.
	Experiments *[]ExperimentTemplate `json:"experiments,omitempty"`

	// Icon A svg of an icon that represents the advice.
	Icon string `json:"icon"`

	// Id A technical ID that is used to uniquely identify this type of advice. You will typically want to use something like `org.example.extension.my-fancy-advice`.
	Id string `json:"id"`

	// Label A human-readable label for the advice.
	Label string `json:"label"`

	// Tags A list of tags that describe the advice.
	Tags *[]string `json:"tags,omitempty"`

	// Version The version of the advice. This is used to identify the version of the advice and is used for compatibility checks.
	Version string `json:"version"`
}

// AdviceDefinitionDescription Provides details about a advice
type AdviceDefinitionDescription struct {
	// ActionNeeded Provides details about a advice actions needed
	ActionNeeded AdviceDefinitionDescriptionActionNeeded `json:"actionNeeded"`

	// Implemented Provides details about a advice implemented
	Implemented AdviceDefinitionDescriptionImplemented `json:"implemented"`

	// ValidationNeeded Provides details about a validation actions needed
	ValidationNeeded AdviceDefinitionDescriptionValidationNeeded `json:"validationNeeded"`
}

// AdviceDefinitionDescriptionActionNeeded Provides details about a advice actions needed
type AdviceDefinitionDescriptionActionNeeded struct {
	// Instruction A human-readable instructions of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Instruction string `json:"instruction"`

	// Motivation A human-readable motivation of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Motivation string `json:"motivation"`

	// Summary A human-readable summary of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceDefinitionDescriptionImplemented Provides details about a advice implemented
type AdviceDefinitionDescriptionImplemented struct {
	// Summary A human-readable summary of the implemented in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceDefinitionDescriptionValidationNeeded Provides details about a validation actions needed
type AdviceDefinitionDescriptionValidationNeeded struct {
	// Summary A human-readable summary of the validation needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceKitError RFC 7807 Problem Details for HTTP APIs compliant response body for error scenarios
type AdviceKitError struct {
	// Detail A human-readable explanation specific to this occurrence of the problem.
	Detail *string `json:"detail,omitempty"`

	// Instance A URI reference that identifies the specific occurrence of the problem.
	Instance *string `json:"instance,omitempty"`

	// Title A short, human-readable summary of the problem type.
	Title string `json:"title"`

	// Type A URI reference that identifies the problem type.
	Type *string `json:"type,omitempty"`
}

// AdviceList Lists all advice that the platform/agent could consume.
type AdviceList struct {
	Advice []DescribingEndpointReference `json:"advice"`
}

// DescribingEndpointReference HTTP endpoint which the Steadybit platform/agent could communicate with.
type DescribingEndpointReference struct {
	// Method HTTP method to use when calling the HTTP endpoint.
	Method DescribingEndpointReferenceMethod `json:"method"`

	// Path Absolute path of the HTTP endpoint.
	Path string `json:"path"`
}

// DescribingEndpointReferenceMethod HTTP method to use when calling the HTTP endpoint.
type DescribingEndpointReferenceMethod string

// Experiment Provides a experiment json exported from the ui
type Experiment = interface{}

// ExperimentTemplate Provides a template about a advice experiment
type ExperimentTemplate struct {
	// Description A human-readable description for the experiment template.
	Description *string `json:"description,omitempty"`

	// Experiment Provides a experiment json exported from the ui
	Experiment Experiment `json:"experiment"`

	// Id A technical ID that is used to uniquely identify this experiment. You will typically want to use something like `org.example.extension.my-fancy-advice.1`.
	Id string `json:"id"`

	// Name A human-readable name for the experiment template.
	Name string `json:"name"`
}

// AdviceDefinitionResponse defines model for AdviceDefinitionResponse.
type AdviceDefinitionResponse struct {
	union json.RawMessage
}

// AdviceListResponse defines model for AdviceListResponse.
type AdviceListResponse struct {
	union json.RawMessage
}

// AsAdviceDefinition returns the union data inside the AdviceDefinitionResponse as a AdviceDefinition
func (t AdviceDefinitionResponse) AsAdviceDefinition() (AdviceDefinition, error) {
	var body AdviceDefinition
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceDefinition overwrites any union data inside the AdviceDefinitionResponse as the provided AdviceDefinition
func (t *AdviceDefinitionResponse) FromAdviceDefinition(v AdviceDefinition) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceDefinition performs a merge with any union data inside the AdviceDefinitionResponse, using the provided AdviceDefinition
func (t *AdviceDefinitionResponse) MergeAdviceDefinition(v AdviceDefinition) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAdviceKitError returns the union data inside the AdviceDefinitionResponse as a AdviceKitError
func (t AdviceDefinitionResponse) AsAdviceKitError() (AdviceKitError, error) {
	var body AdviceKitError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceKitError overwrites any union data inside the AdviceDefinitionResponse as the provided AdviceKitError
func (t *AdviceDefinitionResponse) FromAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceKitError performs a merge with any union data inside the AdviceDefinitionResponse, using the provided AdviceKitError
func (t *AdviceDefinitionResponse) MergeAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AdviceDefinitionResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AdviceDefinitionResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsAdviceList returns the union data inside the AdviceListResponse as a AdviceList
func (t AdviceListResponse) AsAdviceList() (AdviceList, error) {
	var body AdviceList
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceList overwrites any union data inside the AdviceListResponse as the provided AdviceList
func (t *AdviceListResponse) FromAdviceList(v AdviceList) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceList performs a merge with any union data inside the AdviceListResponse, using the provided AdviceList
func (t *AdviceListResponse) MergeAdviceList(v AdviceList) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAdviceKitError returns the union data inside the AdviceListResponse as a AdviceKitError
func (t AdviceListResponse) AsAdviceKitError() (AdviceKitError, error) {
	var body AdviceKitError
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAdviceKitError overwrites any union data inside the AdviceListResponse as the provided AdviceKitError
func (t *AdviceListResponse) FromAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAdviceKitError performs a merge with any union data inside the AdviceListResponse, using the provided AdviceKitError
func (t *AdviceListResponse) MergeAdviceKitError(v AdviceKitError) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AdviceListResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AdviceListResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//knjaWTYYAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
