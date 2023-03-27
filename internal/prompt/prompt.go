package prompt

import (
	"gopkg.in/yaml.v3"
	"os"
)

func WriteToFile(fileName string, b Bundle) error {
	outBytes, err := yaml.Marshal(b)
	if err != nil {
		return err
	}
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(outBytes)
	if err != nil {
		return err
	}
	return nil
}
