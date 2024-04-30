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
			response: `{"id":"com.steadybit.extension_aws.advice.aws-single-zone","version":"1.0.0","label":"AWS Single Zone","icon":"data:image/svg+xml;utf8,<svg ...>","assessmentQueryApplicable":"target.type = com.steadybit.extension_kubernetes.kubernetes-deployment","tags":["AWS"],"status":{"actionNeeded":{"assessmentQuery":"aws.zone > 0 AND aws.zone < 2","description":{"summary":"When availability zone ${target.aws.zone} is failing, your service ${target.attr('k8s.pod.name')} is not available.","motivation":"It is recommended to always split its components into different zones so that in case of a failure of one.","instruction":"Specify the upper limit to be used by defining the   limits   property in your kubernetes manifest: "}},"validationNeeded":{"description":{"summary":"You already took action and configured it. Validate your configuration via the experiment."},"validation":[{"id":"com.steadybit.extension_aws.advice.aws-single-zone.1","name":"AWS Single Zone Template","type":"EXPERIMENT","description":"...","shortDescription":"...","experiment":{"name":"${target.k8s.cluster-name}/${target.k8s.deployment} faultless redundancy during single pod failure","hypothesis":"When a single pod fails of all requests to an endpoint are successful","lanes":{"steps":[{"type":"action","ignoreFailure":false,"parameters":{"duration":"60s","cpuLoad":100,"workers":0},"actionType":"com.steadybit.extension_container.stress_cpu","radius":{"targetType":"com.steadybit.extension_container.container","predicate":{"operator":"AND","predicates":[{"key":"k8s.cluster-name","operator":"EQUALS","values":["${target.k8s.cluster-name}"]},{"key":"k8s.namespace","operator":"EQUALS","values":["${target.k8s.namespace}"]},{"key":"k8s.deployment","operator":"EQUALS","values":["${target.k8s.deployment}"]},{"key":"k8s.container.name","operator":"EQUALS","values":["${target.k8s.container.name}"]}]},"query":null,"percentage":50}}]}}},{"id":"com.steadybit.extension_aws.advice.aws-single-zone.2","name":"Check me","description":"Really check me","shortDescription":"...","type":"TEXT"}]},"implemented":{"description":{"summary":"When availability zone ${target.aws.zones?.[0]} is failing, your service ${target.attr('k8s.pod.name')} is still available."}}}}`,
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
