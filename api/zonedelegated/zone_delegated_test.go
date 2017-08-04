package zonedelegated

import (
	"fmt"
	"github.com/sky-uk/skyinfoblox/api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"github.com/sky-uk/skyinfoblox/api/common"
)

func setupZoneDelegated() *api.BaseAPI {
	delegateObject := common.ExternalServer{
		Address:"172.16.0.1",
		Name: "dns1.testdomain.com",
	}
	disableZone := false
	newZone := ZoneDelegated{
		Address: "10.0.0.1",
		Comment: "This is a comment",
		DelegateTo: delegateObject,
		DelegatedTTL: 1234,
		Disable: &disableZone,
		DnsFqdn: "example.com",
	}


}