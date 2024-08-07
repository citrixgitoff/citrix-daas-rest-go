/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing SitesAPIsDAASService

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

func Test_citrixorchestration_SitesAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SitesAPIsDAASService SitesCheckObjectNameExists", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesCheckObjectNameExists(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetMultipleRemotePCAssignments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetMultipleRemotePCAssignments(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSessionsTrend", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSessionsTrend(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSite(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSiteErrorWarning", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSiteErrorWarning(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSiteMisconfigurationReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSiteMisconfigurationReport(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSiteSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSiteSettings(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSiteStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSiteStatus(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSiteTestReport", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSiteTestReport(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetSites", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetSites(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesGetUpgradePackageVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesGetUpgradePackageVersions(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesPatchMultipleRemotePCAssignments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.SitesAPIsDAAS.SitesPatchMultipleRemotePCAssignments(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesPatchSiteSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		httpRes, err := apiClient.SitesAPIsDAAS.SitesPatchSiteSettings(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SitesAPIsDAASService SitesTestSite", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var nameOrId string

		resp, httpRes, err := apiClient.SitesAPIsDAAS.SitesTestSite(context.Background(), nameOrId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
