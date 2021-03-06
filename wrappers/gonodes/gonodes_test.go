package gonodes_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/alkemics/goflow"
	"github.com/alkemics/goflow/gfutil/gfgo"
	"github.com/alkemics/goflow/wrappers/bind"
	"github.com/alkemics/goflow/wrappers/gonodes"
	"github.com/alkemics/goflow/wrappers/imports"
	"github.com/alkemics/goflow/wrappers/types"
	"github.com/alkemics/goflow/wrappers/varnames"
)

func TestGoNodes(t *testing.T) {
	require := require.New(t)

	wd, err := os.Getwd()
	require.NoError(err)
	require.NoError(os.Chdir("../.."))

	var loader gfgo.NodeLoader
	err = loader.Load("github.com/alkemics/goflow/example/nodes")
	require.NoError(err)

	wraps := []goflow.GraphWrapper{
		gonodes.Wrapper(&loader),
		gonodes.DepWrapper,
		bind.Wrapper,
		varnames.Wrapper,
		types.Wrapper,
		imports.Merger,
		varnames.CompilableWrapper,
	}

	require.NoError(os.Chdir(wd))

	testCases := gfgo.TestCases{
		Tests: []gfgo.TestCase{
			{Deps: "true", Test: "g.Run()"},
			{Deps: "false", Test: "g.Run()"},
		},
	}
	gfgo.TestGenerate(t, wraps, "gonodes.yml", testCases)
}
