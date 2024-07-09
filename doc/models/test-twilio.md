
# Test Twilio

## Structure

`TestTwilio`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `From` | `string` | Required | One of the numbers you have in your Twilio account |
| `To` | `string` | Required | Phone number of the recipient of SMS |
| `TwilioAuthToken` | `string` | Required | Auth Token associated with twilio account |
| `TwilioSid` | `string` | Required | Twilio Account SID |

## Example (as JSON)

```json
{
  "from": "+185051234567",
  "to": "+19999999999",
  "twilio_auth_token": "2135be04736a1a0a314bce432d61721a",
  "twilio_sid": "AC5f4366878d193fb4865ab151739999eb"
}
```

