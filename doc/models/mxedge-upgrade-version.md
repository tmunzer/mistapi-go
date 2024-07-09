
# Mxedge Upgrade Version

version to upgrade for each service, `current` / `latest` / `default` / specific version (e.g. `2.5.100`).\nignored if distro upgrade

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

## Example (as JSON)

```json
{
  "mxagent": "current",
  "mxdas": "current",
  "mxocproxy": "current",
  "radsecproxy": "current",
  "tunterm": "current"
}
```

