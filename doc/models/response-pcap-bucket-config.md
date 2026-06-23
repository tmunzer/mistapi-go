
# Response Pcap Bucket Config

Result of a custom packet capture bucket setup or verification operation

## Structure

`ResponsePcapBucketConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `*string` | Optional | Custom bucket name used for packet capture storage |
| `Detail` | `*string` | Optional | Status or error detail returned for the bucket operation |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responsePcapBucketConfig := models.ResponsePcapBucketConfig{
        Bucket:               models.ToPointer("bucket6"),
        Detail:               models.ToPointer("detail2"),
    }

}
```

