package realtime

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_defaultOrEnv(t *testing.T) {
	const (
		expectedAddr = "http://example.com:8888/api"
		expectedKey  = "testkey"
	)
	defaultAddr, defaultKey := defaultOrEnv()

	// Prepare environment variables for testing
	tests := []struct {
		envVars      map[string]string // environment variables to set for the test
		expectedAddr string            // expected value of addr after defaultOrEnv
		expectedKey  string            // expected value of key after defaultOrEnv
	}{
		{map[string]string{"CENTRIFUGO_ADDR": expectedAddr, "CENTRIFUGO_KEY": expectedKey}, expectedAddr, expectedKey},
		{map[string]string{"CENTRIFUGO_ADDR": expectedAddr}, expectedAddr, defaultKey},
		{map[string]string{}, defaultAddr, defaultKey},
	}

	for _, test := range tests {
		// Set environment variables for the test
		os.Clearenv()
		for k, v := range test.envVars {
			os.Setenv(k, v)
		}

		// Call defaultOrEnv
		addr, key := defaultOrEnv()

		assert.Equal(t, test.expectedAddr, addr)
		assert.Equal(t, test.expectedKey, key)
	}
}
