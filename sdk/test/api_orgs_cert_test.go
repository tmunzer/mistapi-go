/*
Mist API

Testing OrgsCertAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	openapiclient "github.com/tmunzer/mistapi-go/sdk"
)

func Test_mistapigo_OrgsCertAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsCertAPIService ClearOrgCertificates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsCertAPI.ClearOrgCertificates(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsCertAPIService GetOrgCertificates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsCertAPI.GetOrgCertificates(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsCertAPIService TruncateOrgCrlFile", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsCertAPI.TruncateOrgCrlFile(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
