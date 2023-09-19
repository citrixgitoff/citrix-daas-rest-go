/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing AdminAPIs Service

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

func Test_citrixorchestration_AdminAPIs Service(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AdminAPIs Service AdminCheckRoleNameExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.AdminAPIs .AdminCheckRoleNameExists(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminCheckScopeNameExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.AdminAPIs .AdminCheckScopeNameExists(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminCreateAdminRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdminAPIs .AdminCreateAdminRole(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminCreateAdminScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.AdminAPIs .AdminCreateAdminScope(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminDeleteAdminRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.AdminAPIs .AdminDeleteAdminRole(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminDeleteAdminScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.AdminAPIs .AdminDeleteAdminScope(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminAdministrator", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminAdministrator(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminAdministrators", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminAdministrators(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminEffectiveRights", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminEffectiveRights(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminRole(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminRoles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminScope(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminScopedObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminScopedObjects(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetAdminScopes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetAdminScopes(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetPredefinedPermissionGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetPredefinedPermissionGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetPredefinedPermissions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetPredefinedPermissions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminGetPredefinedPermissionsForGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.AdminAPIs .AdminGetPredefinedPermissionsForGroups(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminUpdateAdminRole", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.AdminAPIs .AdminUpdateAdminRole(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AdminAPIs Service AdminUpdateAdminScope", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.AdminAPIs .AdminUpdateAdminScope(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
