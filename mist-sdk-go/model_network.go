/*
Mist API

> Version: **2406.1.14** > > Date: **July 3, 2024**  ---  ### Additional Documentation * [Mist Automation Guide](https://www.juniper.net/documentation/us/en/software/mist/automation-integration/index.html) * [Mist Location SDK](https://www.juniper.net/documentation/us/en/software/mist/location_services/topics/concept/mist-how-get-mist-sdk.html) * [Mist Product Updates](https://www.mist.com/documentation/category/product-updates/)  ---  ### Helpful Resources * [API Sandbox and Exercises](https://api-class.mist.com/) * [Postman Collection, Runners and Webhook Samples](https://www.postman.com/juniper-mist/workspace/mist-systems-s-public-workspace) * [API Demo Apps](https://apps.mist-lab.fr/) * [Juniper Blog](https://blogs.juniper.net/)  --- 

API version: 2406.1.14
Contact: tmunzer@juniper.net
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mistapigo

import (
	"encoding/json"
)

// checks if the Network type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Network{}

// Network Networks are usually subnets that have cross-site significance. `networks`in Org Settings will got merged into `networks`in Site Setting. For gateways, they can be used to define Service Routes.
type Network struct {
	CreatedTime *float32 `json:"created_time,omitempty"`
	// whether to disallow Mist Devices in the network
	DisallowMistServices *bool `json:"disallow_mist_services,omitempty"`
	Gateway *string `json:"gateway,omitempty"`
	Gateway6 *string `json:"gateway6,omitempty"`
	Id *string `json:"id,omitempty"`
	InternalAccess *NetworkInternalAccess `json:"internal_access,omitempty"`
	InternetAccess *NetworkInternetAccess `json:"internet_access,omitempty"`
	// whether to allow clients in the network to talk to each other
	Isolation *bool `json:"isolation,omitempty"`
	ModifiedTime *float32 `json:"modified_time,omitempty"`
	Name *string `json:"name,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	// for a Network (usually LAN), it can be routable to other networks (e.g. OSPF)
	RoutedForNetworks []string `json:"routed_for_networks,omitempty"`
	Subnet *string `json:"subnet,omitempty"`
	Subnet6 *string `json:"subnet6,omitempty"`
	Tenants *map[string]NetworkTenant `json:"tenants,omitempty"`
	VlanId *int32 `json:"vlan_id,omitempty"`
	// Property key is the VPN name. Whether this network can be accessed from vpn
	VpnAccess *map[string]NetworkVpnAccessConfig `json:"vpn_access,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Network Network

// NewNetwork instantiates a new Network object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetwork() *Network {
	this := Network{}
	var disallowMistServices bool = false
	this.DisallowMistServices = &disallowMistServices
	return &this
}

// NewNetworkWithDefaults instantiates a new Network object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkWithDefaults() *Network {
	this := Network{}
	var disallowMistServices bool = false
	this.DisallowMistServices = &disallowMistServices
	return &this
}

// GetCreatedTime returns the CreatedTime field value if set, zero value otherwise.
func (o *Network) GetCreatedTime() float32 {
	if o == nil || IsNil(o.CreatedTime) {
		var ret float32
		return ret
	}
	return *o.CreatedTime
}

// GetCreatedTimeOk returns a tuple with the CreatedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetCreatedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.CreatedTime) {
		return nil, false
	}
	return o.CreatedTime, true
}

// HasCreatedTime returns a boolean if a field has been set.
func (o *Network) HasCreatedTime() bool {
	if o != nil && !IsNil(o.CreatedTime) {
		return true
	}

	return false
}

// SetCreatedTime gets a reference to the given float32 and assigns it to the CreatedTime field.
func (o *Network) SetCreatedTime(v float32) {
	o.CreatedTime = &v
}

// GetDisallowMistServices returns the DisallowMistServices field value if set, zero value otherwise.
func (o *Network) GetDisallowMistServices() bool {
	if o == nil || IsNil(o.DisallowMistServices) {
		var ret bool
		return ret
	}
	return *o.DisallowMistServices
}

// GetDisallowMistServicesOk returns a tuple with the DisallowMistServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetDisallowMistServicesOk() (*bool, bool) {
	if o == nil || IsNil(o.DisallowMistServices) {
		return nil, false
	}
	return o.DisallowMistServices, true
}

// HasDisallowMistServices returns a boolean if a field has been set.
func (o *Network) HasDisallowMistServices() bool {
	if o != nil && !IsNil(o.DisallowMistServices) {
		return true
	}

	return false
}

// SetDisallowMistServices gets a reference to the given bool and assigns it to the DisallowMistServices field.
func (o *Network) SetDisallowMistServices(v bool) {
	o.DisallowMistServices = &v
}

// GetGateway returns the Gateway field value if set, zero value otherwise.
func (o *Network) GetGateway() string {
	if o == nil || IsNil(o.Gateway) {
		var ret string
		return ret
	}
	return *o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetGatewayOk() (*string, bool) {
	if o == nil || IsNil(o.Gateway) {
		return nil, false
	}
	return o.Gateway, true
}

// HasGateway returns a boolean if a field has been set.
func (o *Network) HasGateway() bool {
	if o != nil && !IsNil(o.Gateway) {
		return true
	}

	return false
}

// SetGateway gets a reference to the given string and assigns it to the Gateway field.
func (o *Network) SetGateway(v string) {
	o.Gateway = &v
}

// GetGateway6 returns the Gateway6 field value if set, zero value otherwise.
func (o *Network) GetGateway6() string {
	if o == nil || IsNil(o.Gateway6) {
		var ret string
		return ret
	}
	return *o.Gateway6
}

// GetGateway6Ok returns a tuple with the Gateway6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetGateway6Ok() (*string, bool) {
	if o == nil || IsNil(o.Gateway6) {
		return nil, false
	}
	return o.Gateway6, true
}

// HasGateway6 returns a boolean if a field has been set.
func (o *Network) HasGateway6() bool {
	if o != nil && !IsNil(o.Gateway6) {
		return true
	}

	return false
}

// SetGateway6 gets a reference to the given string and assigns it to the Gateway6 field.
func (o *Network) SetGateway6(v string) {
	o.Gateway6 = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Network) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Network) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Network) SetId(v string) {
	o.Id = &v
}

// GetInternalAccess returns the InternalAccess field value if set, zero value otherwise.
func (o *Network) GetInternalAccess() NetworkInternalAccess {
	if o == nil || IsNil(o.InternalAccess) {
		var ret NetworkInternalAccess
		return ret
	}
	return *o.InternalAccess
}

// GetInternalAccessOk returns a tuple with the InternalAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetInternalAccessOk() (*NetworkInternalAccess, bool) {
	if o == nil || IsNil(o.InternalAccess) {
		return nil, false
	}
	return o.InternalAccess, true
}

// HasInternalAccess returns a boolean if a field has been set.
func (o *Network) HasInternalAccess() bool {
	if o != nil && !IsNil(o.InternalAccess) {
		return true
	}

	return false
}

// SetInternalAccess gets a reference to the given NetworkInternalAccess and assigns it to the InternalAccess field.
func (o *Network) SetInternalAccess(v NetworkInternalAccess) {
	o.InternalAccess = &v
}

// GetInternetAccess returns the InternetAccess field value if set, zero value otherwise.
func (o *Network) GetInternetAccess() NetworkInternetAccess {
	if o == nil || IsNil(o.InternetAccess) {
		var ret NetworkInternetAccess
		return ret
	}
	return *o.InternetAccess
}

// GetInternetAccessOk returns a tuple with the InternetAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetInternetAccessOk() (*NetworkInternetAccess, bool) {
	if o == nil || IsNil(o.InternetAccess) {
		return nil, false
	}
	return o.InternetAccess, true
}

// HasInternetAccess returns a boolean if a field has been set.
func (o *Network) HasInternetAccess() bool {
	if o != nil && !IsNil(o.InternetAccess) {
		return true
	}

	return false
}

// SetInternetAccess gets a reference to the given NetworkInternetAccess and assigns it to the InternetAccess field.
func (o *Network) SetInternetAccess(v NetworkInternetAccess) {
	o.InternetAccess = &v
}

// GetIsolation returns the Isolation field value if set, zero value otherwise.
func (o *Network) GetIsolation() bool {
	if o == nil || IsNil(o.Isolation) {
		var ret bool
		return ret
	}
	return *o.Isolation
}

// GetIsolationOk returns a tuple with the Isolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetIsolationOk() (*bool, bool) {
	if o == nil || IsNil(o.Isolation) {
		return nil, false
	}
	return o.Isolation, true
}

// HasIsolation returns a boolean if a field has been set.
func (o *Network) HasIsolation() bool {
	if o != nil && !IsNil(o.Isolation) {
		return true
	}

	return false
}

// SetIsolation gets a reference to the given bool and assigns it to the Isolation field.
func (o *Network) SetIsolation(v bool) {
	o.Isolation = &v
}

// GetModifiedTime returns the ModifiedTime field value if set, zero value otherwise.
func (o *Network) GetModifiedTime() float32 {
	if o == nil || IsNil(o.ModifiedTime) {
		var ret float32
		return ret
	}
	return *o.ModifiedTime
}

// GetModifiedTimeOk returns a tuple with the ModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetModifiedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ModifiedTime) {
		return nil, false
	}
	return o.ModifiedTime, true
}

// HasModifiedTime returns a boolean if a field has been set.
func (o *Network) HasModifiedTime() bool {
	if o != nil && !IsNil(o.ModifiedTime) {
		return true
	}

	return false
}

// SetModifiedTime gets a reference to the given float32 and assigns it to the ModifiedTime field.
func (o *Network) SetModifiedTime(v float32) {
	o.ModifiedTime = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Network) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Network) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Network) SetName(v string) {
	o.Name = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *Network) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *Network) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *Network) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRoutedForNetworks returns the RoutedForNetworks field value if set, zero value otherwise.
func (o *Network) GetRoutedForNetworks() []string {
	if o == nil || IsNil(o.RoutedForNetworks) {
		var ret []string
		return ret
	}
	return o.RoutedForNetworks
}

// GetRoutedForNetworksOk returns a tuple with the RoutedForNetworks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetRoutedForNetworksOk() ([]string, bool) {
	if o == nil || IsNil(o.RoutedForNetworks) {
		return nil, false
	}
	return o.RoutedForNetworks, true
}

// HasRoutedForNetworks returns a boolean if a field has been set.
func (o *Network) HasRoutedForNetworks() bool {
	if o != nil && !IsNil(o.RoutedForNetworks) {
		return true
	}

	return false
}

// SetRoutedForNetworks gets a reference to the given []string and assigns it to the RoutedForNetworks field.
func (o *Network) SetRoutedForNetworks(v []string) {
	o.RoutedForNetworks = v
}

// GetSubnet returns the Subnet field value if set, zero value otherwise.
func (o *Network) GetSubnet() string {
	if o == nil || IsNil(o.Subnet) {
		var ret string
		return ret
	}
	return *o.Subnet
}

// GetSubnetOk returns a tuple with the Subnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetSubnetOk() (*string, bool) {
	if o == nil || IsNil(o.Subnet) {
		return nil, false
	}
	return o.Subnet, true
}

// HasSubnet returns a boolean if a field has been set.
func (o *Network) HasSubnet() bool {
	if o != nil && !IsNil(o.Subnet) {
		return true
	}

	return false
}

// SetSubnet gets a reference to the given string and assigns it to the Subnet field.
func (o *Network) SetSubnet(v string) {
	o.Subnet = &v
}

// GetSubnet6 returns the Subnet6 field value if set, zero value otherwise.
func (o *Network) GetSubnet6() string {
	if o == nil || IsNil(o.Subnet6) {
		var ret string
		return ret
	}
	return *o.Subnet6
}

// GetSubnet6Ok returns a tuple with the Subnet6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetSubnet6Ok() (*string, bool) {
	if o == nil || IsNil(o.Subnet6) {
		return nil, false
	}
	return o.Subnet6, true
}

// HasSubnet6 returns a boolean if a field has been set.
func (o *Network) HasSubnet6() bool {
	if o != nil && !IsNil(o.Subnet6) {
		return true
	}

	return false
}

// SetSubnet6 gets a reference to the given string and assigns it to the Subnet6 field.
func (o *Network) SetSubnet6(v string) {
	o.Subnet6 = &v
}

// GetTenants returns the Tenants field value if set, zero value otherwise.
func (o *Network) GetTenants() map[string]NetworkTenant {
	if o == nil || IsNil(o.Tenants) {
		var ret map[string]NetworkTenant
		return ret
	}
	return *o.Tenants
}

// GetTenantsOk returns a tuple with the Tenants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetTenantsOk() (*map[string]NetworkTenant, bool) {
	if o == nil || IsNil(o.Tenants) {
		return nil, false
	}
	return o.Tenants, true
}

// HasTenants returns a boolean if a field has been set.
func (o *Network) HasTenants() bool {
	if o != nil && !IsNil(o.Tenants) {
		return true
	}

	return false
}

// SetTenants gets a reference to the given map[string]NetworkTenant and assigns it to the Tenants field.
func (o *Network) SetTenants(v map[string]NetworkTenant) {
	o.Tenants = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *Network) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *Network) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *Network) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetVpnAccess returns the VpnAccess field value if set, zero value otherwise.
func (o *Network) GetVpnAccess() map[string]NetworkVpnAccessConfig {
	if o == nil || IsNil(o.VpnAccess) {
		var ret map[string]NetworkVpnAccessConfig
		return ret
	}
	return *o.VpnAccess
}

// GetVpnAccessOk returns a tuple with the VpnAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Network) GetVpnAccessOk() (*map[string]NetworkVpnAccessConfig, bool) {
	if o == nil || IsNil(o.VpnAccess) {
		return nil, false
	}
	return o.VpnAccess, true
}

// HasVpnAccess returns a boolean if a field has been set.
func (o *Network) HasVpnAccess() bool {
	if o != nil && !IsNil(o.VpnAccess) {
		return true
	}

	return false
}

// SetVpnAccess gets a reference to the given map[string]NetworkVpnAccessConfig and assigns it to the VpnAccess field.
func (o *Network) SetVpnAccess(v map[string]NetworkVpnAccessConfig) {
	o.VpnAccess = &v
}

func (o Network) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Network) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTime) {
		toSerialize["created_time"] = o.CreatedTime
	}
	if !IsNil(o.DisallowMistServices) {
		toSerialize["disallow_mist_services"] = o.DisallowMistServices
	}
	if !IsNil(o.Gateway) {
		toSerialize["gateway"] = o.Gateway
	}
	if !IsNil(o.Gateway6) {
		toSerialize["gateway6"] = o.Gateway6
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InternalAccess) {
		toSerialize["internal_access"] = o.InternalAccess
	}
	if !IsNil(o.InternetAccess) {
		toSerialize["internet_access"] = o.InternetAccess
	}
	if !IsNil(o.Isolation) {
		toSerialize["isolation"] = o.Isolation
	}
	if !IsNil(o.ModifiedTime) {
		toSerialize["modified_time"] = o.ModifiedTime
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.RoutedForNetworks) {
		toSerialize["routed_for_networks"] = o.RoutedForNetworks
	}
	if !IsNil(o.Subnet) {
		toSerialize["subnet"] = o.Subnet
	}
	if !IsNil(o.Subnet6) {
		toSerialize["subnet6"] = o.Subnet6
	}
	if !IsNil(o.Tenants) {
		toSerialize["tenants"] = o.Tenants
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}
	if !IsNil(o.VpnAccess) {
		toSerialize["vpn_access"] = o.VpnAccess
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *Network) UnmarshalJSON(data []byte) (err error) {
	varNetwork := _Network{}

	err = json.Unmarshal(data, &varNetwork)

	if err != nil {
		return err
	}

	*o = Network(varNetwork)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "created_time")
		delete(additionalProperties, "disallow_mist_services")
		delete(additionalProperties, "gateway")
		delete(additionalProperties, "gateway6")
		delete(additionalProperties, "id")
		delete(additionalProperties, "internal_access")
		delete(additionalProperties, "internet_access")
		delete(additionalProperties, "isolation")
		delete(additionalProperties, "modified_time")
		delete(additionalProperties, "name")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "routed_for_networks")
		delete(additionalProperties, "subnet")
		delete(additionalProperties, "subnet6")
		delete(additionalProperties, "tenants")
		delete(additionalProperties, "vlan_id")
		delete(additionalProperties, "vpn_access")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNetwork struct {
	value *Network
	isSet bool
}

func (v NullableNetwork) Get() *Network {
	return v.value
}

func (v *NullableNetwork) Set(val *Network) {
	v.value = val
	v.isSet = true
}

func (v NullableNetwork) IsSet() bool {
	return v.isSet
}

func (v *NullableNetwork) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetwork(val *Network) *NullableNetwork {
	return &NullableNetwork{value: val, isSet: true}
}

func (v NullableNetwork) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetwork) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


