package puzzler

import (
	"errors"
	"io/ioutil"
	"time"
)

func GetAlgorithm() string {
	return algorithm
}

func GetDuration() time.Duration {
	return unhashingTime
}
func saveFile(message string, solutionPath string) {
	err := ioutil.WriteFile(solutionPath, []byte(message), 0644)
	if err != nil {
		println(message)
		panic(err.Error())
	}
}

func getHash(algo string, orig string) (string, error) {
	var hashed string

	switch algo {
	case "md5":
		hashed = Md5Hash(orig)
	case "sha1":
		hashed = Sha1Hash(orig)

	case "sha256":
		hashed = Sha256Hash(orig)
	case "sha512":
		hashed = Sha512Hash(orig)
	case "sha3256":
		hashed = Sha3256(orig)

	default:
		return hashed, errors.New("No appropiate algorithm given")
	}
	return hashed, nil
}
