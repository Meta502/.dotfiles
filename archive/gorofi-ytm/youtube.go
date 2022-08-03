package gorofiytm

import (
	"encoding/json"
	"net/url"
	"os"
	"regexp"

	"github.com/go-resty/resty/v2"
)

const SUGGESTION_URL = "https://suggestqueries.google.com/complete/search?gl=us&ds=yt&client=youtube&q="
const API_URL = "https://www.googleapis.com/youtube/v3"

type YoutubeConfig struct {
	ApiKey string `json:"apiKey"`
}

type YoutubeClient struct {
	client resty.Client
	config YoutubeConfig
}

type SearchPageInfo struct {
	TotalResults int `json:"totalResults"`
	ResultsPerPage int `json:"resultsPerPage"`
}

type SearchListResponse struct {
	Items []SearchResult `json:"items"`
	PageInfo SearchPageInfo `json:"pageInfo"`
}

type VideoID struct {
	Kind string `json:"kind"`
	VideoID string `json:"videoId"`
}

type Thumbnail struct {
	URL string `json:"url"`
	Width int `json:"width"`
	Height int `json:"height"`
}

type Thumbnails struct {
	Default Thumbnail `json:"default"`
	Medium Thumbnail `json:"medium"`
	High Thumbnail `json:"high"`
}

type VideoSnippet struct {
	PublishedAt string `json:"publishedAt"`
	ChannelID string `json:"channelId"`
	Title string `json:"title"`
	Thumbnails Thumbnails `json:"thumbnails"`
	ChannelTitle string `json:"channelTitle"`
}

type SearchResult struct {
	Kind string `json:"kind"`
	ID VideoID `json:"id"`
	Snippet VideoSnippet `json:"snippet"`
}

func NewYoutubeClient() (*YoutubeClient, error) {
	// Open our jsonFile
	apiKey := os.Getenv("API_KEY")

	return &YoutubeClient{
		client: *resty.New(),
		config: YoutubeConfig{
			ApiKey: apiKey,
		},
	}, nil
}

func (y *YoutubeClient) GetSuggestions(query string) []Line {
	regex := regexp.MustCompile(`\[\"(.*?)\"`)
	resp, err := y.client.R().Get(SUGGESTION_URL + url.QueryEscape(query))

	if err != nil {
		return []Line{
			{
				Text: query,
			},
		}
	}

	matches := regex.FindAllString(string(resp.Body()), -1)

	results := []Line{
		{
			Text: query,
		},
	}

	for _, entry := range matches {
		results = append(results, Line{
			Text: entry[2 : len(entry)-1],
		})
	}

	return results
}

func (y *YoutubeClient) GetSearchResults(query string) (*SearchListResponse, error) {
	resp, reqErr := y.client.R().SetQueryParams(map[string]string {
		"q": query,
		"part": "snippet",
		"type": "video",
		"key": y.config.ApiKey,
	}).Get(API_URL + "/search")

	if reqErr != nil {
		return nil, reqErr
	}

	body := resp.Body()

	var searchResponse SearchListResponse

	parseErr := json.Unmarshal(body, &searchResponse)

	if parseErr != nil {
		return nil, parseErr
	}

	return &searchResponse, nil
}
