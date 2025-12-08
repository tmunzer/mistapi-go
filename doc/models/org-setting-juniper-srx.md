
# Org Setting Juniper Srx

## Structure

`OrgSettingJuniperSrx`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.JuniperSrxAutoUpgrade`](../../doc/models/juniper-srx-auto-upgrade.md) | Optional | auto_upgrade device first time it is onboarded |

## Example (as JSON)

```json
{
  "auto_upgrade": {
    "custom_versions": {
      "key0": "custom_versions3",
      "key1": "custom_versions2"
    },
    "enabled": false,
    "snapshot": false
  }
}
```

