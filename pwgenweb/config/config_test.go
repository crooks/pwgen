package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	testFile, err := os.CreateTemp("/tmp", "testcfg")
	if err != nil {
		t.Fatalf("Unable to create a temp file:  %v", err)
	}
	defer os.Remove(testFile.Name())
	testCfg := new(Config)
	testCfg.Standard.MaxSepLength = 3
	testCfg.WriteConfig(testFile.Name())

	cfg, err := ParseConfig(testFile.Name())
	if err != nil {
		t.Fatalf("ParseConfig returned: %v", err)
	}
	if cfg.Standard.MaxSepLength != 3 {
		t.Errorf("Unexpected cfg.Standard.MaxSepLength. Wanted=%d, Got=%d", testCfg.Standard.MaxSepLength, cfg.Standard.MaxSepLength)
	}
}
