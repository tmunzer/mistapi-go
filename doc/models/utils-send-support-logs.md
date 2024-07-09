
# Utils Send Support Logs

## Structure

`UtilsSendSupportLogs`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Info` | [`*models.UtilsSendSupportLogsInfoEnum`](../../doc/models/utils-send-support-logs-info-enum.md) | Optional | optional choice: process, outbound-ssh, and full (default)<br>**Default**: `"full"` |
| `Node` | `*string` | Optional | optional: for SSR, if node is not present, both nodes support files are uploaded |
| `NumMessagesFiles` | `*int` | Optional | optional: number of most recent messages files to upload.<br>**Default**: `1`<br>**Constraints**: `>= 1`, `<= 10` |

## Example (as JSON)

```json
{
  "info": "full",
  "num_messages_files": 1,
  "node": "node2"
}
```

