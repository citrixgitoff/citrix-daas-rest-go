/*
Citrix.CloudServices.Administrators.Api

Testing ServicePrincipalsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ccadmins

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/citrix/citrix-daas-rest-go/ccadmins"
)

func Test_ccadmins_ServicePrincipalsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsDelete(context.Background(), customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsGet(context.Background(), customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsIdAccessPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdAccessPut(context.Background(), id, customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdDelete(context.Background(), id, customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdGet(context.Background(), id, customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsIdRotatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string
		var id string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdRotatePost(context.Background(), customer, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsIdSecondarySecretDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string
		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdSecondarySecretDelete(context.Background(), id, customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsIdUpdateLastAccessedDatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string
		var id string

		httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsIdUpdateLastAccessedDatePost(context.Background(), customer, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsMigratePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string

		httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsMigratePost(context.Background(), customer).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServicePrincipalsAPIService CustomerServicePrincipalsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var customer string

		resp, httpRes, err := apiClient.ServicePrincipalsAPI.CustomerServicePrincipalsPost(context.Background(), customer).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
