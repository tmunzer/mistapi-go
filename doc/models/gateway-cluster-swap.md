
# Gateway Cluster Swap

## Structure

`GatewayClusterSwap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mac` | `*string` | Optional | when `op` ==`replacement_nodeX`, new node1<br>'s MAC, the device has to be standalone and assigned to the same site |
| `Op` | [`models.GatewayClusterSwapOpEnum`](../../doc/models/gateway-cluster-swap-op-enum.md) | Required | **Default**: `"swap"`<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "op": "swap",
  "mac": "mac0"
}
```

