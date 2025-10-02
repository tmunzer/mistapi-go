
# Juniper Srx Auto Upgrade

auto_upgrade device first time it is onboarded

*This model accepts additional fields of type interface{}.*

## Structure

`JuniperSrxAutoUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CustomVersions` | `map[string]string` | Optional | Property key is the SRX Hardware model (e.g. "SRX4600") |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Snapshot` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "snapshot": false,
  "custom_versions": {
    "key0": "custom_versions7",
    "key1": "custom_versions8",
    "key2": "custom_versions9"
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

