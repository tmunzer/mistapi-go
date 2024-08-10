
# Msp License Action

## Structure

`MspLicenseAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmendmentId` | `*string` | Optional | required if `op`==`unamend` |
| `DstOrgId` | `*string` | Optional | required if `op`==`amend`, destination org id<br>**Constraints**: *Minimum Length*: `1` |
| `Notes` | `*string` | Optional | required if `op`== `annotate` |
| `Op` | [`models.MspLicenseActionOperationEnum`](../../doc/models/msp-license-action-operation-enum.md) | Required | enum: `amend`, `annotate`, `delete`, `unamend`<br>**Constraints**: *Minimum Length*: `1` |
| `Quantity` | `*float64` | Optional | required if `op`==`amend` |
| `SubscriptionId` | `*string` | Optional | required if `op`== `annotate`<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "amendment_id": "amendment_id4",
  "dst_org_id": "dst_org_id2",
  "notes": "notes8",
  "op": "amend",
  "quantity": 46.94,
  "subscription_id": "subscription_id8"
}
```

