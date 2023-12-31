/*
OpenBuckets API

Testing FilesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openbuckets

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "https://github.com/openbuckets/sdk-go"
)

func Test_openbuckets_FilesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FilesAPIService SearchFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FilesAPI.SearchFiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
