package pwnapi

import (
	"testing"
	"io/ioutil"
	"strings"
)

var passwordTestCases = []struct {
	password string
	sha1 string
	pathToHashFile string
	expected bool
}{
	{ "Password1", "70CCD9007338D6D81DD3B6271621B9CF9A97EA00", "test_data\\weak_password.txt", true },
	{ "5TronG-p45sw()rd", "2CAF0B7082B8CA0F41C815F14D6E216602784895", "test_data\\strong_password.txt", false },
}

type mockPwnedApi struct {
	mockedDataTxtFilePath string
}

func (m mockPwnedApi) getPwnedHashesForHashPrefix(firstFive string) []string {
	return readTextFileLinesIntoStringArray(m.mockedDataTxtFilePath)	
}

func TestPasswordCheckWithMockApi(t *testing.T) {
	for _, testCase := range passwordTestCases {
		t.Run(testCase.password, func(t *testing.T) {
			apiMock := mockPwnedApi{ mockedDataTxtFilePath : testCase.pathToHashFile }
			result := passwordHashIsPwned(testCase.sha1, apiMock)
			if result != testCase.expected {
				t.Errorf("Expected result for %v was %v but got %v", testCase.password, testCase.expected, result)
			}
		})
	}
}

func readTextFileLinesIntoStringArray(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		//Do something
	}
	return strings.Split(string(content), "\n")
}