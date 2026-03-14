package goenv

import (
	"testing"
)

func TestGet(t *testing.T) {
	tests := []struct {
		name     string
		envVal   string
		expected Type
	}{
		{"empty defaults to prod", "", Prod},
		{"prod", "prod", Prod},
		{"dev", "dev", Dev},
		{"unknown defaults to prod", "staging", Prod},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(EnvVarName, tt.envVal)

			got := Get()
			if got != tt.expected {
				t.Errorf(
					"Get() = %q, want %q",
					got, tt.expected,
				)
			}
		})
	}
}

func TestIsProd(t *testing.T) {
	tests := []struct {
		name     string
		envVal   string
		expected bool
	}{
		{"empty is prod", "", true},
		{"prod is prod", "prod", true},
		{"dev is not prod", "dev", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(EnvVarName, tt.envVal)

			if got := IsProd(); got != tt.expected {
				t.Errorf(
					"IsProd() = %v, want %v",
					got, tt.expected,
				)
			}
		})
	}
}

func TestIsDev(t *testing.T) {
	tests := []struct {
		name     string
		envVal   string
		expected bool
	}{
		{"empty is not dev", "", false},
		{"prod is not dev", "prod", false},
		{"dev is dev", "dev", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv(EnvVarName, tt.envVal)

			if got := IsDev(); got != tt.expected {
				t.Errorf(
					"IsDev() = %v, want %v",
					got, tt.expected,
				)
			}
		})
	}
}
