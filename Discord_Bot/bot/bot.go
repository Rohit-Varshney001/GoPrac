package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rohitvarshney/GoPrac/discord_bot/config"
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
		fmt.Println(err.Error())
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

	content := strings.ToLower(m.Content)

	switch content {
	case "ping":
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	case "hello":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi there! 🙋")
	case "hii":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi there! 🙋")
	case "what's up":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Not much, just chilling. How about you? 😊")
	case "tell me a joke.":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Why don't scientists trust atoms? Because they make up everything! 😄")
	case "can you recommend a movie":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Sure! How about watching 'Inception'? It's a mind-bending thriller. 😁")
	case "how is the weather today":
		_, _ = s.ChannelMessageSend(m.ChannelID, "The weather today is sunny with a few clouds. Perfect for a stroll outside! 🫠")
	case "thanks":
		_, _ = s.ChannelMessageSend(m.ChannelID, "You're welcome! 🤝")
	case "bye":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Goodbye! Have a great day! 👌")
	case "help":
		_, _ = s.ChannelMessageSend(m.ChannelID, "Kya hua bsdk.? 🙋")
	case "what is your purpose?":
		_, _ = s.ChannelMessageSend(m.ChannelID, "I'm here to assist you with whatever you need! 🙂")

	default:
		_, _ = s.ChannelMessageSend(m.ChannelID, "Bhen ke laude bakchodi mat kar tameez ke question pooch.! 🖕")
	}

}
