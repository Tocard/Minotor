package discord

import (
	"encoding/json"
	"fmt"
	"github.com/bensch777/discord-webhook-golang"
	"log"
	"minotor/config"
	"time"
)

func SendDiscordMsgAboutPool(Title, MergedAt, CreatedAt, Url string) {

	embed := discordwebhook.Embed{
		Title:     "New pull request with pool/asset labels",
		Color:     10181046,
		Url:       Url,
		Timestamp: time.Now(),
		Fields: []discordwebhook.Field{
			discordwebhook.Field{
				Name:   "Pull Request Name:",
				Value:  Title,
				Inline: false,
			},
			discordwebhook.Field{
				Name:   "Merged At:",
				Value:  fmt.Sprintf("%s", MergedAt),
				Inline: false,
			},
			discordwebhook.Field{
				Name:   "Created At:",
				Value:  fmt.Sprintf("%s", CreatedAt),
				Inline: false,
			},
		},
		Footer: discordwebhook.Footer{
			Text: "Provided By LFDM for LFDM",
		},
	}
	SendEmbed(config.Cfg.DiscordDefiMinotor, "Minotor | NEW OSMO POOL", config.Cfg.DiscordDefiMinotorLogo, "<@&1094385766655860766>", embed)

}

func SendEmbed(link, name, avatar, grouptag string, embeds discordwebhook.Embed) {

	hook := discordwebhook.Hook{
		Username:   name,
		Avatar_url: avatar,
		Content:    grouptag,
		Embeds:     []discordwebhook.Embed{embeds},
	}

	payload, err := json.Marshal(hook)
	if err != nil {
		log.Println(err)
	}
	err = discordwebhook.ExecuteWebhook(link, payload)
	if err != nil {
		log.Printf("Error on SendEmbed %s", err.Error())
	}
}
