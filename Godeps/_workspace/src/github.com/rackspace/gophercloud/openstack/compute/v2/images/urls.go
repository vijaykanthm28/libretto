package images

import "github.com/apcera/libretto/Godeps/_workspace/src/github.com/rackspace/gophercloud"

func listDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images", "detail")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}
