package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_json_validation(t *testing.T) {
	rClient := resty.New().SetBaseURL("http://localhost:8080")
	httpmock.ActivateNonDefault(rClient.GetClient())
	client := NewAdviceClient("/", rClient)

	tests := []struct {
		name     string
		response string
		wantErr  bool
	}{
		{
			name:     "missing id",
			response: `{}`,
			wantErr:  true,
		},
		{
			name:     "valid",
			response: `{"id":"test","icon":"icon","label":"label","tags":["tags1"],"version":"v1.0.0","assessmentBaseQuery":"target.type='com.steadybit.extension_jvm.application'","assessmentQueryAddon":"target.type='com.steadybit.extension_jvm.application'","description":{"actionNeeded":{"summary":"When availability zone ${target.aws.zones?.[0]} is failing, your service ${target.k8s.pod.name} is not available.","motivation":"It is recommended to always split its components into different zones so that in case of a failure of one.","instruction":"Specify the upper limit to be used by defining the   limits   property in your kubernetes manifest: "},"validationNeeded":{"summary":"You already took action and configured it. Validate your configuration via the experiment."},"implemented":{"summary":"When availability zone ${target.aws.zones?.[0]} is failing, your service ${target.k8s.pod.name} is still available."}}}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "http://localhost:8080/test", httpmock.NewStringResponder(200, tt.response))
			_, err := client.DescribeAdvice(advice_kit_api.DescribingEndpointReference{Path: "/test"})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
