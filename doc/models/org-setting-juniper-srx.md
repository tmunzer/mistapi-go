
# Org Setting Juniper Srx

*This model accepts additional fields of type interface{}.*

## Structure

`OrgSettingJuniperSrx`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.JuniperSrxAutoUpgrade`](../../doc/models/juniper-srx-auto-upgrade.md) | Optional | auto_upgrade device first time it is onboarded |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auto_upgrade": {
    "custom_versions": {
      "key0": "custom_versions3",
      "key1": "custom_versions2"
    },
    "enabled": false,
    "snapshot": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

