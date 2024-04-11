package bot

import (
	"context"
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/rohitvarshney/GoPrac/discord_bot/config"
	openai "github.com/sashabaranov/go-openai"
)

var BotId string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println("here" + err.Error())
		return
	}

	BotId = user.ID

	goBot.AddHandler(messegeHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is Running...")

}

func messegeHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
		return
	} else {
		response, err := getAIResponse(m.Content)
		if err != nil {
			fmt.Println(err.Error())
			_, _ = s.ChannelMessageSend(m.ChannelID, "Sorry, I couldn't generate a response at the moment.")
			return
		}
		_, _ = s.ChannelMessageSend(m.ChannelID, response)
		return
	}

}

func getAIResponse(message string) (string, error) {
	client := openai.NewClient(config.OpenAIAPIKey)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)
	if err != nil {
		return "", err
	}

	res := resp.Choices[0].Message.Content
	return res, nil
}
