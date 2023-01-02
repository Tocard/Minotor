package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"minotor/config"
	"minotor/data"
	"minotor/utils"
)

func CreateGrafanaUser(c *gin.Context) {
	username := c.Param("user")
	password := utils.RandStringBytesMaskImprSrcUnsafe(25)
	email := fmt.Sprintf("%s@ether-source.fr", username)
	User := data.GrafanaUser{Name: username, Email: email, Login: username, Password: password, OrgId: 7}
	url := fmt.Sprintf("%s/api/admin/users", config.Cfg.GrafanaUrl)
	resp, err := utils.DoRequestAuth("POST", url, User)
	if err != nil {
		c.String(resp.StatusCode, fmt.Sprintf("%s error on CreateGrafanaUser", err))
		return
	}
	if resp.StatusCode != 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		c.String(resp.StatusCode, string(body))
		return
	}

	userJson, _ := json.Marshal(User)
	c.String(200, string(userJson))

}
