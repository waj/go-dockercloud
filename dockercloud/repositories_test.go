package dockercloud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_ListRepositories(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("listimages.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/repo/" + repoSubsystemVersion + "/repository/"

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response RepositoryListResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/"

	test_response, err := ListRepositories()
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}

func Test_GetRepository(t *testing.T) {
	User = "test"
	ApiKey = "test"

	fake_response, err := MockupResponse("image.json")
	if err != nil {
		t.Fatal(err)
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, fake_response)
	}))

	defer server.Close()
	url := server.URL + "/api/repo/" + repoSubsystemVersion + "/repository/" + fake_image_name

	res, err := http.Get(url)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	var response Repository
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Fatal(err)
	}

	BaseUrl = server.URL + "/api/"
	test_response, err := GetRepository(fake_image_name)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(test_response, response) != true {
		t.Fatal("Invalid output")
	}
}
