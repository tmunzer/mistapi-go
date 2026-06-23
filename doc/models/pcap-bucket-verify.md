
# Pcap Bucket Verify

Request to verify ownership of a custom packet capture bucket

*This model accepts additional fields of type interface{}.*

## Structure

`PcapBucketVerify`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bucket` | `string` | Required | Customer bucket name being verified for packet capture storage |
| `VerifyToken` | `string` | Required | Token read from the MIST_TOKEN file written during bucket setup |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pcapBucketVerify := models.PcapBucketVerify{
        Bucket:               "company-private-pcap",
        VerifyToken:          "eyJhbGciOiJIUzI1J9.eyJzdWIiOiIxMjM0joiMjgxOG5MDIyfQ.2rzcRvMA3Eg09NnjCAC-1EWMRtxAnFDM",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

