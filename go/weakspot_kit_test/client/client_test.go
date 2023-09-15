package client

import (
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/steadybit/weakspot-kit/go/weakspot_kit_api"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_json_validation(t *testing.T) {
	rClient := resty.New().SetBaseURL("http://localhost:8080")
	httpmock.ActivateNonDefault(rClient.GetClient())
	client := NewWeakspotClient("/", rClient)

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
			response: `{ "id": "test", "icon":"icon", "description": "description", "label": "label", "tags": ["tags1"], "version": "v1.0.0", "assesmentBaseQuery": "target.type='com.steadybit.extension_jvm.application'" , "assesmentQueryAddon": "target.type='com.steadybit.extension_jvm.application'" }`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", "http://localhost:8080/test", httpmock.NewStringResponder(200, tt.response))
			_, err := client.DescribeWeakspot(weakspot_kit_api.DescribingEndpointReference{Path: "/test"})
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
