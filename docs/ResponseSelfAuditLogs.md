# ResponseSelfAuditLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | **int32** |  | 
**Limit** | **int32** |  | 
**Page** | **int32** |  | 
**Results** | [**[]AuditLog**](AuditLog.md) |  | 
**Start** | **int32** |  | 
**Total** | **int32** |  | 

## Methods

### NewResponseSelfAuditLogs

`func NewResponseSelfAuditLogs(end int32, limit int32, page int32, results []AuditLog, start int32, total int32, ) *ResponseSelfAuditLogs`

NewResponseSelfAuditLogs instantiates a new ResponseSelfAuditLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseSelfAuditLogsWithDefaults

`func NewResponseSelfAuditLogsWithDefaults() *ResponseSelfAuditLogs`

NewResponseSelfAuditLogsWithDefaults instantiates a new ResponseSelfAuditLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ResponseSelfAuditLogs) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ResponseSelfAuditLogs) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ResponseSelfAuditLogs) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetLimit

`func (o *ResponseSelfAuditLogs) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *ResponseSelfAuditLogs) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *ResponseSelfAuditLogs) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetPage

`func (o *ResponseSelfAuditLogs) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResponseSelfAuditLogs) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResponseSelfAuditLogs) SetPage(v int32)`

SetPage sets Page field to given value.


### GetResults

`func (o *ResponseSelfAuditLogs) GetResults() []AuditLog`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ResponseSelfAuditLogs) GetResultsOk() (*[]AuditLog, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ResponseSelfAuditLogs) SetResults(v []AuditLog)`

SetResults sets Results field to given value.


### GetStart

`func (o *ResponseSelfAuditLogs) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ResponseSelfAuditLogs) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ResponseSelfAuditLogs) SetStart(v int32)`

SetStart sets Start field to given value.


### GetTotal

`func (o *ResponseSelfAuditLogs) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ResponseSelfAuditLogs) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ResponseSelfAuditLogs) SetTotal(v int32)`

SetTotal sets Total field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


