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
		Advice: []DescribingEndpointReference{
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
		Id:                        "com.steadybit.extension_example.my_advice",
		Version:                   "1.0.0",
		Label:                     "My Advice",
		Icon:                      "data:image/svg+xml,%3Csvg%20width%3D%22",
		Tags:                      &[]string{"tag1", "tag2"},
		AssessmentQueryApplicable: "target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'",
		Status: &AdviceDefinitionStatus{
			ActionNeeded: AdviceDefinitionStatusActionNeeded{
				AssessmentQuery: "target.attributes['k8s.deployment.has.advice'] = 'true'",
				Description: AdviceDefinitionStatusActionNeededDescription{
					Instruction: "",
					Motivation:  "",
					Summary:     "",
				},
			},
			Implemented: AdviceDefinitionStatusImplemented{
				Description: AdviceDefinitionStatusImplementedDescription{
					Summary: "",
				},
			},
			ValidationNeeded: AdviceDefinitionStatusValidationNeeded{
				Description: AdviceDefinitionStatusValidationNeededDescription{
					Summary: "",
				},
				Validation: Ptr([]Validation{
					{
						Description: "some desc",
						Experiment:  Ptr(Experiment("...experiment here as json export from the UI...")),
						Id:          "com.steadybit.extension_example.my_advice.experiment1",
						Name:        "Experiment 1",
						Type:        "experiment",
					}, {
						Description: "please check this",
						Experiment:  nil,
						Id:          "com.steadybit.extension_example.my_advice.check1",
						Name:        "Validation 1",
						Type:        "text",
					},
				}),
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
