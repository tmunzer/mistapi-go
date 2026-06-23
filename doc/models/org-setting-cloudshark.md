
# Org Setting Cloudshark

Packet capture integration settings for CloudShark

## Structure

`OrgSettingCloudshark`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Apitoken` | `*string` | Optional | Token used by Mist to access the CloudShark integration |
| `Url` | `*string` | Optional | CloudShark Enterprise URL, if using a self-hosted CS Enterprise instance |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingCloudshark := models.OrgSettingCloudshark{
        Apitoken:             models.ToPointer("accbd6f10c6d05c3"),
        Url:                  models.ToPointer("https://cloudshark.hosted.domain"),
    }

}
```

