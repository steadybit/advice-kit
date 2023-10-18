// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package advice_kit_api

import (
	"testing"
)

// Note: These test cases only check that the code compiles as intended.
// On compilation errors, we most likely have caused a breaking change of
// the API and need to act accordingly.

func markAsUsed(t *testing.T, v any) {
	if v == nil {
		t.Fail()
	}
}

func TestAdviceList(t *testing.T) {
	v := AdviceList{
		Advices: []DescribingEndpointReference{
			{
				Method: "GET",
				Path:   "/",
			},
		},
	}
	markAsUsed(t, v)
}

func TestAdviceDefinition(t *testing.T) {
	v := AdviceDefinition{
		Id:                  "com.steadybit.extension_example.my_advice",
		Version:             "1.0.0",
		Label:               "My Advice",
		Icon:                "data:image/svg+xml,%3Csvg%20width%3D%22",
		Tags:                &[]string{"tag1", "tag2"},
		AssessmentBaseQuery:  "target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'",
		AssessmentQueryAddon: "target.attributes['k8s.deployment.has.advice'] = 'true'",
		Experiments: &[]Experiment{
			"...experiment here as json export from the UI...",
		},
		Description: AdviceDefinitionDescription{
			ActionNeeded:     AdviceDefinitionDescriptionActionNeeded{
				Instruction: "",
				Motivation:  "",
				Summary:     "",
			},
			Implemented:      AdviceDefinitionDescriptionImplemented{
				Summary: "",
			},
			ValidationNeeded: AdviceDefinitionDescriptionValidationNeeded{
				Summary: "",
			},
		},
	}
	markAsUsed(t, v)
}

func TestAdviceKitError(t *testing.T) {
	v := AdviceKitError{
		Detail:   Ptr("d"),
		Instance: Ptr("i"),
		Title:    "t",
		Type:     Ptr("t"),
	}
	markAsUsed(t, v)
}
