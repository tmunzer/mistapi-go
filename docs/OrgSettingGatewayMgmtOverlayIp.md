# OrgSettingGatewayMgmtOverlayIp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | when it&#39;s going overlay, a routable IP to overlay will be required | [optional] 
**Node1Ip** | Pointer to **string** | for SSR HA cluster, another IP for node1 will be required, too | [optional] 

## Methods

### NewOrgSettingGatewayMgmtOverlayIp

`func NewOrgSettingGatewayMgmtOverlayIp() *OrgSettingGatewayMgmtOverlayIp`

NewOrgSettingGatewayMgmtOverlayIp instantiates a new OrgSettingGatewayMgmtOverlayIp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingGatewayMgmtOverlayIpWithDefaults

`func NewOrgSettingGatewayMgmtOverlayIpWithDefaults() *OrgSettingGatewayMgmtOverlayIp`

NewOrgSettingGatewayMgmtOverlayIpWithDefaults instantiates a new OrgSettingGatewayMgmtOverlayIp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *OrgSettingGatewayMgmtOverlayIp) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *OrgSettingGatewayMgmtOverlayIp) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *OrgSettingGatewayMgmtOverlayIp) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *OrgSettingGatewayMgmtOverlayIp) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNode1Ip

`func (o *OrgSettingGatewayMgmtOverlayIp) GetNode1Ip() string`

GetNode1Ip returns the Node1Ip field if non-nil, zero value otherwise.

### GetNode1IpOk

`func (o *OrgSettingGatewayMgmtOverlayIp) GetNode1IpOk() (*string, bool)`

GetNode1IpOk returns a tuple with the Node1Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNode1Ip

`func (o *OrgSettingGatewayMgmtOverlayIp) SetNode1Ip(v string)`

SetNode1Ip sets Node1Ip field to given value.

### HasNode1Ip

`func (o *OrgSettingGatewayMgmtOverlayIp) HasNode1Ip() bool`

HasNode1Ip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


