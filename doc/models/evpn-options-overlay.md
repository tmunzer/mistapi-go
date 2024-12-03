
# Evpn Options Overlay

*This model accepts additional fields of type interface{}.*

## Structure

`EvpnOptionsOverlay`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `As` | `*int` | Optional | Overlay BGP Local AS Number<br>**Default**: `65000`<br>**Constraints**: `>= 1`, `<= 65535` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "as": 65000,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

