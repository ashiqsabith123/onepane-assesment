package helper

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/ashiqsabith123/onepane-assesment/models"
)

var results []models.Result
var mu sync.Mutex

func GetResults() *[]models.Result {
	return &results
}

func MakeResults(ch chan struct{}, posts []models.Posts, comments []models.Comments, users []models.Users) {

	wg := sync.WaitGroup{}

	for i := 0; i < len(posts); i++ {

		res := new(models.Result)
		res.PostID = posts[i].ID
		res.PostName = posts[i].Title
		res.Body = posts[i].Body

		wg.Add(1)
		go func() {
			defer wg.Done()
			for _, v := range comments {
				if v.PostID == posts[i].ID {
					res.CommentsCount++
				}
			}
		}()

		wg.Add(1)
		go func() {
			defer wg.Done()

			for _, v := range users {
				if v.ID == posts[i].UserID {
					res.UserName = v.Name

					break
				}
			}
		}()

		wg.Wait()
		mu.Lock()
		results = append(results, *res)
		mu.Unlock()

	}

	ch <- struct{}{}
}

func GetResponse(url string, responseType any) error {
	response, err := http.Get(url)

	if err != nil {
		return err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&responseType)
	if err != nil {
		return err
	}
	return nil

}
