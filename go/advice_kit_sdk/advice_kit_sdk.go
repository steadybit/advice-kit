// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Steadybit GmbH

package advice_kit_sdk

import (
	"fmt"

	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"github.com/steadybit/extension-kit/exthttp"
	"github.com/steadybit/extension-kit/extutil"
)

var (
	registeredAdvice = make(map[string]struct{})
)

type AdviceFn func() advice_kit_api.AdviceDefinition

type AdviceConfig struct {
	DisableAdvice      bool     `json:"disableAdvice" required:"false" split_words:"true" default:"false"`
	ActiveAdviceList   []string `json:"activeAdviceList" required:"false" split_words:"true" default:"*"`
	AdviceExcludeQuery string   `json:"adviceExcludeQuery" required:"false" split_words:"true" default:""`
}

func (c AdviceConfig) IsActive(id string) bool {
	if c.DisableAdvice {
		return false
	}

	for _, el := range c.ActiveAdviceList {
		if el == "*" || el == id {
			return true
		}
	}

	return false
}

func RegisterAdvice(cfg AdviceConfig, fn AdviceFn) {
	id := fn().Id

	if !cfg.IsActive(id) {
		return
	}

	registeredAdvice[id] = struct{}{}
	exthttp.RegisterHttpHandler(fmt.Sprintf("GET /advice/%s", id), exthttp.GetterAsHandler(applyExcludeQuery(fn, cfg.AdviceExcludeQuery)))
}

func applyExcludeQuery(fn AdviceFn, query string) AdviceFn {
	if query == "" {
		return fn
	}

	return func() advice_kit_api.AdviceDefinition {
		a := fn()
		if a.AssessmentQueryExclude != nil && *a.AssessmentQueryExclude != "" {
			a.AssessmentQueryExclude = extutil.Ptr(fmt.Sprintf("(%s) OR (%s)", *a.AssessmentQueryExclude, query))
		} else {
			a.AssessmentQueryExclude = &query
		}
		return a
	}
}

// ClearRegisteredAdvice clears all registered advice - used for testing. Warning: This will not remove the registered routes from the http server.
func ClearRegisteredAdvice() {
	registeredAdvice = make(map[string]struct{})
}

// GetAdviceList returns a list of all root endpoints of registered advice.
func GetAdviceList() advice_kit_api.AdviceList {
	var result []advice_kit_api.DescribingEndpointReference

	for id := range registeredAdvice {
		result = append(result, advice_kit_api.DescribingEndpointReference{
			Method: advice_kit_api.GET,
			Path:   fmt.Sprintf("/advice/%s", id),
		})
	}

	return advice_kit_api.AdviceList{Advice: result}
}
