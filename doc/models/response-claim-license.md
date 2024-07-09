
# Response Claim License

## Structure

`ResponseClaimLicense`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InventoryAdded` | [`[]models.ResponseClaimLicenseInventoryItem`](../../doc/models/response-claim-license-inventory-item.md) | Required | **Constraints**: *Unique Items Required* |
| `InventoryDuplicated` | [`[]models.ResponseClaimLicenseInventoryItem`](../../doc/models/response-claim-license-inventory-item.md) | Required | **Constraints**: *Unique Items Required* |
| `LicenseAdded` | [`[]models.ResponseClaimLicenseLicenseItem`](../../doc/models/response-claim-license-license-item.md) | Required | **Constraints**: *Unique Items Required* |
| `LicenseDuplicated` | [`[]models.ResponseClaimLicenseLicenseItem`](../../doc/models/response-claim-license-license-item.md) | Required | **Constraints**: *Unique Items Required* |
| `LicenseError` | [`[]models.ResponseClaimLicenseLicenseErrorItem`](../../doc/models/response-claim-license-license-error-item.md) | Required | **Constraints**: *Unique Items Required* |

## Example (as JSON)

```json
{
  "inventory_added": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6"
    }
  ],
  "inventory_duplicated": [
    {
      "mac": "mac0",
      "magic": "magic6",
      "model": "model4",
      "serial": "serial6",
      "type": "type6"
    }
  ],
  "license_added": [
    {
      "end": 62,
      "quantity": 132,
      "start": 20,
      "type": "type2"
    }
  ],
  "license_duplicated": [
    {
      "end": 18,
      "quantity": 88,
      "start": 232,
      "type": "type8"
    }
  ],
  "license_error": [
    {
      "order": "order2",
      "reason": "reason0"
    }
  ]
}
```

