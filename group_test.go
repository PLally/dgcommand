package dgcommand

import (
	"github.com/plally/dgcommand/embed"
	"io"
	"reflect"
	"testing"
)

func TestGroup(t *testing.T) {
	commandTests := []struct{
		command string
		replys []string
	}{
		{"!ping", []string{"PONG"} },
		{"ping", []string{} },
		{"!mygroup hello", []string{"world"} },
		{"!hello world", []string{"world"} },
	}

	g := Group()

	g.On("ping", func(ctx Context) {
		ctx.Reply("PONG")
	})

	subGroup := Group()

	subGroup.On("hello", func(ctx Context){
		ctx.Reply("world")
	})

	g.Command("hello <arg1>", func(ctx Context){
		ctx.Reply(ctx.Args()[0])
	})

	g.AddHandler("mygroup", subGroup)

	prefixed := OnPrefix("!", g)

	for _, testData := range commandTests {
		t.Run(testData.command, func(t *testing.T) {
			ctx := mockContext{}
			ctx.SetArgs([]string{
				testData.command,
			})

			prefixed.Handle(&ctx)

			success := reflect.DeepEqual(ctx.replys, testData.replys) || len(ctx.replys) == 0 && len(testData.replys) == 0

			if !success {
				t.Errorf("Exepected %v got %v", testData.replys, ctx.replys)
			}

		})
	}
}

type mockContext struct {
	ArgContainer
	replys []string
}

func (m *mockContext) SendFile(name string, r io.Reader) {
	panic("implement me")
}

func (m *mockContext) SendEmbed(embed.Embed) {
	panic("implement me")
}

func (m *mockContext) Reply(s string) {
	m.replys = append(m.replys, s)
}

func (m *mockContext) Error(interface{}) {}