
# Stats Wireless Client Guest

information about this portal

## Structure

`StatsWirelessClientGuest`

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
  "authorized_expiring_time": 186.66,
  "authorized_time": 92.84,
  "company": "company4",
  "email": "email2",
  "field1": "field18",
  "name": "name4"
}
```
