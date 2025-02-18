
# Wlan Airwatch

Airwatch wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanAirwatch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiKey` | `*string` | Optional | API Key |
| `ConsoleUrl` | `*string` | Optional | Console URL |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `Password` | `*string` | Optional | Password |
| `Username` | `*string` | Optional | Username |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "api_key": "aHhlbGxvYXNkZmFzZGZhc2Rmc2RmCg==\"",
  "console_url": "https://hs1.airwatchportals.com",
  "enabled": false,
  "password": "user1",
  "username": "test123",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

