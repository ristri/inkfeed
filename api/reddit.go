package api

import (
	"encoding/json"
	"fmt"
	"inkfeed/models"
	"io"
	"net/http"
	"net/url"
	"time"
)




func FetchComments(permalink string) (*models.RedditSubredditResponse, *models.CommentResponse, error) {
    url := "https://www.reddit.com" + permalink + ".json"
    client := &http.Client{Timeout: 10 * time.Second}
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, nil, err
    }
    
    resp, err := client.Do(req)
    if err != nil {
        return nil, nil, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, nil, err
    }
    
    var response []json.RawMessage
    if err := json.Unmarshal(body, &response); err != nil {
        return nil, nil, err
    }
    
    if len(response) < 2 {
        return nil, nil, fmt.Errorf("invalid response format")
    }
    
    var post models.RedditSubredditResponse
    var comments models.CommentResponse
    
    if err := json.Unmarshal(response[0], &post); err != nil {
        return nil, nil, err
    }
    
    if err := json.Unmarshal(response[1], &comments); err != nil {
        return nil, nil, err
    }
    
    return &post, &comments, nil
}

func FetchRedditData(subreddit string, param url.Values) (*models.RedditSubredditResponse, error) {
    url := "https://www.reddit.com/r/" + subreddit + ".json?" + param.Encode() 
    client := &http.Client{Timeout: 10 * time.Second}
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    
    var redditResp models.RedditSubredditResponse
    if err := json.Unmarshal(body, &redditResp); err != nil {
        return nil, err
    }

    return &redditResp, nil
}
