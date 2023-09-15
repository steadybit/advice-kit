// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2023 Steadybit GmbH

package weakspot_kit_api

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

func TestWeakspotList(t *testing.T) {
	v := WeakspotList{
		Weakspots: []DescribingEndpointReference{
			{
				Method: "GET",
				Path:   "/",
			},
		},
	}
	markAsUsed(t, v)
}

func TestWeakspotDescription(t *testing.T) {
	v := WeakspotDescription{
		Id:                  "com.steadybit.extension_example.my_weakspot",
		Version:             "1.0.0",
		Label:               "My Weakspot",
		Icon:                "data:image/svg+xml,%3Csvg%20width%3D%22",
		Tags:                &[]string{"tag1", "tag2"},
		AssesmentBaseQuery:  "target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'",
		AssesmentQueryAddon: "target.attributes['k8s.deployment.has.weakspot'] = 'true'",
		ExperimentTemplates: &[]ExperimentTemplate{
			"...experiment here",
		},
	}
	markAsUsed(t, v)
}

func TestWeakspotKitError(t *testing.T) {
	v := WeakspotKitError{
		Detail:   Ptr("d"),
		Instance: Ptr("i"),
		Title:    "t",
		Type:     Ptr("t"),
	}
	markAsUsed(t, v)
}
