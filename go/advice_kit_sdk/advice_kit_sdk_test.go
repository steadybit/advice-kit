// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Steadybit GmbH

package advice_kit_sdk

import (
	"reflect"
	"testing"

	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"github.com/steadybit/extension-kit/extutil"
)

func Test_applyExcludeQuery(t *testing.T) {
	tests := []struct {
		name              string
		presentQuery      *string
		configuredExclude string
		wantedQuery       *string
	}{
		{
			name: "no configured exclude, no present query",
		},
		{
			name:         "no configured exclude, present query",
			presentQuery: extutil.Ptr("target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'"),
			wantedQuery:  extutil.Ptr("target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'"),
		},
		{
			name:              "configured exclude, no present query",
			configuredExclude: "k8s.namespace = 'kube-system'",
			wantedQuery:       extutil.Ptr("k8s.namespace = 'kube-system'"),
		},
		{
			name:              "configured exclude, present query",
			presentQuery:      extutil.Ptr("target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'"),
			configuredExclude: "k8s.namespace = 'kube-system'",
			wantedQuery:       extutil.Ptr("(target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment') OR (k8s.namespace = 'kube-system')"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyExcludeQuery(func() advice_kit_api.AdviceDefinition {
				return advice_kit_api.AdviceDefinition{
					AssessmentQueryExclude: tt.presentQuery,
				}
			}, tt.configuredExclude)(); !reflect.DeepEqual(got.AssessmentQueryExclude, tt.wantedQuery) {
				t.Errorf("applyExcludeQuery() = %v, want %v", got.AssessmentQueryExclude, tt.wantedQuery)
			}
		})
	}
}
