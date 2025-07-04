/*
Zibal API

Testing GatewayAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package zibal

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mostafnamazy/zibal"
)

func Test_zibal_GatewayAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GatewayAPIService RequestPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GatewayAPI.RequestPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatewayAPIService StartTrackIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var trackId int64

		httpRes, err := apiClient.GatewayAPI.StartTrackIdGet(context.Background(), trackId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GatewayAPIService VerifyPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GatewayAPI.VerifyPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
