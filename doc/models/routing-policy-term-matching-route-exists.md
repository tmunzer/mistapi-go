
# Routing Policy Term Matching Route Exists

*This model accepts additional fields of type interface{}.*

## Structure

`RoutingPolicyTermMatchingRouteExists`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Route` | `*string` | Optional | - |
| `VrfName` | `*string` | Optional | Name of the vrf instance, it can also be the name of the VPN or wan if they<br><br>**Default**: `"default"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "route": "192.168.0.0/24",
  "vrf_name": "default",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

