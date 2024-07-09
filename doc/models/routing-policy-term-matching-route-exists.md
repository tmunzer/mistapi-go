
# Routing Policy Term Matching Route Exists

## Structure

`RoutingPolicyTermMatchingRouteExists`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Route` | `*string` | Optional | - |
| `VrfName` | `*string` | Optional | name of the vrf instance<br>it can also be the name of the VPN or wan if they<br>**Default**: `"default"` |

## Example (as JSON)

```json
{
  "route": "192.168.0.0/24",
  "vrf_name": "default"
}
```

