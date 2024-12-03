
# Mxedge Upgrade Version

version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nIgnored if distro upgrade, `tunterm`, `radsecproxy`, `mxagent`, `mxocproxy`, `mxdas` or `mxnacedge`

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeUpgradeVersion`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mxagent` | `string` | Required | **Default**: `"current"` |
| `Mxdas` | `*string` | Optional | **Default**: `"current"` |
| `Mxocproxy` | `*string` | Optional | **Default**: `"current"` |
| `Radsecproxy` | `*string` | Optional | **Default**: `"current"` |
| `Tunterm` | `string` | Required | **Default**: `"current"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "mxagent": "current",
  "mxdas": "current",
  "mxocproxy": "current",
  "radsecproxy": "current",
  "tunterm": "current",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

