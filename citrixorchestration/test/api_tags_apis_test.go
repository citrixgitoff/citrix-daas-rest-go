/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing TagsAPIsDAASService

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

func Test_citrixorchestration_TagsAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIsDAASService TagsCheckTagExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.TagsAPIsDAAS.TagsCheckTagExists(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsCreateTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsCreateTag(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsDeleteTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.TagsAPIsDAAS.TagsDeleteTag(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTag(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTagApplicationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTagApplicationGroups(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTagApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTagApplications(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTagDeliveryGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTagDeliveryGroups(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTagMachineCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTagMachineCatalogs(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTagMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTagMachines(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsGetTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsGetTags(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsPatchTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.TagsAPIsDAAS.TagsPatchTag(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsSetTagApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.TagsAPIsDAAS.TagsSetTagApplications(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIsDAASService TagsSetTagDeliveryGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.TagsAPIsDAAS.TagsSetTagDeliveryGroups(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
