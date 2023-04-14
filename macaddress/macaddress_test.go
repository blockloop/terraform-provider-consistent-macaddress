package macaddress

import (
	"testing"
)

func TestGenerateMacAddress(t *testing.T) {
	testCases := []struct {
		seed            string
		expected string
	}{
		{"test-seed", "02:BF:E7:1B:97:A2"},
		{"another-seed", "02:28:C4:77:2F:74"},
		{"skjdajsdfljashdf", "02:00:22:48:CF:4C"},
	}

	for _, tc := range testCases {
		t.Run(tc.seed, func(t *testing.T) {
			macAddress, err := generateMacAddress(tc.seed)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if macAddress != tc.expected {
				t.Errorf("Expected MAC address to match pattern %s, but got %s", tc.expected, macAddress)
			}
		})
	}
}
