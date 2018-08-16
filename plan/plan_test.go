package plan

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parsePlan(t *testing.T) {
	in := `
* something
* hey
- test
- omg
+ old stuff now done
this is a note
again
* yeah
+ sure?
mixed
- nope
`
	p := parse(in)
	s := p.String()
	expected := `* something
* hey
* yeah

+ old stuff now done
+ sure?

- test
- omg
- nope

this is a note
again
mixed
`
	require.Equal(t, expected, s)
}
