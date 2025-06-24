package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type someData struct {
	Value string
}

func TestGet_TTLValid(t *testing.T) {
	expected := &someData{Value: "foo"}

	c := &Cache[someData]{
		data:       expected,
		ttl:        1 * time.Hour,
		lastParsed: time.Now(),
	}

	data := c.Get()

	assert.Equal(t, data, expected)
}

func TestGet_TTLExpired(t *testing.T) {
	expected := &someData{Value: "foo"}

	c := &Cache[someData]{
		data:       expected,
		ttl:        time.Hour,
		lastParsed: time.Now().Add(-2 * time.Hour),
	}

	data := c.Get()

	assert.Nil(t, data)
}

func TestGet_LastParsedNotSet(t *testing.T) {
	expected := &someData{Value: "foo"}

	c := &Cache[someData]{
		data: expected,
		ttl:  time.Hour,
	}

	data := c.Get()

	assert.Nil(t, data)
}

func TestGet_EmptyData(t *testing.T) {
	expected := &someData{Value: ""}

	c := &Cache[someData]{
		data:       expected,
		ttl:        1 * time.Hour,
		lastParsed: time.Now(),
	}

	data := c.Get()

	assert.NotNil(t, data)
	assert.Equal(t, data, expected)
}
