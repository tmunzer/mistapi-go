package models

import (
    "encoding/json"
    "fmt"
)

// JseInventoryItem represents a JseInventoryItem struct.
type JseInventoryItem struct {
    ContractEndTime      *int                      `json:"contract_end_time,omitempty"`
    ContractId           *string                   `json:"contract_id,omitempty"`
    ContractSku          *string                   `json:"contract_sku,omitempty"`
    ContractStartTime    *int                      `json:"contract_start_time,omitempty"`
    // Contract type (Maintenance / Subscription / Premium / Gov AdvCare / Gov TAC / High Sec / AdvCare / Gov Premium)
    ContractType         *string                   `json:"contract_type,omitempty"`
    // PO number associated with this SKU
    CustomerPo           *string                   `json:"customer_po,omitempty"`
    // Distributor name
    Distributor          *string                   `json:"distributor,omitempty"`
    // End of life time
    EolTime              *int                      `json:"eol_time,omitempty"`
    // End of support time
    EosTime              *int                      `json:"eos_time,omitempty"`
    // Address where the device is installed. It is a combination of address , region , country , zip
    InstalledAddress     *string                   `json:"installed_address,omitempty"`
    // Model of the install base inventory
    Model                *string                   `json:"model,omitempty"`
    // Order ID associated with this SKU
    OrderId              *string                   `json:"order_id,omitempty"`
    // Reseller name
    Reseller             *string                   `json:"reseller,omitempty"`
    // Serial Number of the inventory
    Serial               *string                   `json:"serial,omitempty"`
    // Shipped date
    ShippedTime          *float64                  `json:"shipped_time,omitempty"`
    // Serviceable device stock
    Sku                  *string                   `json:"sku,omitempty"`
    // enum: `ap`, `gateway`, `switch`
    Type                 *JseInventoryItemTypeEnum `json:"type,omitempty"`
    WarrantyEndTime      *int                      `json:"warranty_end_time,omitempty"`
    WarrantyStartTime    *int                      `json:"warranty_start_time,omitempty"`
    WarrantyType         *string                   `json:"warranty_type,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for JseInventoryItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JseInventoryItem) String() string {
    return fmt.Sprintf(
    	"JseInventoryItem[ContractEndTime=%v, ContractId=%v, ContractSku=%v, ContractStartTime=%v, ContractType=%v, CustomerPo=%v, Distributor=%v, EolTime=%v, EosTime=%v, InstalledAddress=%v, Model=%v, OrderId=%v, Reseller=%v, Serial=%v, ShippedTime=%v, Sku=%v, Type=%v, WarrantyEndTime=%v, WarrantyStartTime=%v, WarrantyType=%v, AdditionalProperties=%v]",
    	j.ContractEndTime, j.ContractId, j.ContractSku, j.ContractStartTime, j.ContractType, j.CustomerPo, j.Distributor, j.EolTime, j.EosTime, j.InstalledAddress, j.Model, j.OrderId, j.Reseller, j.Serial, j.ShippedTime, j.Sku, j.Type, j.WarrantyEndTime, j.WarrantyStartTime, j.WarrantyType, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JseInventoryItem.
// It customizes the JSON marshaling process for JseInventoryItem objects.
func (j JseInventoryItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "contract_end_time", "contract_id", "contract_sku", "contract_start_time", "contract_type", "customer_po", "distributor", "eol_time", "eos_time", "installed_address", "model", "order_id", "reseller", "serial", "shipped_time", "sku", "type", "warranty_end_time", "warranty_start_time", "warranty_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JseInventoryItem object to a map representation for JSON marshaling.
func (j JseInventoryItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.ContractEndTime != nil {
        structMap["contract_end_time"] = j.ContractEndTime
    }
    if j.ContractId != nil {
        structMap["contract_id"] = j.ContractId
    }
    if j.ContractSku != nil {
        structMap["contract_sku"] = j.ContractSku
    }
    if j.ContractStartTime != nil {
        structMap["contract_start_time"] = j.ContractStartTime
    }
    if j.ContractType != nil {
        structMap["contract_type"] = j.ContractType
    }
    if j.CustomerPo != nil {
        structMap["customer_po"] = j.CustomerPo
    }
    if j.Distributor != nil {
        structMap["distributor"] = j.Distributor
    }
    if j.EolTime != nil {
        structMap["eol_time"] = j.EolTime
    }
    if j.EosTime != nil {
        structMap["eos_time"] = j.EosTime
    }
    if j.InstalledAddress != nil {
        structMap["installed_address"] = j.InstalledAddress
    }
    if j.Model != nil {
        structMap["model"] = j.Model
    }
    if j.OrderId != nil {
        structMap["order_id"] = j.OrderId
    }
    if j.Reseller != nil {
        structMap["reseller"] = j.Reseller
    }
    if j.Serial != nil {
        structMap["serial"] = j.Serial
    }
    if j.ShippedTime != nil {
        structMap["shipped_time"] = j.ShippedTime
    }
    if j.Sku != nil {
        structMap["sku"] = j.Sku
    }
    if j.Type != nil {
        structMap["type"] = j.Type
    }
    if j.WarrantyEndTime != nil {
        structMap["warranty_end_time"] = j.WarrantyEndTime
    }
    if j.WarrantyStartTime != nil {
        structMap["warranty_start_time"] = j.WarrantyStartTime
    }
    if j.WarrantyType != nil {
        structMap["warranty_type"] = j.WarrantyType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JseInventoryItem.
// It customizes the JSON unmarshaling process for JseInventoryItem objects.
func (j *JseInventoryItem) UnmarshalJSON(input []byte) error {
    var temp tempJseInventoryItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "contract_end_time", "contract_id", "contract_sku", "contract_start_time", "contract_type", "customer_po", "distributor", "eol_time", "eos_time", "installed_address", "model", "order_id", "reseller", "serial", "shipped_time", "sku", "type", "warranty_end_time", "warranty_start_time", "warranty_type")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.ContractEndTime = temp.ContractEndTime
    j.ContractId = temp.ContractId
    j.ContractSku = temp.ContractSku
    j.ContractStartTime = temp.ContractStartTime
    j.ContractType = temp.ContractType
    j.CustomerPo = temp.CustomerPo
    j.Distributor = temp.Distributor
    j.EolTime = temp.EolTime
    j.EosTime = temp.EosTime
    j.InstalledAddress = temp.InstalledAddress
    j.Model = temp.Model
    j.OrderId = temp.OrderId
    j.Reseller = temp.Reseller
    j.Serial = temp.Serial
    j.ShippedTime = temp.ShippedTime
    j.Sku = temp.Sku
    j.Type = temp.Type
    j.WarrantyEndTime = temp.WarrantyEndTime
    j.WarrantyStartTime = temp.WarrantyStartTime
    j.WarrantyType = temp.WarrantyType
    return nil
}

// tempJseInventoryItem is a temporary struct used for validating the fields of JseInventoryItem.
type tempJseInventoryItem  struct {
    ContractEndTime   *int                      `json:"contract_end_time,omitempty"`
    ContractId        *string                   `json:"contract_id,omitempty"`
    ContractSku       *string                   `json:"contract_sku,omitempty"`
    ContractStartTime *int                      `json:"contract_start_time,omitempty"`
    ContractType      *string                   `json:"contract_type,omitempty"`
    CustomerPo        *string                   `json:"customer_po,omitempty"`
    Distributor       *string                   `json:"distributor,omitempty"`
    EolTime           *int                      `json:"eol_time,omitempty"`
    EosTime           *int                      `json:"eos_time,omitempty"`
    InstalledAddress  *string                   `json:"installed_address,omitempty"`
    Model             *string                   `json:"model,omitempty"`
    OrderId           *string                   `json:"order_id,omitempty"`
    Reseller          *string                   `json:"reseller,omitempty"`
    Serial            *string                   `json:"serial,omitempty"`
    ShippedTime       *float64                  `json:"shipped_time,omitempty"`
    Sku               *string                   `json:"sku,omitempty"`
    Type              *JseInventoryItemTypeEnum `json:"type,omitempty"`
    WarrantyEndTime   *int                      `json:"warranty_end_time,omitempty"`
    WarrantyStartTime *int                      `json:"warranty_start_time,omitempty"`
    WarrantyType      *string                   `json:"warranty_type,omitempty"`
}
