
# Capture Scan Aps

Property key is the AP MAC address (e.g. "5c5b35000001"). All optionals, parent parameters will be used if not defined

## Structure

`CaptureScanAps`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Band` | [`*models.CaptureScanApsBandEnum`](../../doc/models/capture-scan-aps-band-enum.md) | Optional | Only Single value allowed. enum: `24`, `5`, `6`<br><br>**Default**: `"24"` |
| `Channel` | `*string` | Optional | Specify the channel value where scan PCAP has to be started |
| `TcpdumpExpression` | `*string` | Optional | tcpdump expression, port specific if specified under ports dict, otherwise applicable across ports if specified at top level of payload. Port specific value overrides top level value when both exist. |
| `Width` | `*string` | Optional | Specify the bandwidth value with respect to the channel. |

## Example (as JSON)

```json
{
  "band": "24",
  "channel": "channel2",
  "tcpdump_expression": "tcpdump_expression2",
  "width": "width6"
}
```

