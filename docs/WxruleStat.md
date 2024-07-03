# WxruleStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**WxruleStatsAction**](WxruleStatsAction.md) |  | 
**ClientMac** | **[]string** |  | 
**DstAllowWxtags** | **[]string** |  | 
**DstDenyWxtags** | **[]string** |  | 
**DstWxtags** | **[]string** |  | 
**Name** | **string** |  | 
**Order** | **int32** |  | 
**SrcWxtags** | **[]string** |  | 
**Usage** | [**map[string]WxruleStatsUsageProperties**](WxruleStatsUsageProperties.md) |  | 

## Methods

### NewWxruleStat

`func NewWxruleStat(action WxruleStatsAction, clientMac []string, dstAllowWxtags []string, dstDenyWxtags []string, dstWxtags []string, name string, order int32, srcWxtags []string, usage map[string]WxruleStatsUsageProperties, ) *WxruleStat`

NewWxruleStat instantiates a new WxruleStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWxruleStatWithDefaults

`func NewWxruleStatWithDefaults() *WxruleStat`

NewWxruleStatWithDefaults instantiates a new WxruleStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WxruleStat) GetAction() WxruleStatsAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WxruleStat) GetActionOk() (*WxruleStatsAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WxruleStat) SetAction(v WxruleStatsAction)`

SetAction sets Action field to given value.


### GetClientMac

`func (o *WxruleStat) GetClientMac() []string`

GetClientMac returns the ClientMac field if non-nil, zero value otherwise.

### GetClientMacOk

`func (o *WxruleStat) GetClientMacOk() (*[]string, bool)`

GetClientMacOk returns a tuple with the ClientMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMac

`func (o *WxruleStat) SetClientMac(v []string)`

SetClientMac sets ClientMac field to given value.


### GetDstAllowWxtags

`func (o *WxruleStat) GetDstAllowWxtags() []string`

GetDstAllowWxtags returns the DstAllowWxtags field if non-nil, zero value otherwise.

### GetDstAllowWxtagsOk

`func (o *WxruleStat) GetDstAllowWxtagsOk() (*[]string, bool)`

GetDstAllowWxtagsOk returns a tuple with the DstAllowWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstAllowWxtags

`func (o *WxruleStat) SetDstAllowWxtags(v []string)`

SetDstAllowWxtags sets DstAllowWxtags field to given value.


### GetDstDenyWxtags

`func (o *WxruleStat) GetDstDenyWxtags() []string`

GetDstDenyWxtags returns the DstDenyWxtags field if non-nil, zero value otherwise.

### GetDstDenyWxtagsOk

`func (o *WxruleStat) GetDstDenyWxtagsOk() (*[]string, bool)`

GetDstDenyWxtagsOk returns a tuple with the DstDenyWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstDenyWxtags

`func (o *WxruleStat) SetDstDenyWxtags(v []string)`

SetDstDenyWxtags sets DstDenyWxtags field to given value.


### GetDstWxtags

`func (o *WxruleStat) GetDstWxtags() []string`

GetDstWxtags returns the DstWxtags field if non-nil, zero value otherwise.

### GetDstWxtagsOk

`func (o *WxruleStat) GetDstWxtagsOk() (*[]string, bool)`

GetDstWxtagsOk returns a tuple with the DstWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstWxtags

`func (o *WxruleStat) SetDstWxtags(v []string)`

SetDstWxtags sets DstWxtags field to given value.


### GetName

`func (o *WxruleStat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WxruleStat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WxruleStat) SetName(v string)`

SetName sets Name field to given value.


### GetOrder

`func (o *WxruleStat) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *WxruleStat) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *WxruleStat) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetSrcWxtags

`func (o *WxruleStat) GetSrcWxtags() []string`

GetSrcWxtags returns the SrcWxtags field if non-nil, zero value otherwise.

### GetSrcWxtagsOk

`func (o *WxruleStat) GetSrcWxtagsOk() (*[]string, bool)`

GetSrcWxtagsOk returns a tuple with the SrcWxtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcWxtags

`func (o *WxruleStat) SetSrcWxtags(v []string)`

SetSrcWxtags sets SrcWxtags field to given value.


### GetUsage

`func (o *WxruleStat) GetUsage() map[string]WxruleStatsUsageProperties`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *WxruleStat) GetUsageOk() (*map[string]WxruleStatsUsageProperties, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *WxruleStat) SetUsage(v map[string]WxruleStatsUsageProperties)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


