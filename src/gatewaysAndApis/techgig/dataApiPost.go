package techgig

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/vkjayendravarma/techgig-benz-v2/env"
)

/*
make a post request to techGig data api
https://restmock.techgig.com/merc/

@requestBodyString string : this will be converted as *reader

@endPoint string

return response byteArray
*/
func PostRequest(requestBodyString string, endPoint string) []byte {

	// convert request body to reader
	requestBody := strings.NewReader(requestBodyString)

	// throw api call
	response, err := http.Post(env.TechGigDataApi+endPoint, "application/json", requestBody)
	// check error in api response
	if err != nil {
		return nil
	}

	//close api call after completing my business logic
	defer response.Body.Close()

	// convert response body to byte array
	responseBodyByteArray, _ := ioutil.ReadAll(response.Body)
	return responseBodyByteArray
}
