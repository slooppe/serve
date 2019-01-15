package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVersionCommand(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	var b bytes.Buffer
	err := VersionCommand(&b)

	assert.NoError(err)
	assert.Contains(b.String(), fmt.Sprintf("version %s", version))
}

func TestServerCommand(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	var b bytes.Buffer
	log := log.New(&b, "[test] ", 0)
	opt := flags{Port: 0}

	go func() {
		assert.NoError(ServerCommand(log, opt))
	}()

	time.Sleep(200 * time.Millisecond)
}

func TestNewRouter(t *testing.T) {
	t.Parallel()

	var b bytes.Buffer
	log := log.New(&b, "[test] ", 0)
	opt := flags{Port: 0}

	r := NewRouter(log, opt)

	var _ http.Handler = r
}
