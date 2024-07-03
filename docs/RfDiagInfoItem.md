# RfDiagInfoItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;asset&#x60;, id of the asset | [optional] 
**AssetName** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;asset&#x60;, name of the asset | [optional] 
**ClientName** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60;, hostname of the client | [optional] 
**Duration** | **int32** | recording length in seconds, max is 120 | 
**EndTime** | **int32** | timestamp of end of recording | 
**FrameCount** | **int32** | Number of frames in the output | 
**Id** | Pointer to **string** |  | [optional] [readonly] 
**Mac** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;client&#x60; or &#x60;asset&#x60;, mac of the device | [optional] 
**MapId** | **string** |  | 
**Name** | **string** |  | 
**Next** | Pointer to **string** | Optional. id of the next recoding if present. Only valid for site survey. | [optional] 
**RawEvents** | **string** | URL to a JSON file that contains array of raw location diag events | 
**Ready** | **bool** | whether it’s ready for playback | 
**SdkclientId** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, sdkclient_id of this recording | [optional] 
**SdkclientName** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, name of the sdkclient | [optional] 
**SdkclientUuid** | Pointer to **string** | if &#x60;type&#x60;&#x3D;&#x3D;&#x60;sdkclient&#x60;, device_id of sdkclient | [optional] 
**StartTime** | **int32** | timestamp of the recording (the start) | 
**Type** | [**RfClientType**](RfClientType.md) |  | 
**Url** | **string** | URL to a JSON file that contains an array of frames, each frame is the same format | 

## Methods

### NewRfDiagInfoItem

`func NewRfDiagInfoItem(duration int32, endTime int32, frameCount int32, mapId string, name string, rawEvents string, ready bool, startTime int32, type_ RfClientType, url string, ) *RfDiagInfoItem`

NewRfDiagInfoItem instantiates a new RfDiagInfoItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRfDiagInfoItemWithDefaults

`func NewRfDiagInfoItemWithDefaults() *RfDiagInfoItem`

NewRfDiagInfoItemWithDefaults instantiates a new RfDiagInfoItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *RfDiagInfoItem) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *RfDiagInfoItem) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *RfDiagInfoItem) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *RfDiagInfoItem) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetName

`func (o *RfDiagInfoItem) GetAssetName() string`

GetAssetName returns the AssetName field if non-nil, zero value otherwise.

### GetAssetNameOk

`func (o *RfDiagInfoItem) GetAssetNameOk() (*string, bool)`

GetAssetNameOk returns a tuple with the AssetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetName

`func (o *RfDiagInfoItem) SetAssetName(v string)`

SetAssetName sets AssetName field to given value.

### HasAssetName

`func (o *RfDiagInfoItem) HasAssetName() bool`

HasAssetName returns a boolean if a field has been set.

### GetClientName

`func (o *RfDiagInfoItem) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *RfDiagInfoItem) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *RfDiagInfoItem) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *RfDiagInfoItem) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetDuration

`func (o *RfDiagInfoItem) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *RfDiagInfoItem) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *RfDiagInfoItem) SetDuration(v int32)`

SetDuration sets Duration field to given value.


### GetEndTime

`func (o *RfDiagInfoItem) GetEndTime() int32`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *RfDiagInfoItem) GetEndTimeOk() (*int32, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *RfDiagInfoItem) SetEndTime(v int32)`

SetEndTime sets EndTime field to given value.


### GetFrameCount

`func (o *RfDiagInfoItem) GetFrameCount() int32`

GetFrameCount returns the FrameCount field if non-nil, zero value otherwise.

### GetFrameCountOk

`func (o *RfDiagInfoItem) GetFrameCountOk() (*int32, bool)`

GetFrameCountOk returns a tuple with the FrameCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrameCount

`func (o *RfDiagInfoItem) SetFrameCount(v int32)`

SetFrameCount sets FrameCount field to given value.


### GetId

`func (o *RfDiagInfoItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RfDiagInfoItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RfDiagInfoItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RfDiagInfoItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMac

`func (o *RfDiagInfoItem) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RfDiagInfoItem) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RfDiagInfoItem) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RfDiagInfoItem) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMapId

`func (o *RfDiagInfoItem) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *RfDiagInfoItem) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *RfDiagInfoItem) SetMapId(v string)`

SetMapId sets MapId field to given value.


### GetName

`func (o *RfDiagInfoItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RfDiagInfoItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RfDiagInfoItem) SetName(v string)`

SetName sets Name field to given value.


### GetNext

`func (o *RfDiagInfoItem) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *RfDiagInfoItem) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *RfDiagInfoItem) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *RfDiagInfoItem) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetRawEvents

`func (o *RfDiagInfoItem) GetRawEvents() string`

GetRawEvents returns the RawEvents field if non-nil, zero value otherwise.

### GetRawEventsOk

`func (o *RfDiagInfoItem) GetRawEventsOk() (*string, bool)`

GetRawEventsOk returns a tuple with the RawEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawEvents

`func (o *RfDiagInfoItem) SetRawEvents(v string)`

SetRawEvents sets RawEvents field to given value.


### GetReady

`func (o *RfDiagInfoItem) GetReady() bool`

GetReady returns the Ready field if non-nil, zero value otherwise.

### GetReadyOk

`func (o *RfDiagInfoItem) GetReadyOk() (*bool, bool)`

GetReadyOk returns a tuple with the Ready field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReady

`func (o *RfDiagInfoItem) SetReady(v bool)`

SetReady sets Ready field to given value.


### GetSdkclientId

`func (o *RfDiagInfoItem) GetSdkclientId() string`

GetSdkclientId returns the SdkclientId field if non-nil, zero value otherwise.

### GetSdkclientIdOk

`func (o *RfDiagInfoItem) GetSdkclientIdOk() (*string, bool)`

GetSdkclientIdOk returns a tuple with the SdkclientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclientId

`func (o *RfDiagInfoItem) SetSdkclientId(v string)`

SetSdkclientId sets SdkclientId field to given value.

### HasSdkclientId

`func (o *RfDiagInfoItem) HasSdkclientId() bool`

HasSdkclientId returns a boolean if a field has been set.

### GetSdkclientName

`func (o *RfDiagInfoItem) GetSdkclientName() string`

GetSdkclientName returns the SdkclientName field if non-nil, zero value otherwise.

### GetSdkclientNameOk

`func (o *RfDiagInfoItem) GetSdkclientNameOk() (*string, bool)`

GetSdkclientNameOk returns a tuple with the SdkclientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclientName

`func (o *RfDiagInfoItem) SetSdkclientName(v string)`

SetSdkclientName sets SdkclientName field to given value.

### HasSdkclientName

`func (o *RfDiagInfoItem) HasSdkclientName() bool`

HasSdkclientName returns a boolean if a field has been set.

### GetSdkclientUuid

`func (o *RfDiagInfoItem) GetSdkclientUuid() string`

GetSdkclientUuid returns the SdkclientUuid field if non-nil, zero value otherwise.

### GetSdkclientUuidOk

`func (o *RfDiagInfoItem) GetSdkclientUuidOk() (*string, bool)`

GetSdkclientUuidOk returns a tuple with the SdkclientUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkclientUuid

`func (o *RfDiagInfoItem) SetSdkclientUuid(v string)`

SetSdkclientUuid sets SdkclientUuid field to given value.

### HasSdkclientUuid

`func (o *RfDiagInfoItem) HasSdkclientUuid() bool`

HasSdkclientUuid returns a boolean if a field has been set.

### GetStartTime

`func (o *RfDiagInfoItem) GetStartTime() int32`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *RfDiagInfoItem) GetStartTimeOk() (*int32, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *RfDiagInfoItem) SetStartTime(v int32)`

SetStartTime sets StartTime field to given value.


### GetType

`func (o *RfDiagInfoItem) GetType() RfClientType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RfDiagInfoItem) GetTypeOk() (*RfClientType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RfDiagInfoItem) SetType(v RfClientType)`

SetType sets Type field to given value.


### GetUrl

`func (o *RfDiagInfoItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RfDiagInfoItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RfDiagInfoItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


