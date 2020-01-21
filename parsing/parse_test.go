package parsing

import (
	"reflect"
	"testing"
)


func TestParseDefinition(t *testing.T) {
	definitionStrings := []struct {
		def string
		in string
		result []string
	} {
		{"test <arg1> <arg2> [vararg...]", "arg1 arg2 arg3 continued",
			[]string{"arg1", "arg2", "arg3 continued", },
		},
		{"test <arg1> <arg2> [vararg...]", "arg1 arg2",
			[]string{"arg1", "arg2", "" },
		},


	}
	for _ , d := range definitionStrings {
		c := NewCommandDefinition(d.def	)
		parsedResult, err := c.ParseInput(d.in)
		if err  != nil  {
			t.Errorf("%v errored with: %v", d.def, err.Error())
			return
		}
		if !reflect.DeepEqual(parsedResult, d.result) {
			t.Errorf("testing command def %v failed", d.def)
			return
		}

	}

}