package tempdirdao

import (
	"io/ioutil"
	"math/rand"
	"os"
)

const tempDirBase = "./tmp/"
const randomDirNameLength = 16
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

type TempFileAccess struct {
	tempDirLocation string
}

func generateRandomDirName() string {
	dirName := make([]rune, randomDirNameLength)
	for i := range dirName {
		dirName[i] = letters[rand.Intn(len(letters))]
	}
	return string(dirName)
}

func NewTempFileAccess() (*TempFileAccess, error) {
	dirName := generateRandomDirName()
	tempDirLocation := tempDirBase + dirName + "/"
	errCreatingDir := os.MkdirAll(tempDirLocation, 0755)
	return &TempFileAccess{
		tempDirLocation: tempDirLocation,
	}, errCreatingDir
}

func (tfa *TempFileAccess) LoadDirContents(filename string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(tfa.tempDirLocation)
}

func (tfa *TempFileAccess) LoadFileContents(filename string) ([]byte, error) {
	fileLocation := tfa.tempDirLocation + filename
	return ioutil.ReadFile(fileLocation)
}

func (tfa *TempFileAccess) SaveFile(filename string, fileContents []byte, permissions os.FileMode) error {
	fileLocation := tfa.tempDirLocation + filename
	return ioutil.WriteFile(fileLocation, fileContents, permissions)
}

func (tfa *TempFileAccess) RemoveTempFileAccess() error {
	return os.Remove(tfa.tempDirLocation)
}
