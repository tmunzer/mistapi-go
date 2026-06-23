
# Sdk Invite Sms

Request body for sending an SDK invite by SMS

*This model accepts additional fields of type interface{}.*

## Structure

`SdkInviteSms`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Number` | `string` | Required | Destination phone number for the SDK invite SMS |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    sdkInviteSms := models.SdkInviteSms{
        Number:               "number4",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

