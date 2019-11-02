package tempdirdao

import (
	"errors"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

const tempDirBase = "tmp/"
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

func (tfa *TempFileAccess) GetFullFilePath(filename string) (string, error) {
	dirContents, err := tfa.LoadDirContents()
	if err != nil {
		return "", err
	}
	pwd, _ := os.Getwd()
	path := pwd + "/" + tfa.tempDirLocation
	for i := range dirContents {
		if dirContents[i].Name() == filename {
			return filepath.Join(path, filename), nil
		}
	}
	return "", errors.New("file does not exist")
}

func (tfa *TempFileAccess) LoadDirContents() ([]os.FileInfo, error) {
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
	allFilesInTempDir, err := tfa.LoadDirContents()
	if err != nil {
		return err
	}
	for i := range(allFilesInTempDir) {
		fullFilePath, err := tfa.GetFullFilePath(allFilesInTempDir[i].Name())
		if err != nil {
			return err
		}
		err = os.Remove(fullFilePath)
		if err != nil {
			return err
		}
	}
	return os.Remove(tfa.tempDirLocation)
}
