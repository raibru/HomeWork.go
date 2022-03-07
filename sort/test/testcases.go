package test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

// CaseSort collect the sorttest data by input and expected
type CaseSort struct {
	Input  []int
	Expect []int
}

// ReadTestData read the fixture test data from file
func ReadTestData(filename string, css *[]CaseSort) error {
	testFile := filepath.Join("test-fixtures", filename)
	content, errFile := ioutil.ReadFile(testFile)
	if errFile != nil {
		return errFile
	}
	errConv := json.Unmarshal([]byte(content), css)
	if errConv != nil {
		return errConv
	}

	return nil
}
