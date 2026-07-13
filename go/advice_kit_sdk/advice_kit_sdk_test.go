// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Steadybit GmbH

package advice_kit_sdk

import (
	"reflect"
	"testing"

	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"github.com/stretchr/testify/assert"
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
			presentQuery: new("target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'"),
			wantedQuery:  new("target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'"),
		},
		{
			name:              "configured exclude, no present query",
			configuredExclude: "k8s.namespace = 'kube-system'",
			wantedQuery:       new("k8s.namespace = 'kube-system'"),
		},
		{
			name:              "configured exclude, present query",
			presentQuery:      new("target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment'"),
			configuredExclude: "k8s.namespace = 'kube-system'",
			wantedQuery:       new("(target.type = 'com.steadybit.extension_kubernetes.kubernetes-deployment') OR (k8s.namespace = 'kube-system')"),
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

func TestAdviceConfig_IsActive(t *testing.T) {
	type fields struct {
		DisableAdvice      bool
		ActiveAdviceList   []string
		InactiveAdviceList []string
		AdviceExcludeQuery string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "should deny any advice - disabled",
			fields: fields{
				DisableAdvice:    true,
				ActiveAdviceList: []string{"*"},
			},
			args: args{id: "test"},
			want: false,
		},
		{
			name:   "should deny advice - not enabled",
			fields: fields{},
			args:   args{id: "test"},
			want:   false,
		},
		{
			name: "should deny advice - disabled ",
			fields: fields{
				ActiveAdviceList:   []string{"*"},
				InactiveAdviceList: []string{"test"},
			},
			args: args{id: "test"},
			want: false,
		},
		{
			name: "should allow advice - enabled ",
			fields: fields{
				ActiveAdviceList: []string{"test"},
			},
			args: args{id: "test"},
			want: true,
		},
		{
			name: "should allow advice - all enabled ",
			fields: fields{
				ActiveAdviceList: []string{"*"},
			},
			args: args{id: "test"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := AdviceConfig{
				DisableAdvice:      tt.fields.DisableAdvice,
				ActiveAdviceList:   tt.fields.ActiveAdviceList,
				InactiveAdviceList: tt.fields.InactiveAdviceList,
				AdviceExcludeQuery: tt.fields.AdviceExcludeQuery,
			}
			assert.Equalf(t, tt.want, c.IsActive(tt.args.id), "IsActive(%v)", tt.args.id)
		})
	}
}
