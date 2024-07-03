# Delivery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalEmails** | Pointer to **[]string** | List of additional email string to deliver the alarms via emails | [optional] 
**Enabled** | **bool** | Whether to enable the alarm delivery via emails or not | 
**ToOrgAdmins** | Pointer to **bool** | Whether to deliver the alarms via emails to Org admins or not | [optional] 
**ToSiteAdmins** | Pointer to **bool** | Whether to deliver the alarms via emails to Site admins or not | [optional] 

## Methods

### NewDelivery

`func NewDelivery(enabled bool, ) *Delivery`

NewDelivery instantiates a new Delivery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeliveryWithDefaults

`func NewDeliveryWithDefaults() *Delivery`

NewDeliveryWithDefaults instantiates a new Delivery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalEmails

`func (o *Delivery) GetAdditionalEmails() []string`

GetAdditionalEmails returns the AdditionalEmails field if non-nil, zero value otherwise.

### GetAdditionalEmailsOk

`func (o *Delivery) GetAdditionalEmailsOk() (*[]string, bool)`

GetAdditionalEmailsOk returns a tuple with the AdditionalEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEmails

`func (o *Delivery) SetAdditionalEmails(v []string)`

SetAdditionalEmails sets AdditionalEmails field to given value.

### HasAdditionalEmails

`func (o *Delivery) HasAdditionalEmails() bool`

HasAdditionalEmails returns a boolean if a field has been set.

### GetEnabled

`func (o *Delivery) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Delivery) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Delivery) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetToOrgAdmins

`func (o *Delivery) GetToOrgAdmins() bool`

GetToOrgAdmins returns the ToOrgAdmins field if non-nil, zero value otherwise.

### GetToOrgAdminsOk

`func (o *Delivery) GetToOrgAdminsOk() (*bool, bool)`

GetToOrgAdminsOk returns a tuple with the ToOrgAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToOrgAdmins

`func (o *Delivery) SetToOrgAdmins(v bool)`

SetToOrgAdmins sets ToOrgAdmins field to given value.

### HasToOrgAdmins

`func (o *Delivery) HasToOrgAdmins() bool`

HasToOrgAdmins returns a boolean if a field has been set.

### GetToSiteAdmins

`func (o *Delivery) GetToSiteAdmins() bool`

GetToSiteAdmins returns the ToSiteAdmins field if non-nil, zero value otherwise.

### GetToSiteAdminsOk

`func (o *Delivery) GetToSiteAdminsOk() (*bool, bool)`

GetToSiteAdminsOk returns a tuple with the ToSiteAdmins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSiteAdmins

`func (o *Delivery) SetToSiteAdmins(v bool)`

SetToSiteAdmins sets ToSiteAdmins field to given value.

### HasToSiteAdmins

`func (o *Delivery) HasToSiteAdmins() bool`

HasToSiteAdmins returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


