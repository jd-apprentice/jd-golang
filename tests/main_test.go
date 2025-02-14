package tests

import (
	"jd-golang/src/config"
	"os"
	"reflect"
	"testing"
)

func TestGetConfig(t *testing.T) {
	for _, tt := range config_mock {
		t.Run(tt.name, func(t *testing.T) {
			for key, value := range tt.envVars {
				os.Setenv(key, value)
			}

			defer func() {
				for key := range tt.envVars {
					os.Unsetenv(key)
				}
			}()

			config := config.GetConfig()

			if !reflect.DeepEqual(config, tt.expectedConfig) {
				t.Errorf("GetConfig() = %+v, want %+v", config, tt.expectedConfig)
			}
		})
	}
}
