package pwnapi

import (
	"net/http"
	"fmt"	
 	"io/ioutil"
	"strings"
)

type pwnedApiClient interface {
	getPwnedHashesForHashPrefix(firstFive string) []string
}

type actualPwnedApiClient struct {}

func (a actualPwnedApiClient) getPwnedHashesForHashPrefix(firstFive string) []string {
	url := fmt.Sprintf(PwnedPasswordsUrl, firstFive[0:5])
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

const PwnedPasswordsUrl = "https://api.pwnedpasswords.com/range/%v"
var api = new(actualPwnedApiClient)

func PasswordHashIsPwned(passwordSha1 string) bool {
	return passwordHashIsPwned(passwordSha1, api)
}

func passwordHashIsPwned(passwordSha1 string, apiClient pwnedApiClient) bool {
	firstFive := passwordSha1[0:5]
	pwnedHashes := api.getPwnedHashesForHashPrefix(firstFive)
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