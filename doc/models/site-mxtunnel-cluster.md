
# Site Mxtunnel Cluster

## Structure

`SiteMxtunnelCluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | - |
| `TuntermHosts` | `[]string` | Optional | - |

## Example (as JSON)

```json
{
  "name": "primary",
  "tunterm_hosts": [
    "mxedge1",
    "mxedge2.local"
  ]
}
```

