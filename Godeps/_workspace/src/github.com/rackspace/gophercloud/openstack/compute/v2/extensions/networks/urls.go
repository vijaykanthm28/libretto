package networks

import "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"

const resourcePath = "os-networks"

func resourceURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *gophercloud.ServiceClient) string {
	return resourceURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}
