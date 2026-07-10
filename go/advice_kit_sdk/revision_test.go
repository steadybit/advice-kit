// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2026 Steadybit GmbH

package advice_kit_sdk

import (
	"testing"

	"github.com/steadybit/advice-kit/go/advice_kit_api"
	"github.com/steadybit/extension-kit/exthttp"
	"github.com/stretchr/testify/assert"
)

func TestRegisterAdviceBumpsRevision(t *testing.T) {
	ClearRegisteredAdvice()
	t.Cleanup(ClearRegisteredAdvice)

	before := exthttp.Revision()
	RegisterAdvice(AdviceConfig{ActiveAdviceList: []string{"*"}}, func() advice_kit_api.AdviceDefinition {
		return advice_kit_api.AdviceDefinition{Id: "com.steadybit.advice.test"}
	})
	assert.NotEqual(t, before, exthttp.Revision(), "RegisterAdvice must bump the index revision")
}

func TestRegisterInactiveAdviceDoesNotBumpRevision(t *testing.T) {
	ClearRegisteredAdvice()
	t.Cleanup(ClearRegisteredAdvice)

	before := exthttp.Revision()
	RegisterAdvice(AdviceConfig{DisableAdvice: true}, func() advice_kit_api.AdviceDefinition {
		return advice_kit_api.AdviceDefinition{Id: "com.steadybit.advice.disabled"}
	})
	assert.Equal(t, before, exthttp.Revision(), "an inactive advice registers nothing and must not bump the revision")
}

func TestClearRegisteredAdviceBumpsRevision(t *testing.T) {
	before := exthttp.Revision()
	ClearRegisteredAdvice()
	assert.NotEqual(t, before, exthttp.Revision(), "ClearRegisteredAdvice must bump the index revision")
}
