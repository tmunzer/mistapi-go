
# Module Stat Item Network Resource

## Structure

`ModuleStatItemNetworkResource`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | current usage of the network resource<br><br>**Constraints**: `>= 0` |
| `Limit` | `*int` | Optional | maximum usage of the network resource<br><br>**Constraints**: `>= 0` |
| `Type` | `*string` | Optional | type of the network resource (e.g. FIB, FLOW, ...) |

## Example (as JSON)

```json
{
  "count": 17,
  "limit": 768000,
  "type": "FIB"
}
```

