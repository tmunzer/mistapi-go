
# Webhook Wifi Conn Raw

Sample of the `wifi-conn-raw` webhook payload.

This webhook topic provides raw data from packets emitted by connected devices.

Raw data webhooks are a special subset of webhooks that provide insight into raw data packets emitted by a client,
identified by their advertising MAC address (assets, discovered ble, connected wifi, unconnected wifi).  
The data that client raw data webhooks encompasses are reporting AP information, RSSI Data, and any special packets/telemetry
packets that the client may emit.

Note that client raw webhooks are the raw data coming from the client and do not contain the X,Y location data of the client.
In order to get the location data for a client please see our location webhooks.
Clients can be identified uniquely across these client raw data topics and location webhook topic using MAC address as the Unique identifier (client identifier).

## Structure

`WebhookWifiConnRaw`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Events` | [`[]models.WebhookWifiConnRawEvent`](../../doc/models/webhook-wifi-conn-raw-event.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `Topic` | `string` | Required, Constant | enum: `wifi-conn-raw`<br><br>**Value**: `"wifi-conn-raw"` |

## Example (as JSON)

```json
{
  "events": [
    {
      "map_id": "53f10664-3ce8-4c27-b382-0ef66432349f",
      "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
      "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
      "ap_id": "ap_id2",
      "ap_loc": [
        26.98,
        26.97,
        26.96
      ],
      "client_id": "client_id2",
      "connected_site": false,
      "extended_info_list": [
        {
          "frame_ctrl": 248,
          "payload": "payload2",
          "sequence_ctrl": 42
        },
        {
          "frame_ctrl": 248,
          "payload": "payload2",
          "sequence_ctrl": 42
        }
      ]
    }
  ],
  "topic": "wifi-conn-raw"
}
```

