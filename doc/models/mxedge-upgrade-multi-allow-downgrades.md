
# Mxedge Upgrade Multi Allow Downgrades

Whether downgrade is allowed when running version is higher than expected version for each service

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeUpgradeMultiAllowDowngrades`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mxagent` | `*bool` | Optional | **Default**: `false` |
| `Mxdas` | `*bool` | Optional | **Default**: `false` |
| `Mxocproxy` | `*bool` | Optional | **Default**: `false` |
| `Radsecproxy` | `*bool` | Optional | **Default**: `false` |
| `Tunterm` | `*bool` | Optional | **Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mxagent": false,
  "mxdas": false,
  "mxocproxy": false,
  "radsecproxy": false,
  "tunterm": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

