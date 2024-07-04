# WebhookDelivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | error message, if there is one | [optional] 
**Id** | Pointer to **string** | delivery ID | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**ReqHeaders** | Pointer to **string** | HTTP request headers | [optional] 
**ReqPayload** | Pointer to **string** | HTTP request payload | [optional] 
**ReqUrl** | Pointer to **string** | HTTP request URL | [optional] 
**RespBody** | Pointer to **string** | HTTP response body | [optional] 
**RespHeaders** | Pointer to **string** | HTTP response headers | [optional] 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Status** | Pointer to [**WebhookDeliveryStatus**](WebhookDeliveryStatus.md) |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**Timestamp** | Pointer to **float32** |  | [optional] 
**Topic** | Pointer to [**WebhookDeliveryTopic**](WebhookDeliveryTopic.md) |  | [optional] 
**WebhookId** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookDelivery

`func NewWebhookDelivery() *WebhookDelivery`

NewWebhookDelivery instantiates a new WebhookDelivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookDeliveryWithDefaults

`func NewWebhookDeliveryWithDefaults() *WebhookDelivery`

NewWebhookDeliveryWithDefaults instantiates a new WebhookDelivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *WebhookDelivery) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookDelivery) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookDelivery) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookDelivery) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *WebhookDelivery) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookDelivery) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookDelivery) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookDelivery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrgId

`func (o *WebhookDelivery) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *WebhookDelivery) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *WebhookDelivery) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *WebhookDelivery) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetReqHeaders

`func (o *WebhookDelivery) GetReqHeaders() string`

GetReqHeaders returns the ReqHeaders field if non-nil, zero value otherwise.

### GetReqHeadersOk

`func (o *WebhookDelivery) GetReqHeadersOk() (*string, bool)`

GetReqHeadersOk returns a tuple with the ReqHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqHeaders

`func (o *WebhookDelivery) SetReqHeaders(v string)`

SetReqHeaders sets ReqHeaders field to given value.

### HasReqHeaders

`func (o *WebhookDelivery) HasReqHeaders() bool`

HasReqHeaders returns a boolean if a field has been set.

### GetReqPayload

`func (o *WebhookDelivery) GetReqPayload() string`

GetReqPayload returns the ReqPayload field if non-nil, zero value otherwise.

### GetReqPayloadOk

`func (o *WebhookDelivery) GetReqPayloadOk() (*string, bool)`

GetReqPayloadOk returns a tuple with the ReqPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqPayload

`func (o *WebhookDelivery) SetReqPayload(v string)`

SetReqPayload sets ReqPayload field to given value.

### HasReqPayload

`func (o *WebhookDelivery) HasReqPayload() bool`

HasReqPayload returns a boolean if a field has been set.

### GetReqUrl

`func (o *WebhookDelivery) GetReqUrl() string`

GetReqUrl returns the ReqUrl field if non-nil, zero value otherwise.

### GetReqUrlOk

`func (o *WebhookDelivery) GetReqUrlOk() (*string, bool)`

GetReqUrlOk returns a tuple with the ReqUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqUrl

`func (o *WebhookDelivery) SetReqUrl(v string)`

SetReqUrl sets ReqUrl field to given value.

### HasReqUrl

`func (o *WebhookDelivery) HasReqUrl() bool`

HasReqUrl returns a boolean if a field has been set.

### GetRespBody

`func (o *WebhookDelivery) GetRespBody() string`

GetRespBody returns the RespBody field if non-nil, zero value otherwise.

### GetRespBodyOk

`func (o *WebhookDelivery) GetRespBodyOk() (*string, bool)`

GetRespBodyOk returns a tuple with the RespBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespBody

`func (o *WebhookDelivery) SetRespBody(v string)`

SetRespBody sets RespBody field to given value.

### HasRespBody

`func (o *WebhookDelivery) HasRespBody() bool`

HasRespBody returns a boolean if a field has been set.

### GetRespHeaders

`func (o *WebhookDelivery) GetRespHeaders() string`

GetRespHeaders returns the RespHeaders field if non-nil, zero value otherwise.

### GetRespHeadersOk

`func (o *WebhookDelivery) GetRespHeadersOk() (*string, bool)`

GetRespHeadersOk returns a tuple with the RespHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespHeaders

`func (o *WebhookDelivery) SetRespHeaders(v string)`

SetRespHeaders sets RespHeaders field to given value.

### HasRespHeaders

`func (o *WebhookDelivery) HasRespHeaders() bool`

HasRespHeaders returns a boolean if a field has been set.

### GetSiteId

`func (o *WebhookDelivery) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *WebhookDelivery) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *WebhookDelivery) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *WebhookDelivery) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookDelivery) GetStatus() WebhookDeliveryStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookDelivery) GetStatusOk() (*WebhookDeliveryStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookDelivery) SetStatus(v WebhookDeliveryStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookDelivery) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *WebhookDelivery) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *WebhookDelivery) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *WebhookDelivery) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *WebhookDelivery) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookDelivery) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookDelivery) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookDelivery) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *WebhookDelivery) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTopic

`func (o *WebhookDelivery) GetTopic() WebhookDeliveryTopic`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *WebhookDelivery) GetTopicOk() (*WebhookDeliveryTopic, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *WebhookDelivery) SetTopic(v WebhookDeliveryTopic)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *WebhookDelivery) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetWebhookId

`func (o *WebhookDelivery) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookDelivery) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookDelivery) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.

### HasWebhookId

`func (o *WebhookDelivery) HasWebhookId() bool`

HasWebhookId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


