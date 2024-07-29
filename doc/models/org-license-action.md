
# Org License Action

## Structure

`OrgLicenseAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmendmentId` | `*uuid.UUID` | Optional | if `op`==`unamend`, the ID of the operation to cancel |
| `DstOrgId` | `*uuid.UUID` | Optional | if `op`==`amend`, the id of the org where the license is moved |
| `Notes` | `*string` | Optional | if `op`==`annotate` |
| `Op` | [`models.OrgLicenseActionOperationEnum`](../../doc/models/org-license-action-operation-enum.md) | Required | to move a license, use the `amend` operation. enum: `amend`, `annotate`, `delete`, `unamend` |
| `Quantity` | `*int` | Optional | if `op`==`amend`, the number of licenses to move |
| `SubscriptionId` | `*string` | Optional | if `op`==`amend` or `op`==`delete`, the ID of the subscription to use |

## Example (as JSON)

```json
{
  "amendment_id": "0000018c-0000-0000-0000-000000000000",
  "dst_org_id": "00001982-0000-0000-0000-000000000000",
  "notes": "notes0",
  "op": "amend",
  "quantity": 90,
  "subscription_id": "subscription_id0"
}
```

