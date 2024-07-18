package realtime

import "os"

// defaultOrEnv returns default or environment variable values for Centrifugo address and key.
func defaultOrEnv() (addr, key string) {
	// Addr is Centrifugo API endpoint.
	// For local: "http://host.docker.internal:8888/api"
	// [CENTRIFUGO_ADDR]
	addr = "https://centrifugo-xy3jfcnkpa-de.a.run.app/api"

	// Key is Centrifugo API key.
	// [CENTRIFUGO_KEY]
	key = ""

	if v, ok := os.LookupEnv("CENTRIFUGO_ADDR"); ok {
		addr = v
	}
	if v, ok := os.LookupEnv("CENTRIFUGO_KEY"); ok {
		key = v
	}

	return
}
