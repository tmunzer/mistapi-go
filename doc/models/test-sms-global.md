
# Test Sms Global

*This model accepts additional fields of type interface{}.*

## Structure

`TestSmsGlobal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SmsglobalApiKey` | `string` | Required | SMSGlobal api key |
| `SmsglobalApiSecret` | `string` | Required | SMSGlobal api secret |
| `To` | `string` | Required | Phone number of the recipient of SMS with country code |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "smsglobal_api_key": "123456",
  "smsglobal_api_secret": "abcdef",
  "to": "+911122334455",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

