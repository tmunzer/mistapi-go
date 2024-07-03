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

// checks if the BgpStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BgpStats{}

// BgpStats struct for BgpStats
type BgpStats struct {
	// if this is created for evpn overlay
	EvpnOverlay *bool `json:"evpn_overlay,omitempty"`
	// if this is created for overlay
	ForOverlay *bool `json:"for_overlay,omitempty"`
	// AS
	LocalAs *int32 `json:"local_as,omitempty"`
	// router mac address
	Mac *string `json:"mac,omitempty"`
	Neighbor *string `json:"neighbor,omitempty"`
	NeighborAs *int32 `json:"neighbor_as,omitempty"`
	// if it's another device in the same org
	NeighborMac *string `json:"neighbor_mac,omitempty"`
	// node0/node1
	Node *string `json:"node,omitempty"`
	OrgId *string `json:"org_id,omitempty"`
	RxPkts *int32 `json:"rx_pkts,omitempty"`
	// number of received routes
	RxRoutes *int32 `json:"rx_routes,omitempty"`
	SiteId *string `json:"site_id,omitempty"`
	State *BgpStatsState `json:"state,omitempty"`
	Timestamp *float32 `json:"timestamp,omitempty"`
	TxPkts *int32 `json:"tx_pkts,omitempty"`
	TxRoutes *int32 `json:"tx_routes,omitempty"`
	Up *bool `json:"up,omitempty"`
	Uptime *int32 `json:"uptime,omitempty"`
	VrfName *string `json:"vrf_name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BgpStats BgpStats

// NewBgpStats instantiates a new BgpStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBgpStats() *BgpStats {
	this := BgpStats{}
	return &this
}

// NewBgpStatsWithDefaults instantiates a new BgpStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBgpStatsWithDefaults() *BgpStats {
	this := BgpStats{}
	return &this
}

// GetEvpnOverlay returns the EvpnOverlay field value if set, zero value otherwise.
func (o *BgpStats) GetEvpnOverlay() bool {
	if o == nil || IsNil(o.EvpnOverlay) {
		var ret bool
		return ret
	}
	return *o.EvpnOverlay
}

// GetEvpnOverlayOk returns a tuple with the EvpnOverlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetEvpnOverlayOk() (*bool, bool) {
	if o == nil || IsNil(o.EvpnOverlay) {
		return nil, false
	}
	return o.EvpnOverlay, true
}

// HasEvpnOverlay returns a boolean if a field has been set.
func (o *BgpStats) HasEvpnOverlay() bool {
	if o != nil && !IsNil(o.EvpnOverlay) {
		return true
	}

	return false
}

// SetEvpnOverlay gets a reference to the given bool and assigns it to the EvpnOverlay field.
func (o *BgpStats) SetEvpnOverlay(v bool) {
	o.EvpnOverlay = &v
}

// GetForOverlay returns the ForOverlay field value if set, zero value otherwise.
func (o *BgpStats) GetForOverlay() bool {
	if o == nil || IsNil(o.ForOverlay) {
		var ret bool
		return ret
	}
	return *o.ForOverlay
}

// GetForOverlayOk returns a tuple with the ForOverlay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetForOverlayOk() (*bool, bool) {
	if o == nil || IsNil(o.ForOverlay) {
		return nil, false
	}
	return o.ForOverlay, true
}

// HasForOverlay returns a boolean if a field has been set.
func (o *BgpStats) HasForOverlay() bool {
	if o != nil && !IsNil(o.ForOverlay) {
		return true
	}

	return false
}

// SetForOverlay gets a reference to the given bool and assigns it to the ForOverlay field.
func (o *BgpStats) SetForOverlay(v bool) {
	o.ForOverlay = &v
}

// GetLocalAs returns the LocalAs field value if set, zero value otherwise.
func (o *BgpStats) GetLocalAs() int32 {
	if o == nil || IsNil(o.LocalAs) {
		var ret int32
		return ret
	}
	return *o.LocalAs
}

// GetLocalAsOk returns a tuple with the LocalAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetLocalAsOk() (*int32, bool) {
	if o == nil || IsNil(o.LocalAs) {
		return nil, false
	}
	return o.LocalAs, true
}

// HasLocalAs returns a boolean if a field has been set.
func (o *BgpStats) HasLocalAs() bool {
	if o != nil && !IsNil(o.LocalAs) {
		return true
	}

	return false
}

// SetLocalAs gets a reference to the given int32 and assigns it to the LocalAs field.
func (o *BgpStats) SetLocalAs(v int32) {
	o.LocalAs = &v
}

// GetMac returns the Mac field value if set, zero value otherwise.
func (o *BgpStats) GetMac() string {
	if o == nil || IsNil(o.Mac) {
		var ret string
		return ret
	}
	return *o.Mac
}

// GetMacOk returns a tuple with the Mac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetMacOk() (*string, bool) {
	if o == nil || IsNil(o.Mac) {
		return nil, false
	}
	return o.Mac, true
}

// HasMac returns a boolean if a field has been set.
func (o *BgpStats) HasMac() bool {
	if o != nil && !IsNil(o.Mac) {
		return true
	}

	return false
}

// SetMac gets a reference to the given string and assigns it to the Mac field.
func (o *BgpStats) SetMac(v string) {
	o.Mac = &v
}

// GetNeighbor returns the Neighbor field value if set, zero value otherwise.
func (o *BgpStats) GetNeighbor() string {
	if o == nil || IsNil(o.Neighbor) {
		var ret string
		return ret
	}
	return *o.Neighbor
}

// GetNeighborOk returns a tuple with the Neighbor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetNeighborOk() (*string, bool) {
	if o == nil || IsNil(o.Neighbor) {
		return nil, false
	}
	return o.Neighbor, true
}

// HasNeighbor returns a boolean if a field has been set.
func (o *BgpStats) HasNeighbor() bool {
	if o != nil && !IsNil(o.Neighbor) {
		return true
	}

	return false
}

// SetNeighbor gets a reference to the given string and assigns it to the Neighbor field.
func (o *BgpStats) SetNeighbor(v string) {
	o.Neighbor = &v
}

// GetNeighborAs returns the NeighborAs field value if set, zero value otherwise.
func (o *BgpStats) GetNeighborAs() int32 {
	if o == nil || IsNil(o.NeighborAs) {
		var ret int32
		return ret
	}
	return *o.NeighborAs
}

// GetNeighborAsOk returns a tuple with the NeighborAs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetNeighborAsOk() (*int32, bool) {
	if o == nil || IsNil(o.NeighborAs) {
		return nil, false
	}
	return o.NeighborAs, true
}

// HasNeighborAs returns a boolean if a field has been set.
func (o *BgpStats) HasNeighborAs() bool {
	if o != nil && !IsNil(o.NeighborAs) {
		return true
	}

	return false
}

// SetNeighborAs gets a reference to the given int32 and assigns it to the NeighborAs field.
func (o *BgpStats) SetNeighborAs(v int32) {
	o.NeighborAs = &v
}

// GetNeighborMac returns the NeighborMac field value if set, zero value otherwise.
func (o *BgpStats) GetNeighborMac() string {
	if o == nil || IsNil(o.NeighborMac) {
		var ret string
		return ret
	}
	return *o.NeighborMac
}

// GetNeighborMacOk returns a tuple with the NeighborMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetNeighborMacOk() (*string, bool) {
	if o == nil || IsNil(o.NeighborMac) {
		return nil, false
	}
	return o.NeighborMac, true
}

// HasNeighborMac returns a boolean if a field has been set.
func (o *BgpStats) HasNeighborMac() bool {
	if o != nil && !IsNil(o.NeighborMac) {
		return true
	}

	return false
}

// SetNeighborMac gets a reference to the given string and assigns it to the NeighborMac field.
func (o *BgpStats) SetNeighborMac(v string) {
	o.NeighborMac = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *BgpStats) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *BgpStats) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *BgpStats) SetNode(v string) {
	o.Node = &v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *BgpStats) GetOrgId() string {
	if o == nil || IsNil(o.OrgId) {
		var ret string
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetOrgIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrgId) {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *BgpStats) HasOrgId() bool {
	if o != nil && !IsNil(o.OrgId) {
		return true
	}

	return false
}

// SetOrgId gets a reference to the given string and assigns it to the OrgId field.
func (o *BgpStats) SetOrgId(v string) {
	o.OrgId = &v
}

// GetRxPkts returns the RxPkts field value if set, zero value otherwise.
func (o *BgpStats) GetRxPkts() int32 {
	if o == nil || IsNil(o.RxPkts) {
		var ret int32
		return ret
	}
	return *o.RxPkts
}

// GetRxPktsOk returns a tuple with the RxPkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetRxPktsOk() (*int32, bool) {
	if o == nil || IsNil(o.RxPkts) {
		return nil, false
	}
	return o.RxPkts, true
}

// HasRxPkts returns a boolean if a field has been set.
func (o *BgpStats) HasRxPkts() bool {
	if o != nil && !IsNil(o.RxPkts) {
		return true
	}

	return false
}

// SetRxPkts gets a reference to the given int32 and assigns it to the RxPkts field.
func (o *BgpStats) SetRxPkts(v int32) {
	o.RxPkts = &v
}

// GetRxRoutes returns the RxRoutes field value if set, zero value otherwise.
func (o *BgpStats) GetRxRoutes() int32 {
	if o == nil || IsNil(o.RxRoutes) {
		var ret int32
		return ret
	}
	return *o.RxRoutes
}

// GetRxRoutesOk returns a tuple with the RxRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetRxRoutesOk() (*int32, bool) {
	if o == nil || IsNil(o.RxRoutes) {
		return nil, false
	}
	return o.RxRoutes, true
}

// HasRxRoutes returns a boolean if a field has been set.
func (o *BgpStats) HasRxRoutes() bool {
	if o != nil && !IsNil(o.RxRoutes) {
		return true
	}

	return false
}

// SetRxRoutes gets a reference to the given int32 and assigns it to the RxRoutes field.
func (o *BgpStats) SetRxRoutes(v int32) {
	o.RxRoutes = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *BgpStats) GetSiteId() string {
	if o == nil || IsNil(o.SiteId) {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetSiteIdOk() (*string, bool) {
	if o == nil || IsNil(o.SiteId) {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *BgpStats) HasSiteId() bool {
	if o != nil && !IsNil(o.SiteId) {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *BgpStats) SetSiteId(v string) {
	o.SiteId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BgpStats) GetState() BgpStatsState {
	if o == nil || IsNil(o.State) {
		var ret BgpStatsState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetStateOk() (*BgpStatsState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BgpStats) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given BgpStatsState and assigns it to the State field.
func (o *BgpStats) SetState(v BgpStatsState) {
	o.State = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BgpStats) GetTimestamp() float32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret float32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BgpStats) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float32 and assigns it to the Timestamp field.
func (o *BgpStats) SetTimestamp(v float32) {
	o.Timestamp = &v
}

// GetTxPkts returns the TxPkts field value if set, zero value otherwise.
func (o *BgpStats) GetTxPkts() int32 {
	if o == nil || IsNil(o.TxPkts) {
		var ret int32
		return ret
	}
	return *o.TxPkts
}

// GetTxPktsOk returns a tuple with the TxPkts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetTxPktsOk() (*int32, bool) {
	if o == nil || IsNil(o.TxPkts) {
		return nil, false
	}
	return o.TxPkts, true
}

// HasTxPkts returns a boolean if a field has been set.
func (o *BgpStats) HasTxPkts() bool {
	if o != nil && !IsNil(o.TxPkts) {
		return true
	}

	return false
}

// SetTxPkts gets a reference to the given int32 and assigns it to the TxPkts field.
func (o *BgpStats) SetTxPkts(v int32) {
	o.TxPkts = &v
}

// GetTxRoutes returns the TxRoutes field value if set, zero value otherwise.
func (o *BgpStats) GetTxRoutes() int32 {
	if o == nil || IsNil(o.TxRoutes) {
		var ret int32
		return ret
	}
	return *o.TxRoutes
}

// GetTxRoutesOk returns a tuple with the TxRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetTxRoutesOk() (*int32, bool) {
	if o == nil || IsNil(o.TxRoutes) {
		return nil, false
	}
	return o.TxRoutes, true
}

// HasTxRoutes returns a boolean if a field has been set.
func (o *BgpStats) HasTxRoutes() bool {
	if o != nil && !IsNil(o.TxRoutes) {
		return true
	}

	return false
}

// SetTxRoutes gets a reference to the given int32 and assigns it to the TxRoutes field.
func (o *BgpStats) SetTxRoutes(v int32) {
	o.TxRoutes = &v
}

// GetUp returns the Up field value if set, zero value otherwise.
func (o *BgpStats) GetUp() bool {
	if o == nil || IsNil(o.Up) {
		var ret bool
		return ret
	}
	return *o.Up
}

// GetUpOk returns a tuple with the Up field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetUpOk() (*bool, bool) {
	if o == nil || IsNil(o.Up) {
		return nil, false
	}
	return o.Up, true
}

// HasUp returns a boolean if a field has been set.
func (o *BgpStats) HasUp() bool {
	if o != nil && !IsNil(o.Up) {
		return true
	}

	return false
}

// SetUp gets a reference to the given bool and assigns it to the Up field.
func (o *BgpStats) SetUp(v bool) {
	o.Up = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *BgpStats) GetUptime() int32 {
	if o == nil || IsNil(o.Uptime) {
		var ret int32
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetUptimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *BgpStats) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int32 and assigns it to the Uptime field.
func (o *BgpStats) SetUptime(v int32) {
	o.Uptime = &v
}

// GetVrfName returns the VrfName field value if set, zero value otherwise.
func (o *BgpStats) GetVrfName() string {
	if o == nil || IsNil(o.VrfName) {
		var ret string
		return ret
	}
	return *o.VrfName
}

// GetVrfNameOk returns a tuple with the VrfName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BgpStats) GetVrfNameOk() (*string, bool) {
	if o == nil || IsNil(o.VrfName) {
		return nil, false
	}
	return o.VrfName, true
}

// HasVrfName returns a boolean if a field has been set.
func (o *BgpStats) HasVrfName() bool {
	if o != nil && !IsNil(o.VrfName) {
		return true
	}

	return false
}

// SetVrfName gets a reference to the given string and assigns it to the VrfName field.
func (o *BgpStats) SetVrfName(v string) {
	o.VrfName = &v
}

func (o BgpStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BgpStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EvpnOverlay) {
		toSerialize["evpn_overlay"] = o.EvpnOverlay
	}
	if !IsNil(o.ForOverlay) {
		toSerialize["for_overlay"] = o.ForOverlay
	}
	if !IsNil(o.LocalAs) {
		toSerialize["local_as"] = o.LocalAs
	}
	if !IsNil(o.Mac) {
		toSerialize["mac"] = o.Mac
	}
	if !IsNil(o.Neighbor) {
		toSerialize["neighbor"] = o.Neighbor
	}
	if !IsNil(o.NeighborAs) {
		toSerialize["neighbor_as"] = o.NeighborAs
	}
	if !IsNil(o.NeighborMac) {
		toSerialize["neighbor_mac"] = o.NeighborMac
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.OrgId) {
		toSerialize["org_id"] = o.OrgId
	}
	if !IsNil(o.RxPkts) {
		toSerialize["rx_pkts"] = o.RxPkts
	}
	if !IsNil(o.RxRoutes) {
		toSerialize["rx_routes"] = o.RxRoutes
	}
	if !IsNil(o.SiteId) {
		toSerialize["site_id"] = o.SiteId
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TxPkts) {
		toSerialize["tx_pkts"] = o.TxPkts
	}
	if !IsNil(o.TxRoutes) {
		toSerialize["tx_routes"] = o.TxRoutes
	}
	if !IsNil(o.Up) {
		toSerialize["up"] = o.Up
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	if !IsNil(o.VrfName) {
		toSerialize["vrf_name"] = o.VrfName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *BgpStats) UnmarshalJSON(data []byte) (err error) {
	varBgpStats := _BgpStats{}

	err = json.Unmarshal(data, &varBgpStats)

	if err != nil {
		return err
	}

	*o = BgpStats(varBgpStats)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "evpn_overlay")
		delete(additionalProperties, "for_overlay")
		delete(additionalProperties, "local_as")
		delete(additionalProperties, "mac")
		delete(additionalProperties, "neighbor")
		delete(additionalProperties, "neighbor_as")
		delete(additionalProperties, "neighbor_mac")
		delete(additionalProperties, "node")
		delete(additionalProperties, "org_id")
		delete(additionalProperties, "rx_pkts")
		delete(additionalProperties, "rx_routes")
		delete(additionalProperties, "site_id")
		delete(additionalProperties, "state")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "tx_pkts")
		delete(additionalProperties, "tx_routes")
		delete(additionalProperties, "up")
		delete(additionalProperties, "uptime")
		delete(additionalProperties, "vrf_name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBgpStats struct {
	value *BgpStats
	isSet bool
}

func (v NullableBgpStats) Get() *BgpStats {
	return v.value
}

func (v *NullableBgpStats) Set(val *BgpStats) {
	v.value = val
	v.isSet = true
}

func (v NullableBgpStats) IsSet() bool {
	return v.isSet
}

func (v *NullableBgpStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBgpStats(val *BgpStats) *NullableBgpStats {
	return &NullableBgpStats{value: val, isSet: true}
}

func (v NullableBgpStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBgpStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


