
# Account Sentinelone Config

OAuth linked CrowdStrike apps account details

*This model accepts additional fields of type interface{}.*

## Structure

`AccountSentineloneConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiToken` | `string` | Required | Access token used to authenticate to the SentinelOne account |
| `InstanceUrl` | `string` | Required | Customer account SentinelOne instance URL |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountSentineloneConfig := models.AccountSentineloneConfig{
        ApiToken:             "api_token8",
        InstanceUrl:          "instance_url8",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

