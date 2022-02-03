package lakefs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// curl command to set up LakeFS
//	curl --header "Content-Type: application/json" -X \
//		POST --data '{"username": "ramzi", "key": {"access_key_id": "AKIAIOSFODNN7EXAMPLE", "secret_access_key": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"}}' \
//  	http://127.0.0.1:8000/api/v1/setup_lakefs

func SetupLakeFS() (bool, error) {

	LAKEFS_URL := "http://127.0.0.1:8000/api/v1/setup_lakefs"

	// TODO : replace values with ENV variables
	var jsonStr = []byte(`{
		"username": "lakefs",
		"key": {
			"access_key_id": "AKIAIOSFODNN7EXAMPLE",
			"secret_access_key": "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
		}
	}`)
	req, err := http.NewRequest("POST", LAKEFS_URL, bytes.NewBuffer(jsonStr))

	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("LakeFS set up successfully")
	} else if resp.StatusCode == 409 {
		fmt.Println("LakeFS already set up")
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return true, nil
}

func createRepo() {

}

func uploadFile() {

}
