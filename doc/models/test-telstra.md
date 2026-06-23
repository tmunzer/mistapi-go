
# Test Telstra

Request body for validating Telstra SMS gateway credentials

*This model accepts additional fields of type interface{}.*

## Structure

`TestTelstra`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `TelstraClientId` | `string` | Required | Telstra client identifier used to send the test SMS |
| `TelstraClientSecret` | `string` | Required | Telstra client secret used to send the test SMS |
| `To` | `string` | Required | Phone number of the recipient of SMS with country code |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    testTelstra := models.TestTelstra{
        TelstraClientId:      "123456",
        TelstraClientSecret:  "abcdef",
        To:                   "+911122334455",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

