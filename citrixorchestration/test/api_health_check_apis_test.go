/*
Citrix Virtual Apps and Desktops Rest API for Developer Portal

Testing HealthCheckAPIsDAASService

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

func Test_citrixorchestration_HealthCheckAPIsDAASService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test HealthCheckAPIsDAASService HealthCheckHealthCheck", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.HealthCheckAPIsDAAS.HealthCheckHealthCheck(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
