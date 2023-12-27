package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ashiqsabith123/onepane-assesment/helper"
	"github.com/ashiqsabith123/onepane-assesment/models"
)

func GetPostDetails(w http.ResponseWriter, r *http.Request) {

	var comments []models.Comments
	var posts []models.Posts
	var users []models.Users

	urls := []string{
		"https://jsonplaceholder.typicode.com/comments",
		"https://jsonplaceholder.typicode.com/posts",
		"https://jsonplaceholder.typicode.com/users",
	}

	responces := []any{
		&comments,
		&posts,
		&users,
	}

	var wg sync.WaitGroup
	ch := make(chan error, 3)
	ch1 := make(chan struct{}, 2)

	for i, url := range urls {
		wg.Add(1)
		go func(u string, response interface{}, ch chan error) {
			defer wg.Done()
			err := helper.GetResponse(u, response)
			if err != nil {
				ch <- err
			}
		}(url, responces[i], ch)
	}

	wg.Wait()

	close(ch)

	for v := range ch {
		if v != nil {
			http.Error(w, "Internal Server Error"+v.Error(), http.StatusInternalServerError)
			return
		}
	}

	go helper.MakeResults(0, len(posts)/2, ch1, posts, comments, users)
	go helper.MakeResults(len(posts)/2, len(posts), ch1, posts, comments, users)

	<-ch1
	<-ch1

	close(ch1)

	jsonData, err := json.Marshal(helper.GetResults())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)

}
