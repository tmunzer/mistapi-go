
# Mxedge Das

Configure cloud-assisted dynamic authorization service on this cluster of mist edges

## Structure

`MxedgeDas`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaServers` | [`[]models.MxedgeDasCoaServer`](../../doc/models/mxedge-das-coa-server.md) | Optional | Dynamic authorization clients configured to send CoA\|DM to mist edges on port 3799 |
| `Enabled` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false,
  "coa_servers": [
    {
      "disable_event_timestamp_check": false,
      "enabled": false,
      "host": "host8",
      "port": 28,
      "require_message_authenticator": false
    }
  ]
}
```

