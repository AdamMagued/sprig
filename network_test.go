package sprig

import (
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHostByName(t *testing.T) {
	tpl := `{{"www.google.com" | getHostByName}}`

	resolvedIP, _ := runRaw(tpl, nil)

	ip := net.ParseIP(resolvedIP)
	assert.NotNil(t, ip)
	assert.NotEmpty(t, ip)
}

func TestJoinHostPort(t *testing.T) {
	assert.Equal(t, "127.0.0.1:8080", joinHostPort("127.0.0.1", 8080))
	assert.Equal(t, "127.0.0.1:8080", joinHostPort("127.0.0.1", "8080"))
	assert.Equal(t, "[::1]:443", joinHostPort("::1", 443))
	assert.Equal(t, "[::1]:443", joinHostPort("::1", "443"))
	assert.Equal(t, "localhost:http", joinHostPort("localhost", "http"))

	assert.NoError(t, runt(`{{ joinHostPort "127.0.0.1" 8080 }}`, "127.0.0.1:8080"))
	assert.NoError(t, runt(`{{ joinHostPort "::1" 443 }}`, "[::1]:443"))
	assert.NoError(t, runt(`{{ joinHostPort "localhost" "http" }}`, "localhost:http"))
	assert.NoError(t, runt(`{{ 8080 | joinHostPort "127.0.0.1" }}`, "127.0.0.1:8080"))
	assert.NoError(t, runt(`{{ 443 | joinHostPort "::1" }}`, "[::1]:443"))
}

