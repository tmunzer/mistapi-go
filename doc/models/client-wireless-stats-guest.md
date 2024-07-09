
# Client Wireless Stats Guest

information about this portal

## Structure

`ClientWirelessStatsGuest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Authorized` | `bool` | Required | whether this guest is authorized<br>**Default**: `false` |
| `AuthorizedExpiringTime` | `float64` | Required | when the guest authorization will expire |
| `AuthorizedTime` | `float64` | Required | when the guest is authorized |
| `Company` | `string` | Required | - |
| `Email` | `string` | Required | - |
| `Field1` | `string` | Required | - |
| `Name` | `string` | Required | - |

## Example (as JSON)

```json
{
  "authorized": false,
  "authorized_expiring_time": 228.7,
  "authorized_time": 205.2,
  "company": "company8",
  "email": "email2",
  "field1": "field12",
  "name": "name8"
}
```

