# Guest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorized** | Pointer to **bool** | whether the guest is current authorized | [optional] 
**AuthorizedExpiringTime** | Pointer to **float32** | when the authorization would expire | [optional] 
**AuthorizedTime** | Pointer to **float32** | when the guest was authorized | [optional] 
**Company** | Pointer to **string** | optional, the info provided by user | [optional] 
**Email** | Pointer to **string** | optional, the info provided by user | [optional] 
**Field1** | Pointer to **string** | optional, the info provided by user | [optional] 
**Field2** | Pointer to **string** |  | [optional] 
**Field3** | Pointer to **string** |  | [optional] 
**Field4** | Pointer to **string** |  | [optional] 
**Mac** | Pointer to **string** | mac | [optional] 
**Minutes** | Pointer to **int32** | minutes, the maximum is 259200 (180 days) | [optional] 
**Name** | Pointer to **string** | optional, the info provided by user | [optional] 
**Ssid** | Pointer to **string** |  | [optional] [readonly] 
**WlanId** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewGuest

`func NewGuest() *Guest`

NewGuest instantiates a new Guest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuestWithDefaults

`func NewGuestWithDefaults() *Guest`

NewGuestWithDefaults instantiates a new Guest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorized

`func (o *Guest) GetAuthorized() bool`

GetAuthorized returns the Authorized field if non-nil, zero value otherwise.

### GetAuthorizedOk

`func (o *Guest) GetAuthorizedOk() (*bool, bool)`

GetAuthorizedOk returns a tuple with the Authorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorized

`func (o *Guest) SetAuthorized(v bool)`

SetAuthorized sets Authorized field to given value.

### HasAuthorized

`func (o *Guest) HasAuthorized() bool`

HasAuthorized returns a boolean if a field has been set.

### GetAuthorizedExpiringTime

`func (o *Guest) GetAuthorizedExpiringTime() float32`

GetAuthorizedExpiringTime returns the AuthorizedExpiringTime field if non-nil, zero value otherwise.

### GetAuthorizedExpiringTimeOk

`func (o *Guest) GetAuthorizedExpiringTimeOk() (*float32, bool)`

GetAuthorizedExpiringTimeOk returns a tuple with the AuthorizedExpiringTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedExpiringTime

`func (o *Guest) SetAuthorizedExpiringTime(v float32)`

SetAuthorizedExpiringTime sets AuthorizedExpiringTime field to given value.

### HasAuthorizedExpiringTime

`func (o *Guest) HasAuthorizedExpiringTime() bool`

HasAuthorizedExpiringTime returns a boolean if a field has been set.

### GetAuthorizedTime

`func (o *Guest) GetAuthorizedTime() float32`

GetAuthorizedTime returns the AuthorizedTime field if non-nil, zero value otherwise.

### GetAuthorizedTimeOk

`func (o *Guest) GetAuthorizedTimeOk() (*float32, bool)`

GetAuthorizedTimeOk returns a tuple with the AuthorizedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedTime

`func (o *Guest) SetAuthorizedTime(v float32)`

SetAuthorizedTime sets AuthorizedTime field to given value.

### HasAuthorizedTime

`func (o *Guest) HasAuthorizedTime() bool`

HasAuthorizedTime returns a boolean if a field has been set.

### GetCompany

`func (o *Guest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Guest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Guest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *Guest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *Guest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Guest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Guest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Guest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetField1

`func (o *Guest) GetField1() string`

GetField1 returns the Field1 field if non-nil, zero value otherwise.

### GetField1Ok

`func (o *Guest) GetField1Ok() (*string, bool)`

GetField1Ok returns a tuple with the Field1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField1

`func (o *Guest) SetField1(v string)`

SetField1 sets Field1 field to given value.

### HasField1

`func (o *Guest) HasField1() bool`

HasField1 returns a boolean if a field has been set.

### GetField2

`func (o *Guest) GetField2() string`

GetField2 returns the Field2 field if non-nil, zero value otherwise.

### GetField2Ok

`func (o *Guest) GetField2Ok() (*string, bool)`

GetField2Ok returns a tuple with the Field2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField2

`func (o *Guest) SetField2(v string)`

SetField2 sets Field2 field to given value.

### HasField2

`func (o *Guest) HasField2() bool`

HasField2 returns a boolean if a field has been set.

### GetField3

`func (o *Guest) GetField3() string`

GetField3 returns the Field3 field if non-nil, zero value otherwise.

### GetField3Ok

`func (o *Guest) GetField3Ok() (*string, bool)`

GetField3Ok returns a tuple with the Field3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField3

`func (o *Guest) SetField3(v string)`

SetField3 sets Field3 field to given value.

### HasField3

`func (o *Guest) HasField3() bool`

HasField3 returns a boolean if a field has been set.

### GetField4

`func (o *Guest) GetField4() string`

GetField4 returns the Field4 field if non-nil, zero value otherwise.

### GetField4Ok

`func (o *Guest) GetField4Ok() (*string, bool)`

GetField4Ok returns a tuple with the Field4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField4

`func (o *Guest) SetField4(v string)`

SetField4 sets Field4 field to given value.

### HasField4

`func (o *Guest) HasField4() bool`

HasField4 returns a boolean if a field has been set.

### GetMac

`func (o *Guest) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *Guest) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *Guest) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *Guest) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMinutes

`func (o *Guest) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *Guest) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *Guest) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *Guest) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetName

`func (o *Guest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Guest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Guest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Guest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSsid

`func (o *Guest) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *Guest) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *Guest) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *Guest) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetWlanId

`func (o *Guest) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *Guest) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *Guest) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *Guest) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


