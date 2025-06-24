
# Utils Send Support Logs

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsSendSupportLogs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Info` | [`*models.UtilsSendSupportLogsInfoEnum`](../../doc/models/utils-send-support-logs-info-enum.md) | Optional | Optional, enum:<br><br>* `code-dumps`: Upload all core dump files, if any found<br>* `full`: Upload 1 file with output of `request support information`, 1 file that concatenates all `/var/log/outbound-ssh.log*` files, all core dump files, the 5 most recent `/var/log/messages*` files, and Mist agent logs<br>* `messages`: Upload 1 to 10 `/var/log/messages*` files<br>* `outbound-ssh`: Upload 1 file that concatenates all `/var/log/outbound-ssh.log*` files<br>* `process`: Upload 1 file with output of show `system processes extensive` *``var-logs`: Upload all non-empty files in the`/var/log/` directory<br><br>**Default**: `"full"` |
| `Node` | `*string` | Optional | optional: for SSR, if node is not present, both nodes support files are uploaded |
| `NumMessagesFiles` | `*int` | Optional | optional: number of most recent messages files to upload.<br><br>**Default**: `1`<br><br>**Constraints**: `>= 1`, `<= 10` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "info": "full",
  "num_messages_files": 1,
  "node": "node6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

