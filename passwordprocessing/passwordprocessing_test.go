package passwordprocessing

import (
	"testing"
)

var shaTestCases = []struct {
	input string
	base64Sha string
}{
	{ "Password1", "70CCD9007338D6D81DD3B6271621B9CF9A97EA00" },
	{ "5TronG-p45sw()rd", "2CAF0B7082B8CA0F41C815F14D6E216602784895" },
	{ "letmein", "B7A875FC1EA228B9061041B7CEC4BD3C52AB3CE3" },
}

func TestStringToSha1(t *testing.T) {
	for _, testCase := range shaTestCases {
		t.Run(testCase.input, func(t *testing.T) {
			shaString := StringToSha1Hex(testCase.input)
			if shaString != testCase.base64Sha {
				t.Errorf("Password Encoding to SHA got %v expected %v", shaString, testCase.base64Sha)
			}
		})
	}
}