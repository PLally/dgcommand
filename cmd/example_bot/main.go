package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	// setup discord bot
	TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = session.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	// create and add command handlers
	rootHandler := dgcommand.NewCommandHandler()
	cmd := dgcommand.NewCommand("test <required> [optional] [vararg...]", testCommand)
	rootHandler.AddHandler(cmd.Name, cmd)

	prefixedRootHandler := dgcommand.WithPrefix(rootHandler, func(dgcommand.Context) string {
		return "!"
	})

	session.AddHandler(dgcommand.DiscordHandle(prefixedRootHandler))


	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}

func testCommand(ctx dgcommand.Context) {
	fmt.Println(strings.Join(ctx.Args, ", "))
}
