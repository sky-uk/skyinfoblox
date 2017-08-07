package zoneforward

import (
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func setupRec() ZoneForward {
	forwardToList := make([]common.ExternalServer, 1)
	forwardToList[0] = common.ExternalServer{
		Name: "foo",
	}
	return ZoneForward{
		Fqdn:      "foo.bar.com",
		ForwardTo: forwardToList,
	}
}

func TestCreate(t *testing.T) {
	rec := setupRec()
	api := NewCreate(rec)
	assert.Equal(t, http.MethodPost, api.Method())
}

func TestGetAll(t *testing.T) {
	api := NewGetAll()
	assert.Equal(t, http.MethodGet, api.Method())
}

func TestGet(t *testing.T) {
	api := NewGet("zone_forward/ZG5zLmhvc3QkLZhd3QuaDE", []string{"Name"})
	assert.Equal(t, http.MethodGet, api.Method())
}

func TestUpdate(t *testing.T) {
	rec := setupRec()
	api := NewUpdate(rec, []string{"Name"})
	assert.Equal(t, http.MethodPut, api.Method())
}

func TestDelete(t *testing.T) {
	api := NewDelete("ref")
	assert.Equal(t, http.MethodDelete, api.Method())
}
