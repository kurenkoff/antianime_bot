package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	session, err := discordgo.New("Bot " + "TOKEN HERE")
	if err != nil {
		log.Fatal(err)
	}

	session.AddHandler(createMessage)
	defer session.Close()

	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func createMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if forDelete(m.Content) {
		s.ChannelMessageDelete(m.ChannelID, m.ID)	
	}
}

func forDelete(message string) bool {
	stopSubstr := []string{"anime", "bdo", "jojo"}

	stopSubstr = append(stopSubstr, []string{
		"аниме",
		"бдо",
		"блек дессерт",
		"блэк дессерт",
		"кпоп",
	}...)

	for _, substr := range stopSubstr {
		if strings.Contains(strings.ToLower(message), substr) {
			return true
		}
	}
	
	return false
}