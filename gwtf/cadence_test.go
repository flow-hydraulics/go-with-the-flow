package gwtf

import (
	"testing"

	"github.com/onflow/cadence"
	"github.com/stretchr/testify/assert"
)

func TestCadenceValueToJsonString(t *testing.T) {

	t.Parallel()
	t.Run("Empty value should be empty json object", func(t *testing.T) {
		value := CadenceValueToJsonString(nil)
		assert.Equal(t, "{}", value)
	})

	t.Run("Empty optional should be empty string", func(t *testing.T) {
		value := CadenceValueToJsonString(cadence.NewOptional(nil))
		assert.Equal(t, `""`, value)
	})
	t.Run("Unwrap optional", func(t *testing.T) {
		fooTest, _ := cadence.NewString("foo")
		value := CadenceValueToJsonString(cadence.NewOptional(fooTest))
		assert.Equal(t, `"foo"`, value)
	})
	t.Run("Array", func(t *testing.T) {
		fooTest, _ := cadence.NewString("foo")
		barTest, _ := cadence.NewString("bar")
		value := CadenceValueToJsonString(cadence.NewArray([]cadence.Value{fooTest, barTest}))
		assert.Equal(t, `[
    "foo",
    "bar"
]`, value)
	})

	t.Run("Dictionary", func(t *testing.T) {
		fooTest, _ := cadence.NewString("foo")
		barTest, _ := cadence.NewString("bar")
		dict := cadence.NewDictionary([]cadence.KeyValuePair{{Key: fooTest, Value: barTest}})
		value := CadenceValueToJsonString(dict)
		assert.Equal(t, `{
    "foo": "bar"
}`, value)
	})

	t.Run("Struct", func(t *testing.T) {
		barTest, _ := cadence.NewString("bar")
		s := cadence.Struct{
			Fields: []cadence.Value{barTest},
			StructType: &cadence.StructType{
				Fields: []cadence.Field{{
					Identifier: "foo",
					Type:       cadence.StringType{},
				}},
			},
		}
		value := CadenceValueToJsonString(s)
		assert.Equal(t, `{
    "foo": "bar"
}`, value)
	})

}
