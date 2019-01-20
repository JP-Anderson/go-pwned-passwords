package pwnapi

import (
	"net/http"
	"fmt"	
 	"io/ioutil"
)

const PwnedPasswordsUrl = "https://api.pwnedpasswords.com/range/%v"

func GetPartialHashesForSha(input string) string {
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
	return str
}