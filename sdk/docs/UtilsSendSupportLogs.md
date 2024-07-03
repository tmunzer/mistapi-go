# UtilsSendSupportLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**UtilsSendSupportLogsInfo**](UtilsSendSupportLogsInfo.md) |  | [optional] [default to UTILSSENDSUPPORTLOGSINFO_FULL]
**Node** | Pointer to **string** | optional: for SSR, if node is not present, both nodes support files are uploaded | [optional] 
**NumMessagesFiles** | Pointer to **int32** | optional: number of most recent messages files to upload. | [optional] [default to 1]

## Methods

### NewUtilsSendSupportLogs

`func NewUtilsSendSupportLogs() *UtilsSendSupportLogs`

NewUtilsSendSupportLogs instantiates a new UtilsSendSupportLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilsSendSupportLogsWithDefaults

`func NewUtilsSendSupportLogsWithDefaults() *UtilsSendSupportLogs`

NewUtilsSendSupportLogsWithDefaults instantiates a new UtilsSendSupportLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *UtilsSendSupportLogs) GetInfo() UtilsSendSupportLogsInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *UtilsSendSupportLogs) GetInfoOk() (*UtilsSendSupportLogsInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *UtilsSendSupportLogs) SetInfo(v UtilsSendSupportLogsInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *UtilsSendSupportLogs) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetNode

`func (o *UtilsSendSupportLogs) GetNode() string`

GetNode returns the Node field if non-nil, zero value otherwise.

### GetNodeOk

`func (o *UtilsSendSupportLogs) GetNodeOk() (*string, bool)`

GetNodeOk returns a tuple with the Node field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode

`func (o *UtilsSendSupportLogs) SetNode(v string)`

SetNode sets Node field to given value.

### HasNode

`func (o *UtilsSendSupportLogs) HasNode() bool`

HasNode returns a boolean if a field has been set.

### GetNumMessagesFiles

`func (o *UtilsSendSupportLogs) GetNumMessagesFiles() int32`

GetNumMessagesFiles returns the NumMessagesFiles field if non-nil, zero value otherwise.

### GetNumMessagesFilesOk

`func (o *UtilsSendSupportLogs) GetNumMessagesFilesOk() (*int32, bool)`

GetNumMessagesFilesOk returns a tuple with the NumMessagesFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumMessagesFiles

`func (o *UtilsSendSupportLogs) SetNumMessagesFiles(v int32)`

SetNumMessagesFiles sets NumMessagesFiles field to given value.

### HasNumMessagesFiles

`func (o *UtilsSendSupportLogs) HasNumMessagesFiles() bool`

HasNumMessagesFiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


