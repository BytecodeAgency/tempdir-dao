package tempdirdao

import (
	"io/ioutil"
	"testing"
)

func TestNewTempFileAccess (t *testing.T) {
	_, err := NewTempFileAccess()
	if err != nil {
		t.Error("Fail on DAO creation: " + err.Error())
	}
}

func TestTempFileAccess_SaveAndLoadFile (t *testing.T) {
	testFilename := "testfilename.ext"
	testFileContents := "Some test contents"
	dao, _ := NewTempFileAccess()
	err := dao.SaveFile(testFilename, []byte(testFileContents), 0644)
	if err != nil {
		t.Error("Error while saving file: " + err.Error())
	}
	loadedFileContents, err := dao.LoadFileContents(testFilename)
	if err != nil {
		t.Error("Error while loading file: " + err.Error())
	}
	if string(loadedFileContents) != testFileContents {
		t.Errorf("The loaded file content is different from the test file content, expected %s and got %s", testFileContents, string(loadedFileContents))
	}
}

func TestTempFileAccess_FullFileName (t *testing.T) {
	testFilename := "testfilename2.ext"
	testFileContents := "Some test contents"
	dao, _ := NewTempFileAccess()
	_ = dao.SaveFile(testFilename, []byte(testFileContents), 0664)

	// It should not give an err
	fullFileName, err := dao.GetFullFilePath(testFilename)
	if err != nil {
		t.Error("Got error while generating fullFileName: " + err.Error())
	}

	// The full filename should be longer than the input filename
	if len(fullFileName) <= len(testFilename) {
		t.Errorf("Full filename (%s) is shorter or as long as filename (%s)", fullFileName, testFilename)
	}

	// The given file path should be loaded and should contain the correct file contents
	loadedFullNameFileContents, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		t.Error("Error while loading file with full filename: " + err.Error())
	}
	if string(loadedFullNameFileContents) != testFileContents {
		t.Errorf("The loaded full name file contents (%s) differ from the test content (%s)", string(loadedFullNameFileContents), testFileContents)
	}

	// It should return an err when the asked file does not exist
	fullFileName, err = dao.GetFullFilePath("nonexistentfile.ext")
	if err == nil {
		t.Errorf("Getting full name of file does not trigger error, filename given was %s", fullFileName)
	}
}

func TestTempFileAccess_RemoveTempFileAccess(t *testing.T) {
	testFilename := "testfilename2.ext"
	testFileContents := "Some test contents"
	dao, _ := NewTempFileAccess()

	_ = dao.SaveFile(testFilename, []byte(testFileContents), 0644)
	fullFileName, _ := dao.GetFullFilePath(testFilename)

	// Now the files should be on disk
	loadedFullNameFileContents, err := ioutil.ReadFile(fullFileName)
	if err != nil {
		t.Error("Error while loading file with full filename: " + err.Error())
	}
	if string(loadedFullNameFileContents) != testFileContents {
		t.Errorf("The loaded full name file contents (%s) differ from the test content (%s)", string(loadedFullNameFileContents), testFileContents)
	}

	// No err should be returned when removing the DAO dir
	err = dao.RemoveTempFileAccess()
	if err != nil {
		t.Error("Error while removing DAO directory: " + err.Error())
	}

	// Now the files should be removed
	loadedFullNameFileContents, err = ioutil.ReadFile(fullFileName)
	if err == nil {
		t.Error("An err should be returned if the files are non existent")
	}
	if len(loadedFullNameFileContents) > 0 {
		t.Error("The file contents should be 0, received " + string(len(loadedFullNameFileContents)))
	}
}
