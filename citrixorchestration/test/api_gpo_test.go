/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing GpoDAASService

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

func Test_citrixorchestration_GpoDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GpoDAASService GpoComparePolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoComparePolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoCopyGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GpoDAAS.GpoCopyGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoCopyGpoPolicySet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policySetGuid string

		resp, httpRes, err := apiClient.GpoDAAS.GpoCopyGpoPolicySet(context.Background(), policySetGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoCreateGpoFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoCreateGpoFilter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoCreateGpoPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoCreateGpoPolicy(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoCreateGpoPolicySet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoCreateGpoPolicySet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoCreateGpoSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoCreateGpoSetting(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoDeleteGpoFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filterGuid string

		httpRes, err := apiClient.GpoDAAS.GpoDeleteGpoFilter(context.Background(), filterGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoDeleteGpoPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyGuid string

		httpRes, err := apiClient.GpoDAAS.GpoDeleteGpoPolicy(context.Background(), policyGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoDeleteGpoPolicySet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policySetGuid string

		httpRes, err := apiClient.GpoDAAS.GpoDeleteGpoPolicySet(context.Background(), policySetGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoDeleteGpoSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingGuid string

		httpRes, err := apiClient.GpoDAAS.GpoDeleteGpoSetting(context.Background(), settingGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoDisableGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GpoDAAS.GpoDisableGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoEnableGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GpoDAAS.GpoEnableGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoGetFilterDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoGetFilterDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoGetSettingDefinitions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoGetSettingDefinitions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoGetSettingFullDetail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoGetSettingFullDetail(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoMoveGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GpoDAAS.GpoMoveGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoRankGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoRankGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filterGuid string

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoFilter(context.Background(), filterGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoFilters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoFilters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyGuid string

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoPolicy(context.Background(), policyGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoPolicySet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policySetGuid string

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoPolicySet(context.Background(), policySetGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoPolicySets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoPolicySets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingGuid string

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoSetting(context.Background(), settingGuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoReadGpoSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoReadGpoSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoRemoveGpoPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GpoDAAS.GpoRemoveGpoPolicies(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoRunSimulation", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoRunSimulation(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoSearchFilters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoSearchFilters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoSearchPolicies", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoSearchPolicies(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoSearchPolicySets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoSearchPolicySets(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoSearchSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GpoDAAS.GpoSearchSettings(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoUpdateGpoFilter", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var filterGuid string

		httpRes, err := apiClient.GpoDAAS.GpoUpdateGpoFilter(context.Background(), filterGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoUpdateGpoPolicy", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyGuid string

		httpRes, err := apiClient.GpoDAAS.GpoUpdateGpoPolicy(context.Background(), policyGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoUpdateGpoPolicySet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policySetGuid string

		httpRes, err := apiClient.GpoDAAS.GpoUpdateGpoPolicySet(context.Background(), policySetGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoUpdateGpoSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var settingGuid string

		httpRes, err := apiClient.GpoDAAS.GpoUpdateGpoSetting(context.Background(), settingGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GpoDAASService GpoUpdatePolicySetBlob", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policySetGuid string

		httpRes, err := apiClient.GpoDAAS.GpoUpdatePolicySetBlob(context.Background(), policySetGuid).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
