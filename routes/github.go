package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"minotor/data"
	"minotor/discord"
	"minotor/utils"
	"time"
)

func GetNewOsmoPullRequest(c *gin.Context) {
	var PullRequests []data.PullRequest
	var Result [][]byte

	resp, err := utils.DoRequest("GET", "https://api.github.com/repos/osmosis-labs/osmosis-frontend/pulls", nil)
	if err != nil {
		c.String(resp.StatusCode, fmt.Sprintf("%s error on GetNewOsmoPullRequest -> DoRequest", err.Error()))
	}
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &PullRequests)
	if err != nil {
		c.String(resp.StatusCode, fmt.Sprintf("%s error on GetNewOsmoPullRequest -> Unmarshal", err.Error()))
	}
	var MergedAt string
	var now = time.Now()
	for _, PullRequest := range PullRequests {
		for _, Label := range PullRequest.Labels {
			if Label.Name == "pool/asset" {
				if now.Sub(PullRequest.CreatedAt) < 10*time.Minute || (PullRequest.MergedAt.IsZero() == false && now.Sub(PullRequest.MergedAt) < 10*time.Minute) {
					if PullRequest.MergedAt.IsZero() {
						MergedAt = "Not Yet Merged"
					} else {
						MergedAt = PullRequest.MergedAt.String()
					}
					Result = append(Result, []byte(PullRequest.HtmlUrl))
					discord.SendDiscordMsgAboutPool(PullRequest.Title, MergedAt, PullRequest.CreatedAt.String(), PullRequest.HtmlUrl)
				}
			}
		}
	}
	ResultJson, _ := json.Marshal(Result)
	c.String(201, string(ResultJson))
}
