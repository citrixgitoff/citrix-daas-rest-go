/*
Quick Create Service for DaaS Quick Deploy with AWS WorkSpaces Core - Tech Preview

Testing ImageQCSService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package citrixquickcreate

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixquickcreate"
)

func Test_citrixquickcreate_ImageQCSService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImageQCSService CopyImageAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var accountId string
		var imageId string

		resp, httpRes, err := apiClient.ImageQCS.CopyImageAsync(context.Background(), customerId, accountId, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageQCSService GetImageAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var accountId string
		var imageId string

		resp, httpRes, err := apiClient.ImageQCS.GetImageAsync(context.Background(), customerId, accountId, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageQCSService GetImagesAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var accountId string

		resp, httpRes, err := apiClient.ImageQCS.GetImagesAsync(context.Background(), customerId, accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageQCSService ImportImageAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var accountId string

		resp, httpRes, err := apiClient.ImageQCS.ImportImageAsync(context.Background(), customerId, accountId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImageQCSService RemoveImageAsync", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerId string
		var accountId string
		var imageId string

		httpRes, err := apiClient.ImageQCS.RemoveImageAsync(context.Background(), customerId, accountId, imageId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
