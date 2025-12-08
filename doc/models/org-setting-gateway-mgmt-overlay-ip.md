
# Org Setting Gateway Mgmt Overlay Ip

## Structure

`OrgSettingGatewayMgmtOverlayIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | When it's going overlay, a routable IP to overlay will be required |
| `Node1Ip` | `*string` | Optional | For SSR HA cluster, another IP for node1 will be required, too |

## Example (as JSON)

```json
{
  "ip": "ip2",
  "node1_ip": "node1_ip0"
}
```

