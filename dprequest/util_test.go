package dprequest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlJoin(t *testing.T) {
	uri1 := "http://www.example.com/"
	uri2 := "http://www.example.com"
	uri3 := "http://www.example.com/abc"
	uri4 := "abc/"
	uri5 := "abc"
	uri6 := "/def"
	uri7 := "http://www.example12.com/abc/def"
	assert.Equal(t, "http://www.example12.com/abc/def", BuildUrl(uri1, uri7).String())
	assert.Equal(t, "http://www.example.com", BuildUrl(uri1, uri2).String())
	assert.Equal(t, "http://www.example.com/abc", BuildUrl(uri1, uri3).String())
	assert.Equal(t, "http://www.example.com/abc/", BuildUrl(uri1, uri4).String())
	assert.Equal(t, "http://www.example.com/abc", BuildUrl(uri1, uri5).String())
	assert.Equal(t, "http://www.example.com/abc/abc/", BuildUrl(uri3, uri4).String())
	assert.Equal(t, "http://www.example.com/abc/abc", BuildUrl(uri3, uri5).String())
	assert.Equal(t, "http://www.example.com/abc/def", BuildUrl(uri3, uri6).String())
}
