package models

import "net/url"

type CommentData struct {
    Post      *RedditSubredditResponse
    Comments  *CommentResponse
    Subreddit string
}

type RedditSubredditResponse struct {
    Kind string `json:"kind"`
    Data struct {
        Children []struct {
            Kind string     `json:"kind"`
            Data RedditPost `json:"data"`
        } `json:"children"`
        After    string `json:"after"`
        Before   string `json:"before"`
    } `json:"data"`
}


type SubredditData struct {
		Subreddit string
        IsTopSort bool
        Params url.Values
		Response RedditSubredditResponse
}


type RedditPost struct {
    Title        string  `json:"title"`
    Selftext     string  `json:"selftext"`
    URL          string  `json:"url"`
    Thumbnail    string  `json:"thumbnail"`
    Score        int     `json:"score"`
    Author       string  `json:"author"`
    Created      float64 `json:"created_utc"`
    Permalink    string  `json:"permalink"`
    NumComments  int     `json:"num_comments"`
    Subreddit    string  `json:"subreddit"`
}

type CommentResponse struct {
    Kind string `json:"kind"`
    Data struct {
        Children []struct {
            Kind string  `json:"kind"`
            Data Comment `json:"data"`
        } `json:"children"`
        After  string `json:"after"`
        Before string `json:"before"`
    } `json:"data"`
}

type Comment struct {
    Author    string      `json:"author"`
    Body      string      `json:"body"`
    Created   float64     `json:"created_utc"`
    Score     int         `json:"score"`
    Replies   interface{} `json:"replies"` // Changed back to interface{} to handle both string and object
    ID        string      `json:"id"`
}

type Replies struct {
    Data struct {
        Children []struct {
            Kind string  `json:"kind"`
            Data Comment `json:"data"`
        } `json:"children"`
    } `json:"data"`
}


