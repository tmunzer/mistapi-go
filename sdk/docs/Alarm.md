# Alarm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckAdminId** | Pointer to **string** | UUID of the admin who acked the alarm | [optional] 
**AckAdminName** | Pointer to **string** | Name &amp; Email ID of the admin who acked the alarm | [optional] 
**Acked** | Pointer to **bool** | Whether the alarm is acked or not | [optional] 
**AckedTime** | Pointer to **int32** | Epoch (seconds) when the alarm was acked | [optional] [readonly] 
**Aps** | Pointer to **[]string** | additional information: List of MACs of the APs | [optional] 
**Bssids** | Pointer to **[]string** | List of BSSIDs | [optional] 
**Count** | **int32** | Number of incident within an alarm window | [readonly] 
**Gateways** | Pointer to **[]string** | additional information: List of MACs of the gateways | [optional] 
**Group** | **string** | Group of the alarm | 
**Hostnames** | Pointer to **[]string** | additional information: List of Hostnames of the devices (AP/Switch/Gateway) | [optional] 
**Id** | **string** | UUID of the alarm | [readonly] 
**LastSeen** | **int32** | Epoch (seconds) of the last incident/alarm within an alarm window | 
**Note** | Pointer to **string** | Text describing the alarm | [optional] 
**OrgId** | Pointer to **string** |  | [optional] [readonly] 
**Severity** | **string** | Severity of the alarm | 
**SiteId** | Pointer to **string** |  | [optional] [readonly] 
**Ssids** | Pointer to **[]string** | List of SSIDs | [optional] 
**Switches** | Pointer to **[]string** | additional information: List of MACs of the switches | [optional] 
**Timestamp** | **int32** | Epoch (seconds) of the first incident/alarm | [readonly] 
**Type** | **string** | Key-name of the alarm type | [readonly] 

## Methods

### NewAlarm

`func NewAlarm(count int32, group string, id string, lastSeen int32, severity string, timestamp int32, type_ string, ) *Alarm`

NewAlarm instantiates a new Alarm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlarmWithDefaults

`func NewAlarmWithDefaults() *Alarm`

NewAlarmWithDefaults instantiates a new Alarm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckAdminId

`func (o *Alarm) GetAckAdminId() string`

GetAckAdminId returns the AckAdminId field if non-nil, zero value otherwise.

### GetAckAdminIdOk

`func (o *Alarm) GetAckAdminIdOk() (*string, bool)`

GetAckAdminIdOk returns a tuple with the AckAdminId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAdminId

`func (o *Alarm) SetAckAdminId(v string)`

SetAckAdminId sets AckAdminId field to given value.

### HasAckAdminId

`func (o *Alarm) HasAckAdminId() bool`

HasAckAdminId returns a boolean if a field has been set.

### GetAckAdminName

`func (o *Alarm) GetAckAdminName() string`

GetAckAdminName returns the AckAdminName field if non-nil, zero value otherwise.

### GetAckAdminNameOk

`func (o *Alarm) GetAckAdminNameOk() (*string, bool)`

GetAckAdminNameOk returns a tuple with the AckAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckAdminName

`func (o *Alarm) SetAckAdminName(v string)`

SetAckAdminName sets AckAdminName field to given value.

### HasAckAdminName

`func (o *Alarm) HasAckAdminName() bool`

HasAckAdminName returns a boolean if a field has been set.

### GetAcked

`func (o *Alarm) GetAcked() bool`

GetAcked returns the Acked field if non-nil, zero value otherwise.

### GetAckedOk

`func (o *Alarm) GetAckedOk() (*bool, bool)`

GetAckedOk returns a tuple with the Acked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcked

`func (o *Alarm) SetAcked(v bool)`

SetAcked sets Acked field to given value.

### HasAcked

`func (o *Alarm) HasAcked() bool`

HasAcked returns a boolean if a field has been set.

### GetAckedTime

`func (o *Alarm) GetAckedTime() int32`

GetAckedTime returns the AckedTime field if non-nil, zero value otherwise.

### GetAckedTimeOk

`func (o *Alarm) GetAckedTimeOk() (*int32, bool)`

GetAckedTimeOk returns a tuple with the AckedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckedTime

`func (o *Alarm) SetAckedTime(v int32)`

SetAckedTime sets AckedTime field to given value.

### HasAckedTime

`func (o *Alarm) HasAckedTime() bool`

HasAckedTime returns a boolean if a field has been set.

### GetAps

`func (o *Alarm) GetAps() []string`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *Alarm) GetApsOk() (*[]string, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *Alarm) SetAps(v []string)`

SetAps sets Aps field to given value.

### HasAps

`func (o *Alarm) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetBssids

`func (o *Alarm) GetBssids() []string`

GetBssids returns the Bssids field if non-nil, zero value otherwise.

### GetBssidsOk

`func (o *Alarm) GetBssidsOk() (*[]string, bool)`

GetBssidsOk returns a tuple with the Bssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBssids

`func (o *Alarm) SetBssids(v []string)`

SetBssids sets Bssids field to given value.

### HasBssids

`func (o *Alarm) HasBssids() bool`

HasBssids returns a boolean if a field has been set.

### GetCount

`func (o *Alarm) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Alarm) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Alarm) SetCount(v int32)`

SetCount sets Count field to given value.


### GetGateways

`func (o *Alarm) GetGateways() []string`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *Alarm) GetGatewaysOk() (*[]string, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *Alarm) SetGateways(v []string)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *Alarm) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetGroup

`func (o *Alarm) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Alarm) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Alarm) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetHostnames

`func (o *Alarm) GetHostnames() []string`

GetHostnames returns the Hostnames field if non-nil, zero value otherwise.

### GetHostnamesOk

`func (o *Alarm) GetHostnamesOk() (*[]string, bool)`

GetHostnamesOk returns a tuple with the Hostnames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostnames

`func (o *Alarm) SetHostnames(v []string)`

SetHostnames sets Hostnames field to given value.

### HasHostnames

`func (o *Alarm) HasHostnames() bool`

HasHostnames returns a boolean if a field has been set.

### GetId

`func (o *Alarm) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Alarm) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Alarm) SetId(v string)`

SetId sets Id field to given value.


### GetLastSeen

`func (o *Alarm) GetLastSeen() int32`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *Alarm) GetLastSeenOk() (*int32, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *Alarm) SetLastSeen(v int32)`

SetLastSeen sets LastSeen field to given value.


### GetNote

`func (o *Alarm) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Alarm) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Alarm) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Alarm) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOrgId

`func (o *Alarm) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *Alarm) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *Alarm) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.

### HasOrgId

`func (o *Alarm) HasOrgId() bool`

HasOrgId returns a boolean if a field has been set.

### GetSeverity

`func (o *Alarm) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Alarm) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Alarm) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetSiteId

`func (o *Alarm) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *Alarm) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *Alarm) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *Alarm) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSsids

`func (o *Alarm) GetSsids() []string`

GetSsids returns the Ssids field if non-nil, zero value otherwise.

### GetSsidsOk

`func (o *Alarm) GetSsidsOk() (*[]string, bool)`

GetSsidsOk returns a tuple with the Ssids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsids

`func (o *Alarm) SetSsids(v []string)`

SetSsids sets Ssids field to given value.

### HasSsids

`func (o *Alarm) HasSsids() bool`

HasSsids returns a boolean if a field has been set.

### GetSwitches

`func (o *Alarm) GetSwitches() []string`

GetSwitches returns the Switches field if non-nil, zero value otherwise.

### GetSwitchesOk

`func (o *Alarm) GetSwitchesOk() (*[]string, bool)`

GetSwitchesOk returns a tuple with the Switches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitches

`func (o *Alarm) SetSwitches(v []string)`

SetSwitches sets Switches field to given value.

### HasSwitches

`func (o *Alarm) HasSwitches() bool`

HasSwitches returns a boolean if a field has been set.

### GetTimestamp

`func (o *Alarm) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Alarm) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Alarm) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *Alarm) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Alarm) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Alarm) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


