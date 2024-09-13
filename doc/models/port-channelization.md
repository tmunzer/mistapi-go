
# Port Channelization

## Structure

`PortChannelization`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalProperties` | `*string` | Optional | Property key is the interface name or range (e.g. `et-0/0/47`, `et-0/0/48-49`), Property value is the interface speed (e.g. `25g`, `50g`) |
| `Enabled` | `*bool` | Optional | **Default**: `false` |

## Example (as JSON)

```json
{
  "enabled": false,
  "additionalProperties": "additionalProperties2"
}
```

