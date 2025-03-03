
# Upgrade Org Devices Item

*This model accepts additional fields of type interface{}.*

## Structure

`UpgradeOrgDevicesItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `SiteUpgrades` | [`[]models.UpgradeOrgDevicesItemSiteUpgrade`](../../doc/models/upgrade-org-devices-item-site-upgrade.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "site_upgrades": [
    {
      "site_id": "00000ca4-0000-0000-0000-000000000000",
      "upgrade_id": "000010fe-0000-0000-0000-000000000000",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "site_id": "00000ca4-0000-0000-0000-000000000000",
      "upgrade_id": "000010fe-0000-0000-0000-000000000000",
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

