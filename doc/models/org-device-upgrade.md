
# Org Device Upgrade

## Structure

`OrgDeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional | - |
| `SiteUpgrades` | [`[]models.OrgDeviceUpgradeSiteUpgrade`](../../doc/models/org-device-upgrade-site-upgrade.md) | Optional | - |

## Example (as JSON)

```json
{
  "id": "31223085-405d-4b64-8aea-9c5b98098b4b",
  "site_upgrades": [
    {
      "site_id": "00000ca4-0000-0000-0000-000000000000",
      "upgrade_id": "000010fe-0000-0000-0000-000000000000"
    }
  ]
}
```

