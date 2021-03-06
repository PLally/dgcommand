package parsing

import (
	"errors"
	"github.com/plally/dgcommand/parsing/util"
	"regexp"
	"strings"
)

// TODO write a more efficient definition parser that doesnt use regex
// TODO support --flags
var (
	requiredArg = regexp.MustCompile(`^<[a-zA-Z][a-zA-Z0-9]*>`)
	optionalArg = regexp.MustCompile(`^\[[a-zA-Z][a-zA-Z0-9]*\]`)
	varArg      = regexp.MustCompile(`^\[[a-zA-Z][a-zA-Z0-9]*\.\.\.]`)
	space       = regexp.MustCompile("^ +")
	word        = regexp.MustCompile("^[a-zA-Z0-9]+")
)

type token struct {
	tokenPattern *regexp.Regexp
	value        string
}

func lexCommandDefinition(s string) (tokens []token) {

	i := 0
	length := len(s)

MainLoop:
	for i < length {
		if v := space.FindString(s[i:]); v != "" {
			i += len(v)
			continue MainLoop
		}

		for _, pattern := range []*regexp.Regexp{requiredArg, optionalArg, varArg, word} {
			v := pattern.FindString(s[i:])
			if v != "" {
				tokens = append(tokens, token{pattern, v})
				i += len(v)
				continue MainLoop
			}
		}
		panic("Command Definition Lexer: Invalid command definition")

	}
	return tokens
}

type CommandArgument struct {
	IsOptional bool
	IsVarArg   bool
	Name       string
}

func (a *CommandArgument) String() string {
	var out string

	if a.IsOptional {
		out += "["
	} else {
		out += "<"
	}
	out += a.Name
	if a.IsVarArg {
		out += "..."
	}

	if a.IsOptional {
		out += "]"
	} else {
		out += ">"
	}
	return out
}

type CommandDefinition struct {
	Name   string
	tokens []token
	Args   []CommandArgument
}

func (cmd *CommandDefinition) String() string {
	out := cmd.Name
	for _, arg := range cmd.Args {
		out += " " + arg.String()
	}
	return out
}

func (cmd *CommandDefinition) ParseInput(s string) ([]string, error) {
	var args []string
	if len(cmd.Args) < 1 {
		return args, nil
	}
	util.ConsumeSpaces(&s)

	for _, arg := range cmd.Args {
		if arg.IsVarArg {
			args = append(args, s)
			continue
		}
		a := parseArg(&s)
		if a == "" {
			return args, errors.New("Missing required arg " + arg.Name)
		}
		args = append(args, a)

	}
	return args, nil
}

func parseArg(p *string) (arg string) {
	i := 0
	util.ConsumeSpaces(p)

	i = 0
	for i < len(*p) {
		c := (*p)[i]
		switch c {
		case '"':
			end := strings.Index((*p)[i+1:], `"`)

			if end == -1 {
				end = len(*p) - 1
			}
			arg = (*p)[i+1 : end+1]
			*p = (*p)[end+2:]
			return
		case ' ':
			arg = (*p)[:i]
			*p = (*p)[i+1:]
			return

		}

		i++
	}
	s := *p
	*p = ""
	return s
}

func parseCommandDefinitionTokens(tokens []token) CommandDefinition {
	c := CommandDefinition{
		Name:   "",
		tokens: tokens,
		Args:   make([]CommandArgument, 0),
	}
	for _, t := range tokens {
		switch t.tokenPattern {
		case word:
			if c.Name != "" {
				panic("Command Definition Parser: command names can not contain spaces ")
			}
			c.Name = t.value
		case requiredArg:
			if len(c.Args) > 0 && c.Args[len(c.Args)-1].IsVarArg {
				panic("Command Definition Parser: a Var Arg must be the last argument in a command")
			}
			if len(c.Args) > 0 && c.Args[len(c.Args)-1].IsOptional {
				panic("Command Definition Parser: an IsOptional must not be followed by a required arg")
			}

			c.Args = append(c.Args, CommandArgument{false, false, t.value[1 : len(t.value)-1]})
		case optionalArg:
			if len(c.Args) > 0 && c.Args[len(c.Args)-1].IsVarArg {
				panic("Command Definition Parser: a Var Arg must be the last argument in a command")
			}
			c.Args = append(c.Args, CommandArgument{true, false, t.value[1 : len(t.value)-1]})
		case varArg:
			c.Args = append(c.Args, CommandArgument{true, true, t.value[1 : len(t.value)-4]})
		}
	}
	return c
}

func NewCommandDefinition(s string) CommandDefinition {
	tokens := lexCommandDefinition(s)
	return parseCommandDefinitionTokens(tokens)
}
