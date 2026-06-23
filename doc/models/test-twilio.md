
# Test Twilio

Request body for validating Twilio SMS gateway credentials

*This model accepts additional fields of type interface{}.*

## Structure

`TestTwilio`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `From` | `string` | Required | One of the numbers you have in your Twilio account |
| `To` | `string` | Required | Phone number of the recipient of SMS |
| `TwilioAuthToken` | `string` | Required | Twilio Auth Token used to send the test SMS |
| `TwilioSid` | `string` | Required | Twilio account SID used to send the test SMS |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    testTwilio := models.TestTwilio{
        From:                 "+185051234567",
        To:                   "+19999999999",
        TwilioAuthToken:      "2135be04736a1a0a314bce432d61721a",
        TwilioSid:            "AC5f4366878d193fb4865ab151739999eb",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

