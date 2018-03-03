package main

import (
	"KyrusTech/puzzler"
	"KyrusTech/puzzler_help"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type fullSolution struct {
	Output []solution `json:"output"`
}

type solution struct {
	Layer     string         `json:"layer"`
	Algorithm string         `json:"algorithm"`
	Time      string         `json:"time"`
	Question  string         `json:"question"`
	Body      puzzler.Puzzle `json:"body"`
}

func main() {
	creationMode("puzzlconfig.json", "mypuzzle.txt")
}

func userHandler() {

	mode := flag.String("m", "s", "c = create a puzzle, s = solve a puzzle")
	puzzlepath := flag.String("pp", "Nextsolution.txt", "Path of the puzzle")
	maxLayers := flag.Int("ln", 10, "layers of hashes you are expecting")
	solFilePath := flag.String("sp", "Solution.json", "desired filepath of solution")
	createConfig := flag.String("config", "puzzlconfig.json", "filepath for the congiguration json for puzzle creation")
	messgaeIsFile := flag.Bool("mf", false, "if the \"messgae\" is a file")
	littleEndian := flag.Bool("le", false, "wether the file is little endian ( used when -mf = true )")
	if *messgaeIsFile {
		panic("feaure not availavle.. yet")
	}
	if *littleEndian {
		panic("feaure not availavle.. yet")
	}
	flag.Parse()
	if *maxLayers >= 20 {
		panic("To many layers! What?! Are you crazy?")
	}

	switch *mode {
	case "c":

		creationMode(*createConfig, *puzzlepath)

	case "s":
		solveMode(*puzzlepath, maxLayers, solFilePath)
	case "h":
		puzzler_help.Help()
	default:
		panic("Please use a mode flag \n e.g \n\t -m=c -config=myconfigfil.txt -pp=puzzle.txt ")
	}

}

func solveMode(puzzlepath string, maxLayers *int, solutionPath *string) {
	problem := puzzler.OpenPuzzle(puzzlepath)
	sol := russianDollUnhasher(maxLayers, problem)
	sol.saveSolution(solutionPath)

}
func creationMode(config string, puzzlePath string) {

	layerConfig := puzzler.GetCreationConfig(config)
	layerConfig.MessageAsFile()
	layerConfig.CreatePuzzle(puzzlePath)

}

func russianDollUnhasher(ln *int, p puzzler.Puzzle) fullSolution {
	fullSol := *new(fullSolution)
	for i := 1; i <= *ln; i++ {
		sol := *new(solution)

		s := p.UnhashPuzzleReturnSolution()
		sol.Question = s[0]
		sol.Algorithm = puzzler.GetAlgorithm()
		sol.Time = puzzler.GetDuration().String()
		sol.Body = s[1:]
		sol.Layer = fmt.Sprintf("%d", i)
		fullSol.Output = append(fullSol.Output, sol)
		if len(s) > 2 {
			p = s
		} else {
			return fullSol
		}

	}
	return fullSol
}

func (f *fullSolution) saveSolution(fp *string) {
	marsh, err := json.Marshal(f)
	if err != nil {
		panic(err.Error())
	}
	write_err := ioutil.WriteFile(*fp, marsh, 0664)
	if write_err != nil {
		panic(write_err.Error())
	}
}
