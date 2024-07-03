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

// checks if the JseInventoryItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JseInventoryItem{}

// JseInventoryItem struct for JseInventoryItem
type JseInventoryItem struct {
	ContractEndTime *int32 `json:"contract_end_time,omitempty"`
	ContractId *string `json:"contract_id,omitempty"`
	ContractSku *string `json:"contract_sku,omitempty"`
	ContractStartTime *int32 `json:"contract_start_time,omitempty"`
	// Contract type (Maintenance / Subscription / Premium / Gov AdvCare / Gov TAC / High Sec / AdvCare / Gov Premium)
	ContractType *string `json:"contract_type,omitempty"`
	// po number associated with this SKU
	CustomerPo *string `json:"customer_po,omitempty"`
	// distributor name
	Distributor *string `json:"distributor,omitempty"`
	// end of life time
	EolTime *int32 `json:"eol_time,omitempty"`
	// end of support time
	EosTime *int32 `json:"eos_time,omitempty"`
	// address where the device is installed. It is a combination of address , region , country , zip
	InstalledAddress *string `json:"installed_address,omitempty"`
	// model of the install base inventory
	Model *string `json:"model,omitempty"`
	// order ID associated with this SKU
	OrderId *string `json:"order_id,omitempty"`
	// reseller name
	Reseller *string `json:"reseller,omitempty"`
	// serial Number of the inventory
	Serial *string `json:"serial,omitempty"`
	// Shipped date
	ShippedTime *float32 `json:"shipped_time,omitempty"`
	// serviceable device stock
	Sku *string `json:"sku,omitempty"`
	Type *JseInventoryItemType `json:"type,omitempty"`
	WarrantyEndTime *int32 `json:"warranty_end_time,omitempty"`
	WarrantyStartTime *int32 `json:"warranty_start_time,omitempty"`
	WarrantyType *string `json:"warranty_type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _JseInventoryItem JseInventoryItem

// NewJseInventoryItem instantiates a new JseInventoryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJseInventoryItem() *JseInventoryItem {
	this := JseInventoryItem{}
	return &this
}

// NewJseInventoryItemWithDefaults instantiates a new JseInventoryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJseInventoryItemWithDefaults() *JseInventoryItem {
	this := JseInventoryItem{}
	return &this
}

// GetContractEndTime returns the ContractEndTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetContractEndTime() int32 {
	if o == nil || IsNil(o.ContractEndTime) {
		var ret int32
		return ret
	}
	return *o.ContractEndTime
}

// GetContractEndTimeOk returns a tuple with the ContractEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetContractEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractEndTime) {
		return nil, false
	}
	return o.ContractEndTime, true
}

// HasContractEndTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasContractEndTime() bool {
	if o != nil && !IsNil(o.ContractEndTime) {
		return true
	}

	return false
}

// SetContractEndTime gets a reference to the given int32 and assigns it to the ContractEndTime field.
func (o *JseInventoryItem) SetContractEndTime(v int32) {
	o.ContractEndTime = &v
}

// GetContractId returns the ContractId field value if set, zero value otherwise.
func (o *JseInventoryItem) GetContractId() string {
	if o == nil || IsNil(o.ContractId) {
		var ret string
		return ret
	}
	return *o.ContractId
}

// GetContractIdOk returns a tuple with the ContractId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetContractIdOk() (*string, bool) {
	if o == nil || IsNil(o.ContractId) {
		return nil, false
	}
	return o.ContractId, true
}

// HasContractId returns a boolean if a field has been set.
func (o *JseInventoryItem) HasContractId() bool {
	if o != nil && !IsNil(o.ContractId) {
		return true
	}

	return false
}

// SetContractId gets a reference to the given string and assigns it to the ContractId field.
func (o *JseInventoryItem) SetContractId(v string) {
	o.ContractId = &v
}

// GetContractSku returns the ContractSku field value if set, zero value otherwise.
func (o *JseInventoryItem) GetContractSku() string {
	if o == nil || IsNil(o.ContractSku) {
		var ret string
		return ret
	}
	return *o.ContractSku
}

// GetContractSkuOk returns a tuple with the ContractSku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetContractSkuOk() (*string, bool) {
	if o == nil || IsNil(o.ContractSku) {
		return nil, false
	}
	return o.ContractSku, true
}

// HasContractSku returns a boolean if a field has been set.
func (o *JseInventoryItem) HasContractSku() bool {
	if o != nil && !IsNil(o.ContractSku) {
		return true
	}

	return false
}

// SetContractSku gets a reference to the given string and assigns it to the ContractSku field.
func (o *JseInventoryItem) SetContractSku(v string) {
	o.ContractSku = &v
}

// GetContractStartTime returns the ContractStartTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetContractStartTime() int32 {
	if o == nil || IsNil(o.ContractStartTime) {
		var ret int32
		return ret
	}
	return *o.ContractStartTime
}

// GetContractStartTimeOk returns a tuple with the ContractStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetContractStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ContractStartTime) {
		return nil, false
	}
	return o.ContractStartTime, true
}

// HasContractStartTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasContractStartTime() bool {
	if o != nil && !IsNil(o.ContractStartTime) {
		return true
	}

	return false
}

// SetContractStartTime gets a reference to the given int32 and assigns it to the ContractStartTime field.
func (o *JseInventoryItem) SetContractStartTime(v int32) {
	o.ContractStartTime = &v
}

// GetContractType returns the ContractType field value if set, zero value otherwise.
func (o *JseInventoryItem) GetContractType() string {
	if o == nil || IsNil(o.ContractType) {
		var ret string
		return ret
	}
	return *o.ContractType
}

// GetContractTypeOk returns a tuple with the ContractType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetContractTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ContractType) {
		return nil, false
	}
	return o.ContractType, true
}

// HasContractType returns a boolean if a field has been set.
func (o *JseInventoryItem) HasContractType() bool {
	if o != nil && !IsNil(o.ContractType) {
		return true
	}

	return false
}

// SetContractType gets a reference to the given string and assigns it to the ContractType field.
func (o *JseInventoryItem) SetContractType(v string) {
	o.ContractType = &v
}

// GetCustomerPo returns the CustomerPo field value if set, zero value otherwise.
func (o *JseInventoryItem) GetCustomerPo() string {
	if o == nil || IsNil(o.CustomerPo) {
		var ret string
		return ret
	}
	return *o.CustomerPo
}

// GetCustomerPoOk returns a tuple with the CustomerPo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetCustomerPoOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerPo) {
		return nil, false
	}
	return o.CustomerPo, true
}

// HasCustomerPo returns a boolean if a field has been set.
func (o *JseInventoryItem) HasCustomerPo() bool {
	if o != nil && !IsNil(o.CustomerPo) {
		return true
	}

	return false
}

// SetCustomerPo gets a reference to the given string and assigns it to the CustomerPo field.
func (o *JseInventoryItem) SetCustomerPo(v string) {
	o.CustomerPo = &v
}

// GetDistributor returns the Distributor field value if set, zero value otherwise.
func (o *JseInventoryItem) GetDistributor() string {
	if o == nil || IsNil(o.Distributor) {
		var ret string
		return ret
	}
	return *o.Distributor
}

// GetDistributorOk returns a tuple with the Distributor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetDistributorOk() (*string, bool) {
	if o == nil || IsNil(o.Distributor) {
		return nil, false
	}
	return o.Distributor, true
}

// HasDistributor returns a boolean if a field has been set.
func (o *JseInventoryItem) HasDistributor() bool {
	if o != nil && !IsNil(o.Distributor) {
		return true
	}

	return false
}

// SetDistributor gets a reference to the given string and assigns it to the Distributor field.
func (o *JseInventoryItem) SetDistributor(v string) {
	o.Distributor = &v
}

// GetEolTime returns the EolTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetEolTime() int32 {
	if o == nil || IsNil(o.EolTime) {
		var ret int32
		return ret
	}
	return *o.EolTime
}

// GetEolTimeOk returns a tuple with the EolTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetEolTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EolTime) {
		return nil, false
	}
	return o.EolTime, true
}

// HasEolTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasEolTime() bool {
	if o != nil && !IsNil(o.EolTime) {
		return true
	}

	return false
}

// SetEolTime gets a reference to the given int32 and assigns it to the EolTime field.
func (o *JseInventoryItem) SetEolTime(v int32) {
	o.EolTime = &v
}

// GetEosTime returns the EosTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetEosTime() int32 {
	if o == nil || IsNil(o.EosTime) {
		var ret int32
		return ret
	}
	return *o.EosTime
}

// GetEosTimeOk returns a tuple with the EosTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetEosTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.EosTime) {
		return nil, false
	}
	return o.EosTime, true
}

// HasEosTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasEosTime() bool {
	if o != nil && !IsNil(o.EosTime) {
		return true
	}

	return false
}

// SetEosTime gets a reference to the given int32 and assigns it to the EosTime field.
func (o *JseInventoryItem) SetEosTime(v int32) {
	o.EosTime = &v
}

// GetInstalledAddress returns the InstalledAddress field value if set, zero value otherwise.
func (o *JseInventoryItem) GetInstalledAddress() string {
	if o == nil || IsNil(o.InstalledAddress) {
		var ret string
		return ret
	}
	return *o.InstalledAddress
}

// GetInstalledAddressOk returns a tuple with the InstalledAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetInstalledAddressOk() (*string, bool) {
	if o == nil || IsNil(o.InstalledAddress) {
		return nil, false
	}
	return o.InstalledAddress, true
}

// HasInstalledAddress returns a boolean if a field has been set.
func (o *JseInventoryItem) HasInstalledAddress() bool {
	if o != nil && !IsNil(o.InstalledAddress) {
		return true
	}

	return false
}

// SetInstalledAddress gets a reference to the given string and assigns it to the InstalledAddress field.
func (o *JseInventoryItem) SetInstalledAddress(v string) {
	o.InstalledAddress = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *JseInventoryItem) GetModel() string {
	if o == nil || IsNil(o.Model) {
		var ret string
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetModelOk() (*string, bool) {
	if o == nil || IsNil(o.Model) {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *JseInventoryItem) HasModel() bool {
	if o != nil && !IsNil(o.Model) {
		return true
	}

	return false
}

// SetModel gets a reference to the given string and assigns it to the Model field.
func (o *JseInventoryItem) SetModel(v string) {
	o.Model = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *JseInventoryItem) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *JseInventoryItem) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *JseInventoryItem) SetOrderId(v string) {
	o.OrderId = &v
}

// GetReseller returns the Reseller field value if set, zero value otherwise.
func (o *JseInventoryItem) GetReseller() string {
	if o == nil || IsNil(o.Reseller) {
		var ret string
		return ret
	}
	return *o.Reseller
}

// GetResellerOk returns a tuple with the Reseller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetResellerOk() (*string, bool) {
	if o == nil || IsNil(o.Reseller) {
		return nil, false
	}
	return o.Reseller, true
}

// HasReseller returns a boolean if a field has been set.
func (o *JseInventoryItem) HasReseller() bool {
	if o != nil && !IsNil(o.Reseller) {
		return true
	}

	return false
}

// SetReseller gets a reference to the given string and assigns it to the Reseller field.
func (o *JseInventoryItem) SetReseller(v string) {
	o.Reseller = &v
}

// GetSerial returns the Serial field value if set, zero value otherwise.
func (o *JseInventoryItem) GetSerial() string {
	if o == nil || IsNil(o.Serial) {
		var ret string
		return ret
	}
	return *o.Serial
}

// GetSerialOk returns a tuple with the Serial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetSerialOk() (*string, bool) {
	if o == nil || IsNil(o.Serial) {
		return nil, false
	}
	return o.Serial, true
}

// HasSerial returns a boolean if a field has been set.
func (o *JseInventoryItem) HasSerial() bool {
	if o != nil && !IsNil(o.Serial) {
		return true
	}

	return false
}

// SetSerial gets a reference to the given string and assigns it to the Serial field.
func (o *JseInventoryItem) SetSerial(v string) {
	o.Serial = &v
}

// GetShippedTime returns the ShippedTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetShippedTime() float32 {
	if o == nil || IsNil(o.ShippedTime) {
		var ret float32
		return ret
	}
	return *o.ShippedTime
}

// GetShippedTimeOk returns a tuple with the ShippedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetShippedTimeOk() (*float32, bool) {
	if o == nil || IsNil(o.ShippedTime) {
		return nil, false
	}
	return o.ShippedTime, true
}

// HasShippedTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasShippedTime() bool {
	if o != nil && !IsNil(o.ShippedTime) {
		return true
	}

	return false
}

// SetShippedTime gets a reference to the given float32 and assigns it to the ShippedTime field.
func (o *JseInventoryItem) SetShippedTime(v float32) {
	o.ShippedTime = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *JseInventoryItem) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *JseInventoryItem) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *JseInventoryItem) SetSku(v string) {
	o.Sku = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *JseInventoryItem) GetType() JseInventoryItemType {
	if o == nil || IsNil(o.Type) {
		var ret JseInventoryItemType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetTypeOk() (*JseInventoryItemType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *JseInventoryItem) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given JseInventoryItemType and assigns it to the Type field.
func (o *JseInventoryItem) SetType(v JseInventoryItemType) {
	o.Type = &v
}

// GetWarrantyEndTime returns the WarrantyEndTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetWarrantyEndTime() int32 {
	if o == nil || IsNil(o.WarrantyEndTime) {
		var ret int32
		return ret
	}
	return *o.WarrantyEndTime
}

// GetWarrantyEndTimeOk returns a tuple with the WarrantyEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetWarrantyEndTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WarrantyEndTime) {
		return nil, false
	}
	return o.WarrantyEndTime, true
}

// HasWarrantyEndTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasWarrantyEndTime() bool {
	if o != nil && !IsNil(o.WarrantyEndTime) {
		return true
	}

	return false
}

// SetWarrantyEndTime gets a reference to the given int32 and assigns it to the WarrantyEndTime field.
func (o *JseInventoryItem) SetWarrantyEndTime(v int32) {
	o.WarrantyEndTime = &v
}

// GetWarrantyStartTime returns the WarrantyStartTime field value if set, zero value otherwise.
func (o *JseInventoryItem) GetWarrantyStartTime() int32 {
	if o == nil || IsNil(o.WarrantyStartTime) {
		var ret int32
		return ret
	}
	return *o.WarrantyStartTime
}

// GetWarrantyStartTimeOk returns a tuple with the WarrantyStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetWarrantyStartTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WarrantyStartTime) {
		return nil, false
	}
	return o.WarrantyStartTime, true
}

// HasWarrantyStartTime returns a boolean if a field has been set.
func (o *JseInventoryItem) HasWarrantyStartTime() bool {
	if o != nil && !IsNil(o.WarrantyStartTime) {
		return true
	}

	return false
}

// SetWarrantyStartTime gets a reference to the given int32 and assigns it to the WarrantyStartTime field.
func (o *JseInventoryItem) SetWarrantyStartTime(v int32) {
	o.WarrantyStartTime = &v
}

// GetWarrantyType returns the WarrantyType field value if set, zero value otherwise.
func (o *JseInventoryItem) GetWarrantyType() string {
	if o == nil || IsNil(o.WarrantyType) {
		var ret string
		return ret
	}
	return *o.WarrantyType
}

// GetWarrantyTypeOk returns a tuple with the WarrantyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JseInventoryItem) GetWarrantyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.WarrantyType) {
		return nil, false
	}
	return o.WarrantyType, true
}

// HasWarrantyType returns a boolean if a field has been set.
func (o *JseInventoryItem) HasWarrantyType() bool {
	if o != nil && !IsNil(o.WarrantyType) {
		return true
	}

	return false
}

// SetWarrantyType gets a reference to the given string and assigns it to the WarrantyType field.
func (o *JseInventoryItem) SetWarrantyType(v string) {
	o.WarrantyType = &v
}

func (o JseInventoryItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JseInventoryItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContractEndTime) {
		toSerialize["contract_end_time"] = o.ContractEndTime
	}
	if !IsNil(o.ContractId) {
		toSerialize["contract_id"] = o.ContractId
	}
	if !IsNil(o.ContractSku) {
		toSerialize["contract_sku"] = o.ContractSku
	}
	if !IsNil(o.ContractStartTime) {
		toSerialize["contract_start_time"] = o.ContractStartTime
	}
	if !IsNil(o.ContractType) {
		toSerialize["contract_type"] = o.ContractType
	}
	if !IsNil(o.CustomerPo) {
		toSerialize["customer_po"] = o.CustomerPo
	}
	if !IsNil(o.Distributor) {
		toSerialize["distributor"] = o.Distributor
	}
	if !IsNil(o.EolTime) {
		toSerialize["eol_time"] = o.EolTime
	}
	if !IsNil(o.EosTime) {
		toSerialize["eos_time"] = o.EosTime
	}
	if !IsNil(o.InstalledAddress) {
		toSerialize["installed_address"] = o.InstalledAddress
	}
	if !IsNil(o.Model) {
		toSerialize["model"] = o.Model
	}
	if !IsNil(o.OrderId) {
		toSerialize["order_id"] = o.OrderId
	}
	if !IsNil(o.Reseller) {
		toSerialize["reseller"] = o.Reseller
	}
	if !IsNil(o.Serial) {
		toSerialize["serial"] = o.Serial
	}
	if !IsNil(o.ShippedTime) {
		toSerialize["shipped_time"] = o.ShippedTime
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WarrantyEndTime) {
		toSerialize["warranty_end_time"] = o.WarrantyEndTime
	}
	if !IsNil(o.WarrantyStartTime) {
		toSerialize["warranty_start_time"] = o.WarrantyStartTime
	}
	if !IsNil(o.WarrantyType) {
		toSerialize["warranty_type"] = o.WarrantyType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *JseInventoryItem) UnmarshalJSON(data []byte) (err error) {
	varJseInventoryItem := _JseInventoryItem{}

	err = json.Unmarshal(data, &varJseInventoryItem)

	if err != nil {
		return err
	}

	*o = JseInventoryItem(varJseInventoryItem)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "contract_end_time")
		delete(additionalProperties, "contract_id")
		delete(additionalProperties, "contract_sku")
		delete(additionalProperties, "contract_start_time")
		delete(additionalProperties, "contract_type")
		delete(additionalProperties, "customer_po")
		delete(additionalProperties, "distributor")
		delete(additionalProperties, "eol_time")
		delete(additionalProperties, "eos_time")
		delete(additionalProperties, "installed_address")
		delete(additionalProperties, "model")
		delete(additionalProperties, "order_id")
		delete(additionalProperties, "reseller")
		delete(additionalProperties, "serial")
		delete(additionalProperties, "shipped_time")
		delete(additionalProperties, "sku")
		delete(additionalProperties, "type")
		delete(additionalProperties, "warranty_end_time")
		delete(additionalProperties, "warranty_start_time")
		delete(additionalProperties, "warranty_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableJseInventoryItem struct {
	value *JseInventoryItem
	isSet bool
}

func (v NullableJseInventoryItem) Get() *JseInventoryItem {
	return v.value
}

func (v *NullableJseInventoryItem) Set(val *JseInventoryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableJseInventoryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableJseInventoryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJseInventoryItem(val *JseInventoryItem) *NullableJseInventoryItem {
	return &NullableJseInventoryItem{value: val, isSet: true}
}

func (v NullableJseInventoryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJseInventoryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


