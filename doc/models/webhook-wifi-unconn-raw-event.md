
# Webhook Wifi Unconn Raw Event

## Structure

`WebhookWifiUnconnRawEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApId` | `*string` | Optional | - |
| `ApLoc` | `[]float64` | Optional | optional, coordinates (if any) of reporting AP (updated once in 60s per client) |
| `ClientId` | `*string` | Optional | - |
| `ConnectedSite` | `*bool` | Optional | - |
| `MapId` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organization |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Packets` | [`[]models.WebhookWifiUnconnRawEventPacket`](../../doc/models/webhook-wifi-unconn-raw-event-packet.md) | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |

## Example (as JSON)

```json
{
  "map_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ap_id": "ap_id8",
  "ap_loc": [
    56.46,
    56.47
  ],
  "client_id": "client_id8",
  "connected_site": false
}
```

