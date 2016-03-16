package monitors

import (
	"strconv"

	"github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"
)

const (
	path        = "loadbalancers"
	monitorPath = "healthmonitor"
)

func rootURL(c *gophercloud.ServiceClient, lbID int) string {
	return c.ServiceURL(path, strconv.Itoa(lbID), monitorPath)
}
