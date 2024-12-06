package formatter

import (
	"strings"
)

type IValueGetter interface {
	Get(expression string) interface{}
}

type ValueGetter struct {
	Type       string
	Expression string
}

func ValueGetterFromStr(value string) ValueGetter {
	values := strings.Split(value, "::")
	if len(values) < 2 {
		return ValueGetter{}
	}
	return ValueGetter{
		Type:       values[0],
		Expression: strings.Join(values[1:], "::"),
	}
}

type valueCache struct {
	data  string
	cache map[string]IValueGetter
}

func (f *valueCache) getValue(info ValueGetter) interface{} {
	v, ok := f.cache[info.Type]
	if !ok {
		v = registry[info.Type](f.data)
	}
	return v.Get(info.Expression)
}

func newValueCache(data string) *valueCache {
	return &valueCache{
		data:  data,
		cache: map[string]IValueGetter{},
	}
}

type JsonSchema struct {
	scheme map[string]interface{}
}

type JsonArraySchema struct {
	schema interface{}
}

func (s *JsonSchema) format(cache *valueCache) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range s.scheme {
		switch value.(type) {
		case ValueGetter:
			data := cache.getValue(value.(ValueGetter))
			if darray, ok := data.([]interface{}); ok {
				result[key] = darray[0]
			} else {
				result[key] = data
			}
		case *JsonSchema:
			result[key] = value.(*JsonSchema).format(cache)
		case *JsonArraySchema:
			result[key] = value.(*JsonArraySchema).format(cache)
		}
	}
	return result
}

func (s *JsonSchema) Format(value string) map[string]interface{} {
	return s.format(newValueCache(value))
}

func (s *JsonArraySchema) format(cache *valueCache) []interface{} {
	switch schema := s.schema.(type) {
	case ValueGetter:
		val := cache.getValue(schema)
		if arr, ok := val.([]interface{}); ok {
			// If it's already an array, return it as is
			return arr
		}
		// If single value, wrap in an array
		return []interface{}{val}

	case *JsonSchema:
		// If we have a JsonSchema, it might return a single map with arrays inside.
		singleResult := schema.format(cache)
		// Check if any field in singleResult is an array. If so, replicate the object for each element.
		for k, v := range singleResult {
			if arr, ok := v.([]interface{}); ok {
				// We have found an array field. We will replicate the schema for each element.
				var results []interface{}
				for _, elem := range arr {
					newMap := make(map[string]interface{}, len(singleResult))
					for kk, vv := range singleResult {
						newMap[kk] = vv
					}
					// Replace the array field with the single element
					newMap[k] = elem
					results = append(results, newMap)
				}
				return results
			}
		}
		// No arrays found, just return a single-element array
		return []interface{}{singleResult}
	}
	return nil
}
func (s *JsonArraySchema) Format(value string) []interface{} {
	return s.format(newValueCache(value))
}

func NewJsonArraySchema(schema []interface{}) *JsonArraySchema {
	if len(schema) != 1 {
		// invalid schema
		return nil
	}
	switch schema[0].(type) {
	case string:
		return &JsonArraySchema{
			schema: ValueGetterFromStr(schema[0].(string)),
		}
	case map[string]interface{}:
		return &JsonArraySchema{
			schema: NewJsonSchema(schema[0].(map[string]interface{})),
		}
	}
	return nil
}

func NewJsonSchema(schema map[string]interface{}) *JsonSchema {
	result := make(map[string]interface{})
	for key, value := range schema {
		switch value.(type) {
		case string:
			result[key] = ValueGetterFromStr(value.(string))
		case map[string]interface{}:
			v := NewJsonSchema(value.(map[string]interface{}))
			if v == nil {
				return nil
			}
			result[key] = v
		case []interface{}:
			v := NewJsonArraySchema(value.([]interface{}))
			if v == nil {
				return nil
			}
			result[key] = v
		}
	}
	return &JsonSchema{
		scheme: result,
	}
}
