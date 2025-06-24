
# Switch Auto Upgrade Container

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchAutoUpgradeContainer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SwitchAutoUpgrade`](../../doc/models/switch-auto-upgrade.md) | Optional | - |
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

