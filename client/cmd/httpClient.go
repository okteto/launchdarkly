/*
Copyright © 2023 Okteto Inc.
*/
package cmd

import retryablehttp "github.com/hashicorp/go-retryablehttp"

func getRetryableClient() *retryablehttp.Client {
	client := retryablehttp.NewClient()
	client.Logger = nil
	return client
}
