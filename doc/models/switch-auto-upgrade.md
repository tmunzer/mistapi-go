
# Switch Auto Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Custom version to be used. The Property Key is the switch hardware and the property value is the firmware version |
| `Enabled` | `*bool` | Optional | Enable auto upgrade for the switch |
| `Snapshot` | `*bool` | Optional | Enable snapshot during the upgrade process<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "custom_versions": {
    "QFX5120-32C": "23.4R2-S2.1",
    "QFX5130-32CD": "23.4R2-S2.3"
  },
  "snapshot": false,
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

