
# Response Claim License

Result of claiming licenses or activation codes

## Structure

`ResponseClaimLicense`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InventoryAdded` | [`[]models.ResponseClaimLicenseInventoryItem`](../../doc/models/response-claim-license-inventory-item.md) | Required | Inventory devices added by the claim operation<br><br>**Constraints**: *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseClaimLicenseInventoryItem`](../../doc/models/response-claim-license-inventory-item.md) | Required | Inventory devices already present during the claim operation<br><br>**Constraints**: *Unique Items Required* |
| `InventoryPending` | [`[]models.ResponseClaimLicenseInventoryPendingItem`](../../doc/models/response-claim-license-inventory-pending-item.md) | Optional | Inventory devices pending asynchronous claim processing<br><br>**Constraints**: *Unique Items Required* |
| `LicenseAdded` | [`[]models.ResponseClaimLicenseLicenseItem`](../../doc/models/response-claim-license-license-item.md) | Required | License entitlements added by the claim operation<br><br>**Constraints**: *Unique Items Required* |
| `LicenseDuplicated` | [`[]models.ResponseClaimLicenseLicenseItem`](../../doc/models/response-claim-license-license-item.md) | Required | License entitlements already present during the claim operation<br><br>**Constraints**: *Unique Items Required* |
| `LicenseError` | [`[]models.ResponseClaimLicenseLicenseErrorItem`](../../doc/models/response-claim-license-license-error-item.md) | Required | License claim errors returned by order number<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseClaimLicense := models.ResponseClaimLicense{
        InventoryAdded:       []models.ResponseClaimLicenseInventoryItem{
            models.ResponseClaimLicenseInventoryItem{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        InventoryDuplicated:  []models.ResponseClaimLicenseInventoryItem{
            models.ResponseClaimLicenseInventoryItem{
                Mac:                  "mac0",
                Magic:                "magic6",
                Model:                "model4",
                Serial:               "serial6",
                Type:                 "type6",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        InventoryPending:     []models.ResponseClaimLicenseInventoryPendingItem{
            models.ResponseClaimLicenseInventoryPendingItem{
                Mac:                  models.ToPointer("mac6"),
            },
            models.ResponseClaimLicenseInventoryPendingItem{
                Mac:                  models.ToPointer("mac6"),
            },
        },
        LicenseAdded:         []models.ResponseClaimLicenseLicenseItem{
            models.ResponseClaimLicenseLicenseItem{
                End:                  62,
                Quantity:             132,
                Start:                20,
                Type:                 "type2",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        LicenseDuplicated:    []models.ResponseClaimLicenseLicenseItem{
            models.ResponseClaimLicenseLicenseItem{
                End:                  18,
                Quantity:             88,
                Start:                232,
                Type:                 "type8",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        LicenseError:         []models.ResponseClaimLicenseLicenseErrorItem{
            models.ResponseClaimLicenseLicenseErrorItem{
                Order:                "order2",
                Reason:               "reason0",
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
    }

}
```

