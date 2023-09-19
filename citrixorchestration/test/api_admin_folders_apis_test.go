/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing AdminFoldersAPIs Service

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

func Test_citrixorchestration_AdminFoldersAPIs Service(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdminFoldersAPIs Service AdminFoldersCreateAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersCreateAdminFolder(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersAPIs Service AdminFoldersDeleteAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string

		httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersDeleteAdminFolder(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersAPIs Service AdminFoldersGetAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersGetAdminFolder(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersAPIs Service AdminFoldersGetAdminFolderApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersGetAdminFolderApplications(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersAPIs Service AdminFoldersGetAdminFolderMachineCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersGetAdminFolderMachineCatalogs(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersAPIs Service AdminFoldersGetAdminFolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersGetAdminFolders(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersAPIs Service AdminFoldersUpdateAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string

		resp, httpRes, err := apiClient.AdminFoldersAPIs .AdminFoldersUpdateAdminFolder(context.Background(), pathOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
