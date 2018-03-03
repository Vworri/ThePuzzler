package puzzler

import (
	"strings"
	"time"
)

type Puzzle []string

func (p *Puzzle) UnhashPuzzleReturnSolution() Puzzle {
	s := time.Now()
	emptyHashes.fillEmptyHashStruct()
	emptyHashes.fillEmptyHashStruct()
	p.findAlgo()

	var message string
	for _, line := range *p {
		res, err := findMessage(line, message, algorithm)
		if err != nil {
			panic(err)
		}
		message = res

	}
	solution := *new(Puzzle)
	unhashingTime = time.Since(s)
	solution = strings.Split(string(message), "\n")
	return solution
}

func (p *Puzzle) findAlgo() {

	for _, l := range *p {
		if l == emptyHashes.emptymd5 {
			algorithm = "md5"
			return
		}
		if l == emptyHashes.emptysha1 {
			algorithm = "sha1"
			return
		}
		if l == emptyHashes.emptysha256 {
			algorithm = "sha256"
			return
		}
		if l == emptyHashes.emptysha512 {
			algorithm = "sha512"
			return
		}
		if l == emptyHashes.emptysha3256 {
			algorithm = "sha3256"
			return
		}
	}
	panic("No algorythm found")

}
