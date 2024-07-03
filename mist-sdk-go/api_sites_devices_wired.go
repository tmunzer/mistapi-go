/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type SitesDevicesWiredAPI interface {

	/*
	DeleteSiteLocalSwitchPortConfig deleteSiteLocalSwitchPortConfig

	Sometimes HelpDesk Admin needs to change port configs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param deviceId
	@return ApiDeleteSiteLocalSwitchPortConfigRequest
	*/
	DeleteSiteLocalSwitchPortConfig(ctx context.Context, siteId string, deviceId string) ApiDeleteSiteLocalSwitchPortConfigRequest

	// DeleteSiteLocalSwitchPortConfigExecute executes the request
	DeleteSiteLocalSwitchPortConfigExecute(r ApiDeleteSiteLocalSwitchPortConfigRequest) (*http.Response, error)

	/*
	GetSiteSwitchesMetrics getSiteSwitchesMetrics

	Get version compliance metrics for managed or monitored switches

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@return ApiGetSiteSwitchesMetricsRequest
	*/
	GetSiteSwitchesMetrics(ctx context.Context, siteId string) ApiGetSiteSwitchesMetricsRequest

	// GetSiteSwitchesMetricsExecute executes the request
	//  @return ResponseSwitchMetrics
	GetSiteSwitchesMetricsExecute(r ApiGetSiteSwitchesMetricsRequest) (*ResponseSwitchMetrics, *http.Response, error)

	/*
	UpdateSiteLocalSwitchPortConfig updateSiteLocalSwitchPortConfig

	Sometimes HelpDesk Admin needs to change port configs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param siteId
	@param deviceId
	@return ApiUpdateSiteLocalSwitchPortConfigRequest
	*/
	UpdateSiteLocalSwitchPortConfig(ctx context.Context, siteId string, deviceId string) ApiUpdateSiteLocalSwitchPortConfigRequest

	// UpdateSiteLocalSwitchPortConfigExecute executes the request
	UpdateSiteLocalSwitchPortConfigExecute(r ApiUpdateSiteLocalSwitchPortConfigRequest) (*http.Response, error)
}

// SitesDevicesWiredAPIService SitesDevicesWiredAPI service
type SitesDevicesWiredAPIService service

type ApiDeleteSiteLocalSwitchPortConfigRequest struct {
	ctx context.Context
	ApiService SitesDevicesWiredAPI
	siteId string
	deviceId string
}

func (r ApiDeleteSiteLocalSwitchPortConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSiteLocalSwitchPortConfigExecute(r)
}

/*
DeleteSiteLocalSwitchPortConfig deleteSiteLocalSwitchPortConfig

Sometimes HelpDesk Admin needs to change port configs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param deviceId
 @return ApiDeleteSiteLocalSwitchPortConfigRequest
*/
func (a *SitesDevicesWiredAPIService) DeleteSiteLocalSwitchPortConfig(ctx context.Context, siteId string, deviceId string) ApiDeleteSiteLocalSwitchPortConfigRequest {
	return ApiDeleteSiteLocalSwitchPortConfigRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		deviceId: deviceId,
	}
}

// Execute executes the request
func (a *SitesDevicesWiredAPIService) DeleteSiteLocalSwitchPortConfigExecute(r ApiDeleteSiteLocalSwitchPortConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesDevicesWiredAPIService.DeleteSiteLocalSwitchPortConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/devices/{device_id}/local_port_config"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponseHttp400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSiteSwitchesMetricsRequest struct {
	ctx context.Context
	ApiService SitesDevicesWiredAPI
	siteId string
	type_ *SwitchMetricType
	scope *SwitchMetricScope
	switchMac *string
}

func (r ApiGetSiteSwitchesMetricsRequest) Type_(type_ SwitchMetricType) ApiGetSiteSwitchesMetricsRequest {
	r.type_ = &type_
	return r
}

func (r ApiGetSiteSwitchesMetricsRequest) Scope(scope SwitchMetricScope) ApiGetSiteSwitchesMetricsRequest {
	r.scope = &scope
	return r
}

// switch mac, used only with metric &#x60;type&#x60;&#x3D;&#x3D;&#x60;active_ports_summary&#x60;
func (r ApiGetSiteSwitchesMetricsRequest) SwitchMac(switchMac string) ApiGetSiteSwitchesMetricsRequest {
	r.switchMac = &switchMac
	return r
}

func (r ApiGetSiteSwitchesMetricsRequest) Execute() (*ResponseSwitchMetrics, *http.Response, error) {
	return r.ApiService.GetSiteSwitchesMetricsExecute(r)
}

/*
GetSiteSwitchesMetrics getSiteSwitchesMetrics

Get version compliance metrics for managed or monitored switches

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @return ApiGetSiteSwitchesMetricsRequest
*/
func (a *SitesDevicesWiredAPIService) GetSiteSwitchesMetrics(ctx context.Context, siteId string) ApiGetSiteSwitchesMetricsRequest {
	return ApiGetSiteSwitchesMetricsRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
	}
}

// Execute executes the request
//  @return ResponseSwitchMetrics
func (a *SitesDevicesWiredAPIService) GetSiteSwitchesMetricsExecute(r ApiGetSiteSwitchesMetricsRequest) (*ResponseSwitchMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponseSwitchMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesDevicesWiredAPIService.GetSiteSwitchesMetrics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/stats/switches/metrics"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.type_ != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "type", r.type_, "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "")
	}
	if r.switchMac != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "switch_mac", r.switchMac, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponseHttp400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateSiteLocalSwitchPortConfigRequest struct {
	ctx context.Context
	ApiService SitesDevicesWiredAPI
	siteId string
	deviceId string
	junosPortConfig *JunosPortConfig
}

func (r ApiUpdateSiteLocalSwitchPortConfigRequest) JunosPortConfig(junosPortConfig JunosPortConfig) ApiUpdateSiteLocalSwitchPortConfigRequest {
	r.junosPortConfig = &junosPortConfig
	return r
}

func (r ApiUpdateSiteLocalSwitchPortConfigRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateSiteLocalSwitchPortConfigExecute(r)
}

/*
UpdateSiteLocalSwitchPortConfig updateSiteLocalSwitchPortConfig

Sometimes HelpDesk Admin needs to change port configs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param siteId
 @param deviceId
 @return ApiUpdateSiteLocalSwitchPortConfigRequest
*/
func (a *SitesDevicesWiredAPIService) UpdateSiteLocalSwitchPortConfig(ctx context.Context, siteId string, deviceId string) ApiUpdateSiteLocalSwitchPortConfigRequest {
	return ApiUpdateSiteLocalSwitchPortConfigRequest{
		ApiService: a,
		ctx: ctx,
		siteId: siteId,
		deviceId: deviceId,
	}
}

// Execute executes the request
func (a *SitesDevicesWiredAPIService) UpdateSiteLocalSwitchPortConfigExecute(r ApiUpdateSiteLocalSwitchPortConfigRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SitesDevicesWiredAPIService.UpdateSiteLocalSwitchPortConfig")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/sites/{site_id}/devices/{device_id}/local_port_config"
	localVarPath = strings.Replace(localVarPath, "{"+"site_id"+"}", url.PathEscape(parameterValueToString(r.siteId, "siteId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"device_id"+"}", url.PathEscape(parameterValueToString(r.deviceId, "deviceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.junosPortConfig
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["apiToken"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponseHttp400
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ResponseHttp401
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ResponseHttp403
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ResponseHttp404
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v ResponseHttp429
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}