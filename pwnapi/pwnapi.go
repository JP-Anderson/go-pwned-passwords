package pwnapi

import (
	"net/http"
	"fmt"	
 	"io/ioutil"
	"strings"
)

const PwnedPasswordsUrl = "https://api.pwnedpasswords.com/range/%v"

func PasswordHashIsPwned(passwordSha1 string) bool {
	firstFive := passwordSha1[0:5]
	pwnedHashes := getPartialHashesForSha(firstFive)
	partialHash := passwordSha1[5:]
	
	hashIsPwned := false
	for _, hashCount := range pwnedHashes {
		hash := strings.Split(hashCount, ":")[0]
		if partialHash == hash {
			hashIsPwned = true
		}
	}
	
	return hashIsPwned
}

func getPartialHashesForSha(input string) []string {
	url := fmt.Sprintf(PwnedPasswordsUrl, input[0:5])
	resp, err := http.Get(url)
	if err != nil {
		//todo
	}
	
	defer resp.Body.Close()
	
	data, err2 := ioutil.ReadAll(resp.Body)
	
	if err2 != nil {
		//todo
	}
	
	str := string(data)
	
	lines := strings.Split(str, "\n")
	
	return lines
}