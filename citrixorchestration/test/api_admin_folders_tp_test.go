/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

Testing AdminFoldersTPApiService

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

func Test_citrixorchestration_AdminFoldersTPApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPCreateAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPCreateAdminFolder(context.Background(), customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPDeleteAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPDeleteAdminFolder(context.Background(), pathOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPGetAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolder(context.Background(), pathOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPGetAdminFolderApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolderApplications(context.Background(), pathOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPGetAdminFolderMachineCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolderMachineCatalogs(context.Background(), pathOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPGetAdminFolders", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPGetAdminFolders(context.Background(), customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminFoldersTPApiService AdminFoldersTPUpdateAdminFolder", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var pathOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.AdminFoldersTPApi.AdminFoldersTPUpdateAdminFolder(context.Background(), pathOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
