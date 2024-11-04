
# Org Device Upgrade

## Structure

`OrgDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `SiteUpgrades` | [`[]models.OrgDeviceUpgradeSiteUpgrade`](../../doc/models/org-device-upgrade-site-upgrade.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "site_upgrades": [
    {
      "site_id": "00000ca4-0000-0000-0000-000000000000",
      "upgrade_id": "000010fe-0000-0000-0000-000000000000"
    },
    {
      "site_id": "00000ca4-0000-0000-0000-000000000000",
      "upgrade_id": "000010fe-0000-0000-0000-000000000000"
    }
  ]
}
```

