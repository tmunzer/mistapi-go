
# Upgrade Org Devices Item Site Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevicesItemSiteUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `UpgradeId` | `*uuid.UUID` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "upgrade_id": "ebbdbd0b-1bcf-4e55-8a6a-3416049a52b1",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

