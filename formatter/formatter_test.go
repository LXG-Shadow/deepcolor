package formatter

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewJsonSchema(t *testing.T) {
	scheme := map[string]interface{}{
		"asdf": "json::a.b",
		"values": map[string]interface{}{
			"a": "json::g.x",
			"value": []interface{}{
				map[string]interface{}{
					"a": "json::g.c.#.x",
				},
			},
			"value2": []interface{}{"json::g.c.#.x"},
			"value3": []interface{}{"regex::f."},
		},
		"data":  "regex::world.",
		"data2": "regex::f.",
	}
	testdata, _ := json.Marshal(map[string]interface{}{
		"a": map[string]interface{}{
			"b": "c",
		},
		"g": map[string]interface{}{
			"x": "value1",
			"c": []interface{}{
				map[string]int{
					"x": 1,
				},
				map[string]int{
					"x": 2,
				},
			},
		},
		"hello":     "world1",
		"otherdata": "f1f2f3f4",
	})
	schema := NewJsonSchema(scheme)
	require.NotNil(t, schema)
	value := schema.Format(string(testdata))
	v, err := json.Marshal(value)
	require.Equal(t, `{"asdf":"c","data":"world1","data2":"f1","values":{"a":"value1","value":[{"a":1}],"value2":[1,2],"value3":["f1","f2","f3","f4"]}}`, string(v))
	v, err = json.MarshalIndent(value, "", "  ")
	require.NoError(t, err)
	t.Log(string(v))
}
