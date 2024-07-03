/*
Mist API

Testing OrgsAlarmTemplatesAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package mistapigo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/tmunzer/msitapi-go"
)

func Test_mistapigo_OrgsAlarmTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OrgsAlarmTemplatesAPIService CreateOrgAlarmTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsAlarmTemplatesAPI.CreateOrgAlarmTemplate(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService DeleteOrgAlarmTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var alarmtemplateId string

		httpRes, err := apiClient.OrgsAlarmTemplatesAPI.DeleteOrgAlarmTemplate(context.Background(), orgId, alarmtemplateId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService GetOrgAlarmTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var alarmtemplateId string

		resp, httpRes, err := apiClient.OrgsAlarmTemplatesAPI.GetOrgAlarmTemplate(context.Background(), orgId, alarmtemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService ListOrgAlarmTemplates", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsAlarmTemplatesAPI.ListOrgAlarmTemplates(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService ListOrgSuppressedAlarms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		resp, httpRes, err := apiClient.OrgsAlarmTemplatesAPI.ListOrgSuppressedAlarms(context.Background(), orgId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService SuppressOrgAlarm", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsAlarmTemplatesAPI.SuppressOrgAlarm(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService UnsuppressOrgSuppressedAlarms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string

		httpRes, err := apiClient.OrgsAlarmTemplatesAPI.UnsuppressOrgSuppressedAlarms(context.Background(), orgId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test OrgsAlarmTemplatesAPIService UpdateOrgAlarmTemplate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var orgId string
		var alarmtemplateId string

		resp, httpRes, err := apiClient.OrgsAlarmTemplatesAPI.UpdateOrgAlarmTemplate(context.Background(), orgId, alarmtemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}