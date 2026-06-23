
# Utils Send Support Logs

Request to upload support files from a device

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsSendSupportLogs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Info` | [`*models.UtilsSendSupportLogsInfoEnum`](../../doc/models/utils-send-support-logs-info-enum.md) | Optional | Optional, enum:<br><br>* `code-dumps`: Upload all core dump files, if any found. Uploads for all members of VC on switches.<br>* `full`: Upload 1 file with output of `request support information`, 1 file that concatenates all `/var/log/outbound-ssh.log*` files, all core dump files, the 5 most recent `/var/log/messages*` files, and Mist agent logs<br>* `messages`: Upload 1 to 10 `/var/log/messages*` files<br>* `outbound-ssh`: Upload 1 file that concatenates all `/var/log/outbound-ssh.log*` files<br>* `process`: Upload 1 file with output of show ```system processes extensive`` * ```var-logs`: Upload all non-empty files in the `/var/log/` directory<br><br>**Default**: `"full"` |
| `Node` | `*string` | Optional | Optional for SSR: if node is not present, both nodes support files are uploaded |
| `NumMessagesFiles` | `*int` | Optional | Number of most recent messages files to upload when `info`==`messages`<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1`, `<= 10` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsSendSupportLogs := models.UtilsSendSupportLogs{
        Info:                 models.ToPointer(models.UtilsSendSupportLogsInfoEnum_FULL),
        Node:                 models.ToPointer("node0"),
        NumMessagesFiles:     models.ToPointer(1),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

