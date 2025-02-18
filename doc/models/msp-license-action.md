
# Msp License Action

*This model accepts additional fields of type interface{}.*

## Structure

`MspLicenseAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmendmentId` | `*string` | Optional | Required if `op`==`unamend` |
| `DstOrgId` | `*uuid.UUID` | Optional | Required if `op`==`amend`, destination org id |
| `Notes` | `*string` | Optional | Required if `op`==`annotate` |
| `Op` | [`models.MspLicenseActionOperationEnum`](../../doc/models/msp-license-action-operation-enum.md) | Required | enum: `amend`, `annotate`, `delete`, `unamend`<br>**Constraints**: *Minimum Length*: `1` |
| `Quantity` | `*float64` | Optional | Required if `op`==`amend` |
| `SubscriptionId` | `*string` | Optional | Required if `op`==`annotate`<br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "amendment_id": "amendment_id4",
  "dst_org_id": "00000976-0000-0000-0000-000000000000",
  "notes": "notes8",
  "op": "amend",
  "quantity": 46.94,
  "subscription_id": "subscription_id8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

