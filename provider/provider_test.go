package provider

import "testing"

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("provider schema is invalid: %s", err)
	}
}

// Variable values can hold secrets, so they must be redacted in plan output.
func TestVariableValuesAreSensitive(t *testing.T) {
	resources := Provider().ResourcesMap
	for _, name := range []string{"expo_app_variable", "expo_account_variable"} {
		resource, ok := resources[name]
		if !ok {
			t.Fatalf("resource %s is not registered", name)
		}
		if !resource.Schema["value"].Sensitive {
			t.Errorf("%s.value must be Sensitive", name)
		}
	}
}
