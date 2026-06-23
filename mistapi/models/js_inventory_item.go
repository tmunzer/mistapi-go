// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// JsInventoryItem represents a JsInventoryItem struct.
// Juniper Support inventory item with entitlement, lifecycle, and software metadata
type JsInventoryItem struct {
	// Current operational availability status of the device; only returned for onboarded (claimed) devices
	Availability *string `json:"availability,omitempty"`
	// Indicates if the device is claimed by any org
	Claimed *bool `json:"claimed,omitempty"`
	// Expiration date of the service contract; only returned for onboarded (claimed) devices
	ContractEndDate *string `json:"contract_end_date,omitempty"`
	// Name of the reseller associated with the contract; only returned for onboarded (claimed) devices
	ContractReseller *string `json:"contract_reseller,omitempty"`
	// Official commencement date of the service contract; only returned for onboarded (claimed) devices
	ContractStartDate *string `json:"contract_start_date,omitempty"`
	// General classification of the contract; only returned for onboarded (claimed) devices
	ContractType *string `json:"contract_type,omitempty"`
	// Current status of the contract (e.g., 'Active', 'Expired'); only returned for onboarded (claimed) devices
	CurrentContractFlag *string `json:"current_contract_flag,omitempty"`
	// Name of the device
	DeviceName *string `json:"device_name,omitempty"`
	// Name of the distributor who provided the device; only returned for onboarded (claimed) devices
	Distributor *string `json:"distributor,omitempty"`
	// End of sale epoch timestamp
	EndOfSaleTime *int `json:"end_of_sale_time,omitempty"`
	// Product support notice associated with the device end-of-life milestone
	EolPsn *string `json:"eol_psn,omitempty"`
	// End of support time
	EosTime *int `json:"eos_time,omitempty"`
	// Indicates if the device is covered under active support contract
	HasSupport *bool `json:"has_support,omitempty"`
	// Physical installation street address of the asset; only returned for onboarded (claimed) devices
	IaAddress *string `json:"ia_address,omitempty"`
	// Country of installation; only returned for onboarded (claimed) devices
	IaCountry *string `json:"ia_country,omitempty"`
	// Geographic region where the device is installed; only returned for onboarded (claimed) devices
	IaRegion *string `json:"ia_region,omitempty"`
	// ZIP or postal code of the installation site; only returned for onboarded (claimed) devices
	IaZipPostal *string `json:"ia_zip_postal,omitempty"`
	// Indicates whether it is Master
	Master *bool `json:"master,omitempty"`
	// Hardware model recorded in the install base inventory
	Model *string `json:"model,omitempty"`
	// Unique identifier of a Mist organization
	OrgId *uuid.UUID `json:"org_id,omitempty"`
	// Device serial number recorded in the install base inventory
	Serial *string `json:"serial,omitempty"`
	// Unique identifier for the service contract; only returned for onboarded (claimed) devices
	ServiceContractNo *string `json:"service_contract_no,omitempty"`
	// Specific service level or contract category for the device; only returned for onboarded (claimed) devices
	ServiceContractType *string `json:"service_contract_type,omitempty"`
	// Indicates if the service offer was declined (e.g., 'Y'/'N'); only returned for onboarded (claimed) devices
	ServiceDeclineFlag *string `json:"service_decline_flag,omitempty"`
	// Indicator of whether the device is eligible for service coverage ('Yes', 'No'); only returned for onboarded (claimed) devices
	ServiceEligible *string `json:"service_eligible,omitempty"`
	// Calculated or actual date the device was shipped; only returned for onboarded (claimed) devices
	ShipDateCalc *string `json:"ship_date_calc,omitempty"`
	// Serviceable SKU associated with the device
	Sku *string `json:"sku,omitempty"`
	// Current inventory status reported for the device
	Status *string `json:"status,omitempty"`
	// Recommended software version for the device
	SuggestedVersion *string `json:"suggested_version,omitempty"`
	// enum: `ap`, `gateway`, `switch`
	Type *DeviceTypeEnum `json:"type,omitempty"`
	// Software version currently running on the device
	Version *string `json:"version,omitempty"`
	// Release description associated with the reported software version
	VersionDescription *string `json:"version_description,omitempty"`
	// End of Service of version
	VersionEosTime *int `json:"version_eos_time,omitempty"`
	// FRS date of the version
	VersionTime *int `json:"version_time,omitempty"`
	// Coverage details associated with the device warranty
	Warranty *string `json:"warranty,omitempty"`
	// Expiration date of the warranty period; only returned for onboarded (claimed) devices
	WarrantyEnd *string `json:"warranty_end,omitempty"`
	// Start date of the manufacturer or secondary warranty; only returned for onboarded (claimed) devices
	WarrantyStart *string `json:"warranty_start,omitempty"`
	// Timestamp when warranty needs to be renewed
	WarrantyTime *int `json:"warranty_time,omitempty"`
	// Warranty type label for Juniper Support Insight (JSI) devices. enum: `Standard Hardware Warranty`, `Enhanced Hardware Warranty`, `Dead On Arrival Warranty`, `Limited Lifetime Warranty`, `Software Warranty`, `Limited Lifetime Warranty for WLA`, `Warranty-JCPO EOL (DOA Not Included)`, `MIST Enhanced Hardware Warranty`, `MIST Standard Warranty`, `Determine Lifetime warranty`
	WarrantyType         *JsiWarrantyTypeEnum   `json:"warranty_type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsInventoryItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsInventoryItem) String() string {
	return fmt.Sprintf(
		"JsInventoryItem[Availability=%v, Claimed=%v, ContractEndDate=%v, ContractReseller=%v, ContractStartDate=%v, ContractType=%v, CurrentContractFlag=%v, DeviceName=%v, Distributor=%v, EndOfSaleTime=%v, EolPsn=%v, EosTime=%v, HasSupport=%v, IaAddress=%v, IaCountry=%v, IaRegion=%v, IaZipPostal=%v, Master=%v, Model=%v, OrgId=%v, Serial=%v, ServiceContractNo=%v, ServiceContractType=%v, ServiceDeclineFlag=%v, ServiceEligible=%v, ShipDateCalc=%v, Sku=%v, Status=%v, SuggestedVersion=%v, Type=%v, Version=%v, VersionDescription=%v, VersionEosTime=%v, VersionTime=%v, Warranty=%v, WarrantyEnd=%v, WarrantyStart=%v, WarrantyTime=%v, WarrantyType=%v, AdditionalProperties=%v]",
		j.Availability, j.Claimed, j.ContractEndDate, j.ContractReseller, j.ContractStartDate, j.ContractType, j.CurrentContractFlag, j.DeviceName, j.Distributor, j.EndOfSaleTime, j.EolPsn, j.EosTime, j.HasSupport, j.IaAddress, j.IaCountry, j.IaRegion, j.IaZipPostal, j.Master, j.Model, j.OrgId, j.Serial, j.ServiceContractNo, j.ServiceContractType, j.ServiceDeclineFlag, j.ServiceEligible, j.ShipDateCalc, j.Sku, j.Status, j.SuggestedVersion, j.Type, j.Version, j.VersionDescription, j.VersionEosTime, j.VersionTime, j.Warranty, j.WarrantyEnd, j.WarrantyStart, j.WarrantyTime, j.WarrantyType, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsInventoryItem.
// It customizes the JSON marshaling process for JsInventoryItem objects.
func (j JsInventoryItem) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(j.AdditionalProperties,
		"availability", "claimed", "contract_end_date", "contract_reseller", "contract_start_date", "contract_type", "current_contract_flag", "device_name", "distributor", "end_of_sale_time", "eol_psn", "eos_time", "has_support", "ia_address", "ia_country", "ia_region", "ia_zip_postal", "master", "model", "org_id", "serial", "service_contract_no", "service_contract_type", "service_decline_flag", "service_eligible", "ship_date_calc", "sku", "status", "suggested_version", "type", "version", "version_description", "version_eos_time", "version_time", "warranty", "warranty_end", "warranty_start", "warranty_time", "warranty_type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(j.toMap())
}

// toMap converts the JsInventoryItem object to a map representation for JSON marshaling.
func (j JsInventoryItem) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, j.AdditionalProperties)
	if j.Availability != nil {
		structMap["availability"] = j.Availability
	}
	if j.Claimed != nil {
		structMap["claimed"] = j.Claimed
	}
	if j.ContractEndDate != nil {
		structMap["contract_end_date"] = j.ContractEndDate
	}
	if j.ContractReseller != nil {
		structMap["contract_reseller"] = j.ContractReseller
	}
	if j.ContractStartDate != nil {
		structMap["contract_start_date"] = j.ContractStartDate
	}
	if j.ContractType != nil {
		structMap["contract_type"] = j.ContractType
	}
	if j.CurrentContractFlag != nil {
		structMap["current_contract_flag"] = j.CurrentContractFlag
	}
	if j.DeviceName != nil {
		structMap["device_name"] = j.DeviceName
	}
	if j.Distributor != nil {
		structMap["distributor"] = j.Distributor
	}
	if j.EndOfSaleTime != nil {
		structMap["end_of_sale_time"] = j.EndOfSaleTime
	}
	if j.EolPsn != nil {
		structMap["eol_psn"] = j.EolPsn
	}
	if j.EosTime != nil {
		structMap["eos_time"] = j.EosTime
	}
	if j.HasSupport != nil {
		structMap["has_support"] = j.HasSupport
	}
	if j.IaAddress != nil {
		structMap["ia_address"] = j.IaAddress
	}
	if j.IaCountry != nil {
		structMap["ia_country"] = j.IaCountry
	}
	if j.IaRegion != nil {
		structMap["ia_region"] = j.IaRegion
	}
	if j.IaZipPostal != nil {
		structMap["ia_zip_postal"] = j.IaZipPostal
	}
	if j.Master != nil {
		structMap["master"] = j.Master
	}
	if j.Model != nil {
		structMap["model"] = j.Model
	}
	if j.OrgId != nil {
		structMap["org_id"] = j.OrgId
	}
	if j.Serial != nil {
		structMap["serial"] = j.Serial
	}
	if j.ServiceContractNo != nil {
		structMap["service_contract_no"] = j.ServiceContractNo
	}
	if j.ServiceContractType != nil {
		structMap["service_contract_type"] = j.ServiceContractType
	}
	if j.ServiceDeclineFlag != nil {
		structMap["service_decline_flag"] = j.ServiceDeclineFlag
	}
	if j.ServiceEligible != nil {
		structMap["service_eligible"] = j.ServiceEligible
	}
	if j.ShipDateCalc != nil {
		structMap["ship_date_calc"] = j.ShipDateCalc
	}
	if j.Sku != nil {
		structMap["sku"] = j.Sku
	}
	if j.Status != nil {
		structMap["status"] = j.Status
	}
	if j.SuggestedVersion != nil {
		structMap["suggested_version"] = j.SuggestedVersion
	}
	if j.Type != nil {
		structMap["type"] = j.Type
	}
	if j.Version != nil {
		structMap["version"] = j.Version
	}
	if j.VersionDescription != nil {
		structMap["version_description"] = j.VersionDescription
	}
	if j.VersionEosTime != nil {
		structMap["version_eos_time"] = j.VersionEosTime
	}
	if j.VersionTime != nil {
		structMap["version_time"] = j.VersionTime
	}
	if j.Warranty != nil {
		structMap["warranty"] = j.Warranty
	}
	if j.WarrantyEnd != nil {
		structMap["warranty_end"] = j.WarrantyEnd
	}
	if j.WarrantyStart != nil {
		structMap["warranty_start"] = j.WarrantyStart
	}
	if j.WarrantyTime != nil {
		structMap["warranty_time"] = j.WarrantyTime
	}
	if j.WarrantyType != nil {
		structMap["warranty_type"] = j.WarrantyType
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsInventoryItem.
// It customizes the JSON unmarshaling process for JsInventoryItem objects.
func (j *JsInventoryItem) UnmarshalJSON(input []byte) error {
	var temp tempJsInventoryItem
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "availability", "claimed", "contract_end_date", "contract_reseller", "contract_start_date", "contract_type", "current_contract_flag", "device_name", "distributor", "end_of_sale_time", "eol_psn", "eos_time", "has_support", "ia_address", "ia_country", "ia_region", "ia_zip_postal", "master", "model", "org_id", "serial", "service_contract_no", "service_contract_type", "service_decline_flag", "service_eligible", "ship_date_calc", "sku", "status", "suggested_version", "type", "version", "version_description", "version_eos_time", "version_time", "warranty", "warranty_end", "warranty_start", "warranty_time", "warranty_type")
	if err != nil {
		return err
	}
	j.AdditionalProperties = additionalProperties

	j.Availability = temp.Availability
	j.Claimed = temp.Claimed
	j.ContractEndDate = temp.ContractEndDate
	j.ContractReseller = temp.ContractReseller
	j.ContractStartDate = temp.ContractStartDate
	j.ContractType = temp.ContractType
	j.CurrentContractFlag = temp.CurrentContractFlag
	j.DeviceName = temp.DeviceName
	j.Distributor = temp.Distributor
	j.EndOfSaleTime = temp.EndOfSaleTime
	j.EolPsn = temp.EolPsn
	j.EosTime = temp.EosTime
	j.HasSupport = temp.HasSupport
	j.IaAddress = temp.IaAddress
	j.IaCountry = temp.IaCountry
	j.IaRegion = temp.IaRegion
	j.IaZipPostal = temp.IaZipPostal
	j.Master = temp.Master
	j.Model = temp.Model
	j.OrgId = temp.OrgId
	j.Serial = temp.Serial
	j.ServiceContractNo = temp.ServiceContractNo
	j.ServiceContractType = temp.ServiceContractType
	j.ServiceDeclineFlag = temp.ServiceDeclineFlag
	j.ServiceEligible = temp.ServiceEligible
	j.ShipDateCalc = temp.ShipDateCalc
	j.Sku = temp.Sku
	j.Status = temp.Status
	j.SuggestedVersion = temp.SuggestedVersion
	j.Type = temp.Type
	j.Version = temp.Version
	j.VersionDescription = temp.VersionDescription
	j.VersionEosTime = temp.VersionEosTime
	j.VersionTime = temp.VersionTime
	j.Warranty = temp.Warranty
	j.WarrantyEnd = temp.WarrantyEnd
	j.WarrantyStart = temp.WarrantyStart
	j.WarrantyTime = temp.WarrantyTime
	j.WarrantyType = temp.WarrantyType
	return nil
}

// tempJsInventoryItem is a temporary struct used for validating the fields of JsInventoryItem.
type tempJsInventoryItem struct {
	Availability        *string              `json:"availability,omitempty"`
	Claimed             *bool                `json:"claimed,omitempty"`
	ContractEndDate     *string              `json:"contract_end_date,omitempty"`
	ContractReseller    *string              `json:"contract_reseller,omitempty"`
	ContractStartDate   *string              `json:"contract_start_date,omitempty"`
	ContractType        *string              `json:"contract_type,omitempty"`
	CurrentContractFlag *string              `json:"current_contract_flag,omitempty"`
	DeviceName          *string              `json:"device_name,omitempty"`
	Distributor         *string              `json:"distributor,omitempty"`
	EndOfSaleTime       *int                 `json:"end_of_sale_time,omitempty"`
	EolPsn              *string              `json:"eol_psn,omitempty"`
	EosTime             *int                 `json:"eos_time,omitempty"`
	HasSupport          *bool                `json:"has_support,omitempty"`
	IaAddress           *string              `json:"ia_address,omitempty"`
	IaCountry           *string              `json:"ia_country,omitempty"`
	IaRegion            *string              `json:"ia_region,omitempty"`
	IaZipPostal         *string              `json:"ia_zip_postal,omitempty"`
	Master              *bool                `json:"master,omitempty"`
	Model               *string              `json:"model,omitempty"`
	OrgId               *uuid.UUID           `json:"org_id,omitempty"`
	Serial              *string              `json:"serial,omitempty"`
	ServiceContractNo   *string              `json:"service_contract_no,omitempty"`
	ServiceContractType *string              `json:"service_contract_type,omitempty"`
	ServiceDeclineFlag  *string              `json:"service_decline_flag,omitempty"`
	ServiceEligible     *string              `json:"service_eligible,omitempty"`
	ShipDateCalc        *string              `json:"ship_date_calc,omitempty"`
	Sku                 *string              `json:"sku,omitempty"`
	Status              *string              `json:"status,omitempty"`
	SuggestedVersion    *string              `json:"suggested_version,omitempty"`
	Type                *DeviceTypeEnum      `json:"type,omitempty"`
	Version             *string              `json:"version,omitempty"`
	VersionDescription  *string              `json:"version_description,omitempty"`
	VersionEosTime      *int                 `json:"version_eos_time,omitempty"`
	VersionTime         *int                 `json:"version_time,omitempty"`
	Warranty            *string              `json:"warranty,omitempty"`
	WarrantyEnd         *string              `json:"warranty_end,omitempty"`
	WarrantyStart       *string              `json:"warranty_start,omitempty"`
	WarrantyTime        *int                 `json:"warranty_time,omitempty"`
	WarrantyType        *JsiWarrantyTypeEnum `json:"warranty_type,omitempty"`
}
