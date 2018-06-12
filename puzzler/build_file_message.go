package puzzler

import (
	"encoding/base32"
	"io/ioutil"
	"os"
)

func openMessageFile(filename string) string {
	helpfile, err := os.Open(filename)
	defer helpfile.Close()
	if err != nil {
		panic(err.Error())
	}
	helptxt, err := ioutil.ReadAll(helpfile)
	if err != nil {
		panic(err.Error())
	}
	encodedFile := base32.StdEncoding.EncodeToString(helptxt)
	return encodedFile

}

func (s *unhahsedSolution) MessageAsFile() {

	for _, layer := range s.Input {
		if _, err := os.Stat(layer.Message); err == nil {
			layer.Message = layer.Message + "\n" + openMessageFile(layer.Message)

		}
	}
}
