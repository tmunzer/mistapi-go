
# Utils Service Ping

## Structure

`UtilsServicePing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | **Default**: `10` |
| `Host` | `string` | Required | - |
| `Service` | `string` | Required | ping packet takes the same path as the service |
| `Size` | `*int` | Optional | **Default**: `56`<br>**Constraints**: `>= 56`, `<= 65535` |
| `Tenant` | `*string` | Optional | tenant context in which the packet is sent |

## Example (as JSON)

```json
{
  "count": 10,
  "host": "host6",
  "service": "service8",
  "size": 56,
  "tenant": "tenant8"
}
```

