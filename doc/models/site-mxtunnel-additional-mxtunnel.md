
# Site Mxtunnel Additional Mxtunnel

## Structure

`SiteMxtunnelAdditionalMxtunnel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Clusters` | [`[]models.SiteMxtunnelCluster`](../../doc/models/site-mxtunnel-cluster.md) | Optional | for AP, how to connect to tunterm or radsecproxy |
| `HelloInterval` | `*int` | Optional | in seconds, used as heartbeat to detect if a tunnel is alive. AP will try another peer after missing N hellos specified by hello_retries<br>**Default**: `60`<br>**Constraints**: `>= 1`, `<= 300` |
| `HelloRetries` | `*int` | Optional | **Default**: `7`<br>**Constraints**: `>= 2`, `<= 30` |
| `Protocol` | [`*models.SiteMxtunnelProtocolEnum`](../../doc/models/site-mxtunnel-protocol-enum.md) | Optional | enum: `ip`, `udp` |
| `VlanIds` | `[]int` | Optional | - |

## Example (as JSON)

```json
{
  "hello_interval": 60,
  "hello_retries": 3,
  "protocol": "udp",
  "vlan_ids": [
    300,
    310,
    320
  ],
  "clusters": [
    {
      "name": "name6",
      "tunterm_hosts": [
        "tunterm_hosts6"
      ]
    },
    {
      "name": "name6",
      "tunterm_hosts": [
        "tunterm_hosts6"
      ]
    },
    {
      "name": "name6",
      "tunterm_hosts": [
        "tunterm_hosts6"
      ]
    }
  ]
}
```

