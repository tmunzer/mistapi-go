
# Account Vmware Config

VMware account credentials and webhook settings for OAuth application linking

*This model accepts additional fields of type interface{}.*

## Structure

`AccountVmwareConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClientId` | `string` | Required | Customer account Client ID |
| `ClientSecret` | `string` | Required | Customer account Client Secret |
| `InstanceUrl` | `string` | Required | Customer account VMware instance URL |
| `WebhookEnabled` | `bool` | Required | Enables or disables the webhook integration |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    accountVmwareConfig := models.AccountVmwareConfig{
        ClientId:             "client_id0",
        ClientSecret:         "client_secret6",
        InstanceUrl:          "instance_url2",
        WebhookEnabled:       false,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

