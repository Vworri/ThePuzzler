package puzzler

import (
	"io/ioutil"
	"os"
	"strings"
)

func OpenPuzzle(filename string) Puzzle {
	puzzlefile, err := os.Open(filename)
	defer puzzlefile.Close()
	if err != nil {
		panic(err.Error())
	}
	puzzle, err := ioutil.ReadAll(puzzlefile)
	if err != nil {
		panic(err.Error())
	}
	p := strings.Split(string(puzzle), "\n")

	return p
}

func findMessage(hashed string, unhahsed string, algo string) (string, error) {
	for _, character := range prinable {
		expText := unhahsed + string(character)

		expHash, err := getHash(algorithm, expText)
		if err != nil {
			return expText, err
		}

		if expHash == hashed {
			return expText, nil
		}

	}
	return unhahsed, nil

}
