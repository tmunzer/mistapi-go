
# Org License Action

*This model accepts additional fields of type interface{}.*

## Structure

`OrgLicenseAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmendmentId` | `*uuid.UUID` | Optional | If `op`==`unamend`, the ID of the operation to cancel |
| `DstOrgId` | `*uuid.UUID` | Optional | If `op`==`amend`, the id of the org where the license is moved |
| `Notes` | `*string` | Optional | If `op`==`annotate` |
| `Op` | [`models.OrgLicenseActionOperationEnum`](../../doc/models/org-license-action-operation-enum.md) | Required | to move a license, use the `amend` operation. enum: `amend`, `annotate`, `delete`, `unamend` |
| `Quantity` | `*int` | Optional | If `op`==`amend`, the number of licenses to move |
| `SubscriptionId` | `*string` | Optional | If `op`==`amend` or `op`==`delete`, the ID of the subscription to use |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "amendment_id": "00000f82-0000-0000-0000-000000000000",
  "dst_org_id": "00000b8c-0000-0000-0000-000000000000",
  "notes": "notes4",
  "op": "delete",
  "quantity": 80,
  "subscription_id": "subscription_id4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

