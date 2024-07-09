
# Test Telstra

## Structure

`TestTelstra`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TelstraClientId` | `string` | Required | Telstra client id |
| `TelstraClientSecret` | `string` | Required | Telstra client secret |
| `To` | `string` | Required | Phone number of the recipient of SMS with country code |

## Example (as JSON)

```json
{
  "telstra_client_id": "123456",
  "telstra_client_secret": "abcdef",
  "to": "+911122334455"
}
```

