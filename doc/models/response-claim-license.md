
# Response Claim License

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClaimLicense`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InventoryAdded` | [`[]models.ResponseClaimLicenseInventoryItem`](../../doc/models/response-claim-license-inventory-item.md) | Required | **Constraints**: *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseClaimLicenseInventoryItem`](../../doc/models/response-claim-license-inventory-item.md) | Required | **Constraints**: *Unique Items Required* |
| `InventoryPending` | [`[]models.ResponseClaimLicenseInventoryPendingItem`](../../doc/models/response-claim-license-inventory-pending-item.md) | Optional | for async claim<br><br>**Constraints**: *Unique Items Required* |
| `LicenseAdded` | [`[]models.ResponseClaimLicenseLicenseItem`](../../doc/models/response-claim-license-license-item.md) | Required | **Constraints**: *Unique Items Required* |
| `LicenseDuplicated` | [`[]models.ResponseClaimLicenseLicenseItem`](../../doc/models/response-claim-license-license-item.md) | Required | **Constraints**: *Unique Items Required* |
| `LicenseError` | [`[]models.ResponseClaimLicenseLicenseErrorItem`](../../doc/models/response-claim-license-license-error-item.md) | Required | **Constraints**: *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "inventory_added": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "inventory_duplicated": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "license_added": [
    {
      "end": 62,
      "quantity": 132,
      "start": 20,
      "type": "type2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "license_duplicated": [
    {
      "end": 18,
      "quantity": 88,
      "start": 232,
      "type": "type8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "license_error": [
    {
      "order": "order2",
      "reason": "reason0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "inventory_pending": [
    {
      "mac": "mac6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "mac": "mac6",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

