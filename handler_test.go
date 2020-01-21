package dgcommand

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"strings"
	"testing"
)

var testFullBot = flag.Bool("fullbot", false, "Should we create a full discordbot for testing")

func TestFullBot(t *testing.T) {
	if !*testFullBot{
		return
	}

	TOKEN := os.Getenv("DISCORD_BOT_TOKEN")
	session, err := discordgo.New("Bot " + TOKEN)
	if err != nil {
		t.Fatal(err.Error())
	}

	handler := NewCommandHandler()
	handler.AddCommandFunc("echo <name>", testCommand)
	session.AddHandler(handler.DiscordHandle)
}

func testCommand(ctx CommandContext) {

}

func ExampleParseCommand() {
	args := parseCommand(`>ping test`)
	fmt.Println(strings.Join(args, ", "))

	args = parseCommand(`cmd --flag1 something --flag2 "test nothing"`)

	fmt.Println(strings.Join(args, ", "))


	// Output:
	//>ping, test
	//cmd, --flag1, something, --flag2, test nothing
}

func BenchmarkParseCommand(b *testing.B) {
	commandText := `cmd --flag1 something --flag2 "test nothing"`
	for n:= 0; n < b.N; n++ {
		parseCommand(commandText)
	}
}