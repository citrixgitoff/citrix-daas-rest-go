/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing DeliveryGroupsAPIsDAASService

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

func Test_citrixorchestration_DeliveryGroupsAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsAddDeliveryGroupMachineCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupMachineCatalog(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsAddDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var tagNameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsAddDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsCheckDeliveryGroupExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCheckDeliveryGroupExists(context.Background(), name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsCheckRebootScheduleNameExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var name string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCheckRebootScheduleNameExists(context.Background(), nameOrId, name).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsCreateDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroup(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsCreateDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupPowerTimeScheme(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsCreateDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsCreateDeliveryGroupRebootSchedule(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDeleteDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var policyId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var schemeNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDeleteDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var scheduleNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDeleteDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDoAddApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddApplications(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDoAddMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoAddMachines(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDoDeliveryGroupSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoDeliveryGroupSearch(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDoRemoveApplicationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var appGroupNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveApplicationGroups(context.Background(), nameOrId, appGroupNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDoRemoveApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var appNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveApplications(context.Background(), nameOrId, appNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsDoRemoveMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var machineNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsDoRemoveMachines(context.Background(), nameOrId, machineNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeiliveryGroupProjectedAutoscaleMachines(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicies(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var policyId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupApplications(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupMachines(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var schemeNameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupPowerTimeSchemes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupPowerTimeSchemes(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var scheduleNameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupRebootSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupRebootSchedules(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupStartMenuApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupStartMenuApplications(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTags(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupTestReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupTestReport(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroups(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupsAdministrators", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsAdministrators(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupsApplicationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsApplicationGroups(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupsDesktops", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsDesktops(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var id string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogDetails(context.Background(), nameOrId, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupsMachineCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsMachineCatalogs(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsGetDeliveryGroupsUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsGetDeliveryGroupsUsage(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsPatchDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var policyId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsPatchDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var schemeNameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsPatchDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var scheduleNameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsPatchDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsRemoveDeliveryGroupMachineCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var catalogNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsRemoveDeliveryGroupMachineCatalog(context.Background(), nameOrId, catalogNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsRemoveDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var tagNameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsRemoveDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsSetDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsSetDeliveryGroupTags(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsStartDeliveryGroupRebootCycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsStartDeliveryGroupRebootCycle(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsTestDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroup(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsAPIsDAASService DeliveryGroupsTestDeliveryGroupExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.DeliveryGroupsAPIsDAAS.DeliveryGroupsTestDeliveryGroupExists(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
