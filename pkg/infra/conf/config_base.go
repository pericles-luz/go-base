package conf

import (
	"errors"
	"fmt"
	"os"

	"github.com/pericles-luz/go-base/pkg/utils"
)

type ConfigBase struct{}

func (cfg *ConfigBase) ReadConfigurationFile(fileName string) ([]byte, error) {
	configPath := utils.GetBaseDirectory("config")
	filePath := fmt.Sprintf("%s/%s.json", configPath, fileName)
	if !utils.FileExists(filePath) {
		return nil, errors.New("configuration file does not exists: " + filePath)
	}
	json, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return json, nil
}
