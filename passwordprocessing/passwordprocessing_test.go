package passwordprocessing

import (
	"testing"
	"encoding/hex"
	"strings"
)

var shaTestCases = []struct {
	input string
	base64Sha string
}{
	{ "Password1", "70CCD9007338D6D81DD3B6271621B9CF9A97EA00" },
}

func TestStringToSha1(t *testing.T) {
	for _, testCase := range shaTestCases {
		t.Run(testCase.input, func(t *testing.T) {
			output := PasswordStringToSHABytes(testCase.input)
			shaString := strings.ToUpper(hex.EncodeToString(output))
			if shaString != testCase.base64Sha {
				t.Errorf("Password Encoding to SHA got %v expected %v", shaString, testCase.base64Sha)
			}
		})
	}
}