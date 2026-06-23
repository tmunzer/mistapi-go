
# Pcap Bucket

Request to configure a custom packet capture bucket

*This model accepts additional fields of type interface{}.*

## Structure

`PcapBucket`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `string` | Required | Customer bucket name to use for packet capture files |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pcapBucket := models.PcapBucket{
        Bucket:               "company-private-pcap",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

