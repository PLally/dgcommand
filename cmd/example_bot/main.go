package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/plally/dgcommand"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		log.Fatal(err.Error())
	}

	handler := dgcommand.NewCommandHandler()
	cmd := dgcommand.NewCommand("test <required> [optional] [vararg...]", testCommand)
	session.AddHandler(handler.DiscordHandle)
	handler.AddHandler(cmd.Name, cmd)

	err = session.Open()
	if err != nil {
		log.Fatal("error opening connection,", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	session.Close()
}

func testCommand(ctx dgcommand.CommandContext) {
	fmt.Println(ctx.Args)
}
