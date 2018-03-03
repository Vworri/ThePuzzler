package puzzler

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

var puzzleCreationConfig unhahsedSolution

type unhahsedSolution struct {
	Input []solutionLayer `json:"input"`
}
type solutionLayer struct {
	Layer     int    `json:"layer"`
	Algorithm string `json:"algorithm"`
	Message   string `json:"message"`
}

func GetCreationConfig(congigPath string) unhahsedSolution {
	puzzleCreationConfig.OpenPuzzleCreator(congigPath)
	return puzzleCreationConfig
}

func (layers unhahsedSolution) CreatePuzzle(pp string) {
	var thisSolution string
	if len(layers.Input) > 20 {
		panic("Too mny layers for comfort. Please limit to 20 or less!")
	}
	for _, layer := range layers.Input {

		thisPuzzle, err := layer.CreatLayer(thisSolution)
		if err != nil {
			panic(err.Error())
		}
		thisSolution = thisPuzzle.join()

	}
	saveFile(thisSolution, pp)

}

func (pc *unhahsedSolution) OpenPuzzleCreator(fp string) {
	puzzleCreationFile, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err.Error())
	}
	json.Unmarshal(puzzleCreationFile, &pc)
}

func (s *solutionLayer) CreatLayer(prevMessage string) (Puzzle, error) {
	var solText string
	algo := s.Algorithm
	sol := s.Message + "\n" + prevMessage
	puzzle := *new(Puzzle)

	emptyhash, err := getHash(algo, "")

	if err != nil {
		panic(err.Error())
	}
	puzzle = append(puzzle, emptyhash)

	for _, character := range sol {
		solText = solText + string(character)
		hashed, err := getHash(algo, solText)
		if err != nil {
			panic(err.Error())
		}
		puzzle = append(puzzle, hashed)

	}
	return puzzle, nil

}

func (p *Puzzle) join() string {
	return strings.Join(*p, "\n")
}
