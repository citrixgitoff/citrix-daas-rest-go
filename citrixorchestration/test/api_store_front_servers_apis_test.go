/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing StoreFrontServersAPIsDAASService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package citrixorchestration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/citrixorchestration"
)

func Test_citrixorchestration_StoreFrontServersAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StoreFrontServersAPIsDAASService StoreFrontServersCreateStoreFrontServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersCreateStoreFrontServer(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreFrontServersAPIsDAASService StoreFrontServersDeleteStoreFrontServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersDeleteStoreFrontServer(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreFrontServersAPIsDAASService StoreFrontServersGetStoreFrontDeliveryGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontDeliveryGroups(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreFrontServersAPIsDAASService StoreFrontServersGetStoreFrontServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServer(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreFrontServersAPIsDAASService StoreFrontServersGetStoreFrontServers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersGetStoreFrontServers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test StoreFrontServersAPIsDAASService StoreFrontServersUpdateStoreFrontServer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.StoreFrontServersAPIsDAAS.StoreFrontServersUpdateStoreFrontServer(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
