/*
Global App Config Admin

Testing DiscoveryControllerDAASService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package globalappconfiguration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/globalappconfiguration"
)

func Test_globalappconfiguration_DiscoveryControllerDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DiscoveryControllerDAASService DeleteDiscoveryApiUsingDELETE", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var domain string

		httpRes, err := apiClient.DiscoveryControllerDAAS.DeleteDiscoveryApiUsingDELETE(context.Background(), app, domain).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryControllerDAASService GetAllDiscoveryApiUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DiscoveryControllerDAAS.GetAllDiscoveryApiUsingGET(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryControllerDAASService GetDiscoveryApiUsingGET", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var domain string

		resp, httpRes, err := apiClient.DiscoveryControllerDAAS.GetDiscoveryApiUsingGET(context.Background(), app, domain).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryControllerDAASService PostDiscoveryApiUsingPOST", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string

		resp, httpRes, err := apiClient.DiscoveryControllerDAAS.PostDiscoveryApiUsingPOST(context.Background(), app).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiscoveryControllerDAASService PutDiscoveryApiUsingPUT", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var domain string

		httpRes, err := apiClient.DiscoveryControllerDAAS.PutDiscoveryApiUsingPUT(context.Background(), app, domain).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
