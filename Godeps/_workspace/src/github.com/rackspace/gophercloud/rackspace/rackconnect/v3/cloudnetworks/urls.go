package cloudnetworks

import "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"

func listURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL("cloud_networks")
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL("cloud_networks", id)
}
