package zoneforward

import (
	"github.com/sky-uk/skyinfoblox/api/common"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func testZoneForwardStruct(t *testing.T) {
	rec := new(ZoneForward)
	flag := false

	// just to test each field is settable...
	rec.Ref = "ref"
	rec.Address = "address"
	rec.Comment = "comment"
	rec.Disable = flag
	rec.DisplayDomain = "DisplayDomain"
	rec.DNSFqdn = "DNSFqdn"
	rec.ForwardTo = []common.ExternalServer{{Name: "Name"}}
	rec.ForwardersOnly = flag
	rec.ForwardingServers = []common.ForwardingMemberServer{{Name: "Name"}}
	rec.Fqdn = "Fqdn"
	rec.Locked = flag
	rec.LockedBy = "LockedBy"
	rec.MaskPrefix = "MaskPrefix"
	rec.Parent = "Parent"
	rec.Prefix = "Prefix"
	rec.UsingSrgAssociations = flag
	rec.View = "View"
	rec.ZoneFormat = "FORWARD"

}

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
	assert.Contains(t, api.Endpoint(), "_return_fields=Name")
}

func TestGetNoFields(t *testing.T) {
	api := NewGet("zone_forward/ZG5zLmhvc3QkLZhd3QuaDE", nil)
	assert.Equal(t, http.MethodGet, api.Method())
	assert.NotContains(t, api.Endpoint(), "_return_fields")
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
