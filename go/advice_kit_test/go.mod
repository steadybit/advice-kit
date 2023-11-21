module github.com/steadybit/advice-kit/go/advice_kit_test

go 1.21

require (
	github.com/getkin/kin-openapi v0.120.0
	github.com/go-resty/resty/v2 v2.10.0
	github.com/jarcoal/httpmock v1.3.1
	github.com/steadybit/advice-kit/go/advice_kit_api v0.0.1-beta.3
	github.com/stretchr/testify v1.8.4
	golang.org/x/text v0.14.0
)

replace (
	github.com/steadybit/advice-kit/go/advice_kit_api v0.0.1-beta.3 => ../advice_kit_api
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/jsonpointer v0.20.0 // indirect
	github.com/go-openapi/swag v0.22.4 // indirect
	github.com/google/uuid v1.4.0 // indirect
	github.com/invopop/yaml v0.2.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oapi-codegen/runtime v1.1.0 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
