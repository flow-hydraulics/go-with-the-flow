package main

import (
	"testing"

	"github.com/bjartek/go-with-the-flow/v2/gwtf"
	"github.com/stretchr/testify/assert"
)

func TestScript(t *testing.T) {
	g := gwtf.NewTestingEmulator()
	t.Parallel()

	t.Run("Raw account argument", func(t *testing.T) {
		value := g.ScriptFromFile("test", []byte{}).RawAccountArgument("0x1cf0e2f2f715450").RunReturnsInterface()
		assert.Equal(t, value, "0x1cf0e2f2f715450")
	})

	t.Run("Raw account argument", func(t *testing.T) {
		value := g.ScriptFromFile("test", []byte{}).AccountArgument("first").RunReturnsInterface()
		assert.Equal(t, value, "0x1cf0e2f2f715450")
	})

	t.Run("Script should report failure", func(t *testing.T) {
		value, err := g.Script("asdf").RunReturns()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "Parsing failed")
		assert.Nil(t, value)

	})

}
