// Package advice_kit_api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
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

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/oapi-codegen/runtime"
)

// Defines values for AdviceDefinitionTemplateEngine.
const (
	FREEMARKER AdviceDefinitionTemplateEngine = "FREEMARKER"
	LEGACY     AdviceDefinitionTemplateEngine = "LEGACY"
)

// Defines values for DescribingEndpointReferenceMethod.
const (
	GET DescribingEndpointReferenceMethod = "GET"
)

// Defines values for ValidationType.
const (
	EXPERIMENT ValidationType = "EXPERIMENT"
	TEXT       ValidationType = "TEXT"
)

// AdviceDefinition Provides details about a advice
type AdviceDefinition struct {
	// AssessmentQueryApplicable A Assessment Target Base Query that is used identifies targets that could have this advice.
	AssessmentQueryApplicable string `json:"assessmentQueryApplicable"`

	// Icon A svg of an icon that represents the advice.
	Icon string `json:"icon"`

	// Id A technical ID that is used to uniquely identify this type of advice. You will typically want to use something like `org.example.extension.my-fancy-advice`.
	Id string `json:"id"`

	// Label A human-readable label for the advice.
	Label string `json:"label"`

	// Status Provides details about a advice
	Status AdviceDefinitionStatus `json:"status"`

	// Tags A list of tags that describe the advice.
	Tags *[]string `json:"tags,omitempty"`

	// TemplateEngine The template engine that is used to render different aspects including validation experiments of the advice.
	TemplateEngine *AdviceDefinitionTemplateEngine `json:"templateEngine,omitempty"`

	// Version The version of the advice. This is used to identify the version of the advice and is used for compatibility checks.
	Version string `json:"version"`
}

// AdviceDefinitionTemplateEngine The template engine that is used to render different aspects including validation experiments of the advice.
type AdviceDefinitionTemplateEngine string

// AdviceDefinitionStatus Provides details about a advice
type AdviceDefinitionStatus struct {
	// ActionNeeded Provides details about a advice lifecycle status actions needed
	ActionNeeded AdviceDefinitionStatusActionNeeded `json:"actionNeeded"`

	// Implemented Provides details about a advice lifecycle status implemented
	Implemented AdviceDefinitionStatusImplemented `json:"implemented"`

	// ValidationNeeded Provides details about a advice lifecycle status validation needed
	ValidationNeeded AdviceDefinitionStatusValidationNeeded `json:"validationNeeded"`
}

// AdviceDefinitionStatusActionNeeded Provides details about a advice lifecycle status actions needed
type AdviceDefinitionStatusActionNeeded struct {
	// AssessmentQuery A Assessment Target Query Addon that is used to identify targets with this advice in the target list of the assessmentQueryApplicable
	AssessmentQuery string `json:"assessmentQuery"`

	// Description Provides details description about a advice lifecycle status actions needed
	Description AdviceDefinitionStatusActionNeededDescription `json:"description"`
}

// AdviceDefinitionStatusActionNeededDescription Provides details description about a advice lifecycle status actions needed
type AdviceDefinitionStatusActionNeededDescription struct {
	// Instruction A human-readable instructions of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Instruction string `json:"instruction"`

	// Motivation A human-readable motivation of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Motivation string `json:"motivation"`

	// Summary A human-readable summary of the action needed in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceDefinitionStatusImplemented Provides details about a advice lifecycle status implemented
type AdviceDefinitionStatusImplemented struct {
	// Description Provides details description about a advice lifecycle status implemented
	Description AdviceDefinitionStatusImplementedDescription `json:"description"`
}

// AdviceDefinitionStatusImplementedDescription Provides details description about a advice lifecycle status implemented
type AdviceDefinitionStatusImplementedDescription struct {
	// Summary A human-readable summary of the implemented in mark down format. (you can use placeholder like ${target.k8s.pod.name})
	Summary string `json:"summary"`
}

// AdviceDefinitionStatusValidationNeeded Provides details about a advice lifecycle status validation needed
type AdviceDefinitionStatusValidationNeeded struct {
	// Description Provides details description about a advice lifecycle status validation needed
	Description AdviceDefinitionStatusValidationNeededDescription `json:"description"`

	// Validation A list of validations that are available for this advice.
	Validation *[]Validation `json:"validation,omitempty"`
}

// AdviceDefinitionStatusValidationNeededDescription Provides details description about a advice lifecycle status validation needed
type AdviceDefinitionStatusValidationNeededDescription struct {
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

// Validation Provides a either a template about a advice experiment or a textual validation like a checklist item
type Validation struct {
	// Description A human-readable description for the validation or for the experiment template. Markdown supported.
	Description *string `json:"description,omitempty"`

	// Experiment Provides a experiment json exported from the ui
	Experiment *Experiment `json:"experiment,omitempty"`

	// Id A technical ID that is used to uniquely identify this validation. You will typically want to use something like `org.example.extension.my-fancy-advice-validation.1`.
	Id string `json:"id"`

	// Name A human-readable name for the validation.
	Name string `json:"name"`

	// ShortDescription A human-readable short description for the validation or for the experiment template. Text Only.
	ShortDescription string `json:"shortDescription"`

	// Type The type of the validation. Either `EXPERIMENT` or `TEXT`.
	Type ValidationType `json:"type"`
}

// ValidationType The type of the validation. Either `EXPERIMENT` or `TEXT`.
type ValidationType string

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

	merged, err := runtime.JSONMerge(t.union, b)
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

	merged, err := runtime.JSONMerge(t.union, b)
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

	merged, err := runtime.JSONMerge(t.union, b)
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

	merged, err := runtime.JSONMerge(t.union, b)
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
