package lakefs

import (
	"bytes"
	"encoding/json"
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

// curl --user AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY \
// 		-X POST -H "Content-Type: application/json" \
//		--data '{
//			"name": "uncompressed",
//			"storage_namespace": "local://example-bucket/",
//			"default_branch": "main" }'
//		http://127.0.0.1:8000/api/v1/repositories
func CreateRepo(repoName string, storage_namespace string, default_branch string) (string, error) {
	URL := "http://127.0.0.1:8000/api/v1/repositories"
	reqBody, _ := json.Marshal(map[string]string{
		"name":              repoName,
		"storage_namespace": storage_namespace,
		"default_branch":    default_branch,
	})

	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 201:
		fmt.Printf("Repository %s seccessfully created\n", repoName)
	case 409:
		fmt.Printf("Repo %s already existing\n", repoName)
	}

	// if resp.StatusCode == 201 {
	// 	fmt.Printf("Repository %s seccessfully created\n", repoName)
	// } else {
	// 	fmt.Printf("Failed to create repo %s\n", repoName)
	// }
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return "", nil
}

// curl --user AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY -X POST -H "Content-Type: multipart/form-data" -F "content=@`pwd`/out.pdf" http://127.0.0.1:8000/api/v1/repositories/uncompressed/branches/main/objects?path=out.pdf

// func UploadFile() {
// 	URL := "http://127.0.0.1:8000/api/v1/repositories"
// 	reqBody, _ := json.Marshal(map[string]string{
// 		"name":              repoName,
// 		"storage_namespace": storage_namespace,
// 		"default_branch":    default_branch,
// 	})

// 	req, err := http.NewRequest("POST", URL, bytes.NewBuffer(reqBody))
// 	if err != nil {
// 		return "", err
// 	}
// 	req.Header.Set("Content-Type", "multipart/form-data")
// 	req.SetBasicAuth("AKIAIOSFODNN7EXAMPLE", "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()
// }

// curl --user AKIAIOSFODNN7EXAMPLE:wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY http://127.0.0.1:8000/api/v1/repositories/asd/refs/main/objects?path=cisco-flowtable.bib
