package conf

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pericles-luz/go-base/pkg/utils"
)

type ConfigBase struct {
	raw []byte
}

func (cfg *ConfigBase) ReadConfigurationFile(fileName string) ([]byte, error) {
	configPath := utils.GetBaseDirectory("config")
	log.Println("configPath:", configPath)
	fileName = strings.ReplaceAll(fileName, "..", "")
	fileName = strings.ReplaceAll(fileName, "/", "")
	filePath := fmt.Sprintf("%s/%s.json", configPath, fileName)
	if _, err := os.Stat(filePath); nil != err {
		return nil, err
	}
	json, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	cfg.raw = json
	return cfg.raw, nil
}
