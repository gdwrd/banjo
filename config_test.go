package banjo

import "testing"

func TestDefaultConfig(t *testing.T) {
	config := DefaultConfig()

	if config.host != "127.0.0.1" {
		t.Errorf("Host should be default")
	}

	if config.port != "4321" {
		t.Errorf("Port should be default")
	}

	if config.debug != false {
		t.Errorf("Debug filed should be default")
	}
}
