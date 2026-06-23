
# Js Inventory Item

Juniper Support inventory item with entitlement, lifecycle, and software metadata

## Structure

`JsInventoryItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Availability` | `*string` | Optional | Current operational availability status of the device; only returned for onboarded (claimed) devices |
| `Claimed` | `*bool` | Optional | Indicates if the device is claimed by any org |
| `ContractEndDate` | `*string` | Optional | Expiration date of the service contract; only returned for onboarded (claimed) devices |
| `ContractReseller` | `*string` | Optional | Name of the reseller associated with the contract; only returned for onboarded (claimed) devices |
| `ContractStartDate` | `*string` | Optional | Official commencement date of the service contract; only returned for onboarded (claimed) devices |
| `ContractType` | `*string` | Optional | General classification of the contract; only returned for onboarded (claimed) devices |
| `CurrentContractFlag` | `*string` | Optional | Current status of the contract (e.g., 'Active', 'Expired'); only returned for onboarded (claimed) devices |
| `DeviceName` | `*string` | Optional | Name of the device |
| `Distributor` | `*string` | Optional | Name of the distributor who provided the device; only returned for onboarded (claimed) devices |
| `EndOfSaleTime` | `*int` | Optional | End of sale epoch timestamp |
| `EolPsn` | `*string` | Optional | Product support notice associated with the device end-of-life milestone |
| `EosTime` | `*int` | Optional | End of support time |
| `HasSupport` | `*bool` | Optional | Indicates if the device is covered under active support contract |
| `IaAddress` | `*string` | Optional | Physical installation street address of the asset; only returned for onboarded (claimed) devices |
| `IaCountry` | `*string` | Optional | Country of installation; only returned for onboarded (claimed) devices |
| `IaRegion` | `*string` | Optional | Geographic region where the device is installed; only returned for onboarded (claimed) devices |
| `IaZipPostal` | `*string` | Optional | ZIP or postal code of the installation site; only returned for onboarded (claimed) devices |
| `Master` | `*bool` | Optional | Indicates whether it is Master |
| `Model` | `*string` | Optional | Hardware model recorded in the install base inventory |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Serial` | `*string` | Optional | Device serial number recorded in the install base inventory |
| `ServiceContractNo` | `*string` | Optional | Unique identifier for the service contract; only returned for onboarded (claimed) devices |
| `ServiceContractType` | `*string` | Optional | Specific service level or contract category for the device; only returned for onboarded (claimed) devices |
| `ServiceDeclineFlag` | `*string` | Optional | Indicates if the service offer was declined (e.g., 'Y'/'N'); only returned for onboarded (claimed) devices |
| `ServiceEligible` | `*string` | Optional | Indicator of whether the device is eligible for service coverage ('Yes', 'No'); only returned for onboarded (claimed) devices |
| `ShipDateCalc` | `*string` | Optional | Calculated or actual date the device was shipped; only returned for onboarded (claimed) devices |
| `Sku` | `*string` | Optional | Serviceable SKU associated with the device |
| `Status` | `*string` | Optional | Current inventory status reported for the device |
| `SuggestedVersion` | `*string` | Optional | Recommended software version for the device |
| `Type` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |
| `Version` | `*string` | Optional | Software version currently running on the device |
| `VersionDescription` | `*string` | Optional | Release description associated with the reported software version |
| `VersionEosTime` | `*int` | Optional | End of Service of version |
| `VersionTime` | `*int` | Optional | FRS date of the version |
| `Warranty` | `*string` | Optional | Coverage details associated with the device warranty |
| `WarrantyEnd` | `*string` | Optional | Expiration date of the warranty period; only returned for onboarded (claimed) devices |
| `WarrantyStart` | `*string` | Optional | Start date of the manufacturer or secondary warranty; only returned for onboarded (claimed) devices |
| `WarrantyTime` | `*int` | Optional | Timestamp when warranty needs to be renewed |
| `WarrantyType` | [`*models.JsiWarrantyTypeEnum`](../../doc/models/jsi-warranty-type-enum.md) | Optional | Warranty type label for Juniper Support Insight (JSI) devices. enum: `Standard Hardware Warranty`, `Enhanced Hardware Warranty`, `Dead On Arrival Warranty`, `Limited Lifetime Warranty`, `Software Warranty`, `Limited Lifetime Warranty for WLA`, `Warranty-JCPO EOL (DOA Not Included)`, `MIST Enhanced Hardware Warranty`, `MIST Standard Warranty`, `Determine Lifetime warranty` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    jsInventoryItem := models.JsInventoryItem{
        Availability:         models.ToPointer("availability4"),
        Claimed:              models.ToPointer(false),
        ContractEndDate:      models.ToPointer("contract_end_date4"),
        ContractReseller:     models.ToPointer("contract_reseller6"),
        ContractStartDate:    models.ToPointer("contract_start_date4"),
        EolPsn:               models.ToPointer("TSB18097"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
    }

}
```

