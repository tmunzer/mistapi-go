
# Test Sms Global

Request body for validating SMSGlobal SMS gateway credentials

*This model accepts additional fields of type interface{}.*

## Structure

`TestSmsGlobal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `SmsglobalApiKey` | `string` | Required | SMSGlobal API key used to send the test SMS |
| `SmsglobalApiSecret` | `string` | Required | SMSGlobal API secret used to send the test SMS |
| `SmsglobalSender` | `*string` | Optional | Optional sender's number or sender ID. If not provided, uses the default number associated with the account |
| `To` | `string` | Required | Phone number of the recipient of SMS with country code |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    testSmsGlobal := models.TestSmsGlobal{
        SmsglobalApiKey:      "123456",
        SmsglobalApiSecret:   "abcdef",
        SmsglobalSender:      models.ToPointer("61400000002"),
        To:                   "+911122334455",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

