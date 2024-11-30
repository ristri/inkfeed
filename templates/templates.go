package templates

import (
	"encoding/json"
	"html/template"
	"inkfeed/api"
	"inkfeed/models"
	"net/http"
	"strings"
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Templates struct {
	tmpl *template.Template
}

func LoadTemplates() (*Templates, error) {
	funcMap := template.FuncMap{
		"formatTime":       formatTime,
		"getNestedReplies": getNestedReplies,
		"hasReplies":       hasReplies,
		"markdownToHTML":   markdownToHTML,
	}

	tmpl := template.New("").Funcs(funcMap)

	templateFiles := []string{
		"templates/list.html",
		"templates/post.html",
		"templates/comments.html",
		"templates/styles.html",
		"templates/scripts.html",
	}

	for _, file := range templateFiles {
		_, err := tmpl.ParseFiles(file)
		if err != nil {
			return nil, err
		}
	}

	return &Templates{tmpl: tmpl}, nil
}

func (t *Templates) HandleSubreddit() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subredditPath := strings.TrimPrefix(r.URL.Path, "/r/")
		params := r.URL.Query()
		isTopSort := false
		if len(strings.SplitN(subredditPath, "/", 2)) > 1 {
			isTopSort = strings.Contains(strings.SplitN(subredditPath, "/", 2)[1], "top")
		}

		if subredditPath == "" {
			subredditPath = "popular" // Default subreddit
		}

		subredditName := strings.SplitN(subredditPath, "/", 2)[0]

		data, err := api.FetchRedditData(subredditPath, params)
		if err != nil {
			http.Error(w, "Error fetching Reddit data: "+err.Error(), http.StatusInternalServerError)
			return
		}

		params.Del("before")
		params.Del("after")
		params.Del("limit")
        params.Del("count")

		subredditData := models.SubredditData{
			Subreddit: subredditName,
			IsTopSort: isTopSort,
			Params:    params,
			Response:  *data,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := t.tmpl.ExecuteTemplate(w, "list.html", subredditData); err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (t *Templates) HandleComments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		permalink := strings.TrimPrefix(r.URL.Path, "/comments")

		post, comments, err := api.FetchComments(permalink)
		if err != nil {
			http.Error(w, "Error fetching comments: "+err.Error(), http.StatusInternalServerError)
			return
		}

		parts := strings.Split(permalink, "/")
		subreddit := parts[2]

		data := models.CommentData{
			Post:      post,
			Comments:  comments,
			Subreddit: subreddit,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := t.tmpl.ExecuteTemplate(w, "post.html", data); err != nil {
			http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func getNestedReplies(comment models.Comment) *models.CommentResponse {
	if comment.Replies == nil {
		return nil
	}

	// Handle empty string replies
	if _, ok := comment.Replies.(string); ok {
		return nil
	}

	// Try to convert to map first
	repliesMap, ok := comment.Replies.(map[string]interface{})
	if !ok {
		return nil
	}

	// Create a new CommentResponse
	var response models.CommentResponse

	// Marshal and unmarshal to convert the map to struct
	jsonBytes, err := json.Marshal(repliesMap)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		return nil
	}

	return &response
}

func markdownToHTML(md string) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	return template.HTML(markdown.ToHTML([]byte(md), p, nil))
}

func hasReplies(comment models.Comment) bool {
	replies := getNestedReplies(comment)
	return replies != nil && len(replies.Data.Children) > 0
}

func formatTime(timestamp float64) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("Jan 02, 2006 15:04 MST")
}
