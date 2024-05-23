/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing PvsStreamingAPIsDAASService

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

func Test_citrixorchestration_PvsStreamingAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PvsStreamingAPIsDAASService PvsStreamingGetPvsStreamingSites", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingSites(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PvsStreamingAPIsDAASService PvsStreamingGetPvsStreamingStores", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var farmId string

		resp, httpRes, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingStores(context.Background(), farmId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PvsStreamingAPIsDAASService PvsStreamingGetPvsStreamingVDisks", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingGetPvsStreamingVDisks(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PvsStreamingAPIsDAASService PvsStreamingTestDeviceCollectionExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.PvsStreamingAPIsDAAS.PvsStreamingTestDeviceCollectionExists(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
