// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// const (
// 	apiKey  = "cjuRu3aF4gEa4XSiX9WKdRRM"
// 	baseURL = "https://www.searchapi.io/api/v1/google-jobs"
// )

// type Job struct {
// 	Title       string `json:"title"`
// 	Company     string `json:"company_name"`
// 	Description string `json:"description"`
// 	Location    string `json:"location"`
// }

// type Response struct {
// 	Jobs []Job `json:"jobs"`
// }

// func searchJobs(query string) (*Response, error) {
// 	// Construct the URL with query parameters
// 	params := url.Values{}
// 	params.Add("query", query)
// 	params.Add("apikey", apiKey)

// 	reqURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

// 	// Make the HTTP request
// 	resp, err := http.Get(reqURL)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer resp.Body.Close()

// 	// Check the HTTP status code
// 	if resp.StatusCode != http.StatusOK {
// 		body, _ := ioutil.ReadAll(resp.Body)
// 		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, body)
// 	}

// 	// Read and parse the response body
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var response Response
// 	if err := json.Unmarshal(body, &response); err != nil {
// 		return nil, err
// 	}

// 	return &response, nil
// }

// func main() {
// 	query := "TN visa sponsor"
// 	response, err := searchJobs(query)
// 	if err != nil {
// 		log.Fatalf("Error searching for jobs: %v", err)
// 	}

// 	for _, job := range response.Jobs {
// 		fmt.Printf("Title: %s\nCompany: %s\nLocation: %s\nDescription: %s\n\n", job.Title, job.Company, job.Location, job.Description)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	baseURL := "https://www.searchapi.io/api/v1/search"
	queryParams := url.Values{
		"engine":   []string{"google_jobs"},
		"q":        []string{"TN visa sponsor engineer"},
		"location": []string{"California,United States"},
		"api_key":  []string{"cjuRu3aF4gEa4XSiX9WKdRRM"},
	}

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())
	req, _ := http.NewRequest("GET", fullURL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error making the request: %v", err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var jobsResponse struct {
		Jobs []struct {
			Title                 string `json:"title"`
			Company               string `json:"company"`
			Location              string `json:"location"`
			Description           string `json:"description"`
			CompanyWebResultsLink string `json:"company_web_results_link"`
			ShareLink             string `json:"sharing_link"`
			Thumbnail             string `json:"thumbnail"`
			JobId                 string `json:"job_id"`
		} `json:"jobs"`
	}

	if err := json.Unmarshal(body, &jobsResponse); err != nil {
		log.Fatalf("Error unmarshalling response: %v", err)
	}

	for i, job := range jobsResponse.Jobs {
		if i >= 5 {
			break
		}
		fmt.Printf("Title: %s\nCompany: %s\nLocation: %s\nDescription: %s\nCompany Web Link: %s\nShare Link: %s\nThumbnail: %s\nJob ID: %s\n\n",
			job.Title, job.Company, job.Location, job.Description, job.CompanyWebResultsLink, job.ShareLink, job.Thumbnail, job.JobId)
		fmt.Println("==============================================================")
	}
}
