/*
Citrix Virtual Apps and Desktops REST API TECHPREVIEW

Testing DeliveryGroupsTPApiService

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

func Test_citrixorchestration_DeliveryGroupsTPApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPAddDeliveryGroupMachineCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupMachineCatalog(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPAddDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var tagNameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPAddDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPCheckDeliveryGroupExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var name string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCheckDeliveryGroupExists(context.Background(), name, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPCheckRebootScheduleNameExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var name string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCheckRebootScheduleNameExists(context.Background(), nameOrId, name, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPCreateDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroup(context.Background(), customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPCreateDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPCreateDeliveryGroupRebootSchedule(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDeleteDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var policyId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var schemeNameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var scheduleNameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDeleteDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDoAddApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoAddApplications(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDoAddMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoAddMachines(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDoDeliveryGroupSearch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoDeliveryGroupSearch(context.Background(), customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDoRemoveApplicationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var appGroupNameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveApplicationGroups(context.Background(), nameOrId, appGroupNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDoRemoveApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var appNameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveApplications(context.Background(), nameOrId, appNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPDoRemoveMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var machineNameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPDoRemoveMachines(context.Background(), nameOrId, machineNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeiliveryGroupProjectedAutoscaleMachines(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicies(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var policyId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupApplications(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupMachines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupMachines(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var schemeNameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupPowerTimeSchemes(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var scheduleNameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupRebootSchedules", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupRebootSchedules(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupStartMenuApplications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupStartMenuApplications(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTags(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupTestReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupTestReport(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroups(context.Background(), customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupsAdministrators", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsAdministrators(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupsApplicationGroups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsApplicationGroups(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupsDesktops", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsDesktops(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var id string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogDetails(context.Background(), nameOrId, id, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsMachineCatalogs(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPGetDeliveryGroupsUsage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPGetDeliveryGroupsUsage(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPPatchDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var policyId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupAdvancedAccessPolicy(context.Background(), nameOrId, policyId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var schemeNameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupPowerTimeScheme(context.Background(), nameOrId, schemeNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPPatchDeliveryGroupRebootSchedule", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var scheduleNameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPPatchDeliveryGroupRebootSchedule(context.Background(), nameOrId, scheduleNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var catalogNameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupMachineCatalog(context.Background(), nameOrId, catalogNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPRemoveDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var tagNameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPRemoveDeliveryGroupTags(context.Background(), nameOrId, tagNameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPSetDeliveryGroupTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPSetDeliveryGroupTags(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPStartDeliveryGroupRebootCycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPStartDeliveryGroupRebootCycle(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPTestDeliveryGroup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string
		var customerid string
		var siteid string

		resp, httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroup(context.Background(), nameOrId, customerid, siteid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DeliveryGroupsTPApiService DeliveryGroupsTPTestDeliveryGroupExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customerid string
		var siteid string

		httpRes, err := apiClient.DeliveryGroupsTPApi.DeliveryGroupsTPTestDeliveryGroupExists(context.Background(), customerid, siteid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
