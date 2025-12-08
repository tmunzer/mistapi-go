
# Mxedge Upgrade Multi Allow Downgrades

Whether downgrade is allowed when running version is higher than expected version for each service

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

## Example (as JSON)

```json
{
  "mxagent": false,
  "mxdas": false,
  "mxocproxy": false,
  "radsecproxy": false,
  "tunterm": false
}
```

