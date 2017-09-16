package translate

// TODO: TranslationConfig
// e.g. /content as basedir
// e.g. mainContent
// e.g. slug

// e.g. homepage -> _index.md
// etc

import (
	"os"

	"github.com/naoina/toml"
)

type TransConfig struct {
	Encoding string
}

func LoadConfig() TransConfig {
	fileName := "extract-config.toml"
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return TransConfig{}
	}

	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var config TransConfig
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}

	return config
}
