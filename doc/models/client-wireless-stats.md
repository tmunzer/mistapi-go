
# Client Wireless Stats

## Structure

`ClientWirelessStats`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ttl` | `float64` | Required | TTL of the validity of the stat |
| `Accuracy` | `*int` | Optional | estimated client location accuracy, in meter |
| `AirespaceIfname` | `*string` | Optional | - |
| `Airwatch` | [`*models.ClientWirelessStatsAirwatch`](../../doc/models/client-wireless-stats-airwatch.md) | Optional | information if airwatch enabled |
| `ApId` | `uuid.UUID` | Required | AP ID the client is connected to |
| `ApMac` | `string` | Required | AP the client is connected to |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `6` |
| `Channel` | `int` | Required | current channel |
| `DualBand` | `bool` | Required | whether the client is dual_band capable (determined by whether we’ve seen probe requests from both bands) |
| `Family` | `string` | Required | device family, through fingerprinting. iPod / Nexus Galaxy / Windows Mobile or CE … |
| `Guest` | [`*models.ClientWirelessStatsGuest`](../../doc/models/client-wireless-stats-guest.md) | Optional | information about this portal |
| `Hostname` | `string` | Required | hostname that we learned from sniffing DHCP |
| `IdleTime` | `float64` | Required | how long, in seconds, has the client been idle (since the last RX packet) |
| `Ip` | `string` | Required | - |
| `IsGuest` | `bool` | Required | whether this is a guest<br>**Default**: `false` |
| `KeyMgmt` | `string` | Required | e.g. WPA2-PSK/CCMP |
| `LastSeen` | `float64` | Required | last seen timestamp |
| `Mac` | `string` | Required | client mac |
| `Manufacture` | `string` | Required | device manufacture, through fingerprinting or OUI |
| `MapId` | `*uuid.UUID` | Optional | estimated client location - map_id |
| `Model` | `string` | Required | device model, may be available if we can identify them |
| `NumLocatingAps` | `*int` | Optional | number of APs used to locate this client |
| `Os` | `string` | Required | device os, through fingerprinting |
| `PowerSaving` | `bool` | Required | if it’s currently in power-save mode |
| `Proto` | [`models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Required | enum: `a`, `ac`, `ax`, `b`, `g`, `n` |
| `PskId` | `*uuid.UUID` | Optional | PSK id (if multi-psk is used) |
| `Rssi` | `float64` | Required | signal strength |
| `Rssizones` | [`[]models.ClientWirelessStatsRssiZone`](../../doc/models/client-wireless-stats-rssi-zone.md) | Optional | list of rssizone_id’s where client is in and since when (if known) |
| `RxBps` | `float64` | Required | rate of receiving traffic from the clients, bits/seconds, last known |
| `RxBytes` | `float64` | Required | amount of traffic received from client since client connects |
| `RxPackets` | `float64` | Required | amount of traffic received from client since client connects |
| `RxRate` | `float64` | Required | RX Rate, Mbps |
| `RxRetries` | `float64` | Required | amount of rx retries |
| `Snr` | `float64` | Required | signal over noise |
| `Ssid` | `string` | Required | SSID the client is connected to |
| `TxBps` | `float64` | Required | rate of transmitting traffic to the clients, bits/seconds, last known |
| `TxBytes` | `float64` | Required | amount of traffic sent to client since client connects |
| `TxPackets` | `float64` | Required | amount of traffic sent to client since client connects |
| `TxRate` | `float64` | Required | TX Rate, Mbps |
| `TxRetries` | `float64` | Required | amount of tx retries |
| `Type` | `*string` | Optional | client’s type, regular / vip / resource / blocked (if client object is created) |
| `Uptime` | `float64` | Required | how long, in seconds, has the client been connected |
| `Username` | `string` | Required | username that we learned from 802.1X exchange or Per_user PSK or User Portal |
| `Vbeacons` | [`[]models.ClientWirelessStatsVbeacon`](../../doc/models/client-wireless-stats-vbeacon.md) | Optional | list of beacon_id’s where the client is in and since when (if known) |
| `VlanId` | `*int` | Optional | vlan id, could be empty (from older AP) |
| `WlanId` | `uuid.UUID` | Required | WLAN ID the client is connected to |
| `WxruleId` | `*uuid.UUID` | Optional | current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched. |
| `WxruleUsage` | [`[]models.ClientWirelessStatsWxruleUsage`](../../doc/models/client-wireless-stats-wxrule-usage.md) | Optional | current WxlanRule usage per tag_id |
| `X` | `*float64` | Optional | estimated clinet location in pixels |
| `XM` | `*float64` | Optional | estimated client location in meter |
| `Y` | `*float64` | Optional | estimated clinet location in pixels |
| `YM` | `*float64` | Optional | estimated client location in meter |
| `Zones` | [`[]models.ClientWirelessStatsZone`](../../doc/models/client-wireless-stats-zone.md) | Optional | list of zone_id’s where client is in and since when (if known) |

## Example (as JSON)

```json
{
  "_ttl": 141.98,
  "ap_id": "00000938-0000-0000-0000-000000000000",
  "ap_mac": "ap_mac0",
  "band": "24",
  "channel": 242,
  "dual_band": false,
  "family": "family0",
  "hostname": "hostname6",
  "idle_time": 80.9,
  "ip": "ip2",
  "is_guest": false,
  "key_mgmt": "key_mgmt0",
  "last_seen": 70.88,
  "mac": "mac2",
  "manufacture": "manufacture2",
  "model": "model4",
  "os": "os6",
  "power_saving": false,
  "proto": "ax",
  "rssi": 34.4,
  "rx_bps": 141.08,
  "rx_bytes": 36.76,
  "rx_packets": 88.94,
  "rx_rate": 29.82,
  "rx_retries": 252.6,
  "snr": 95.86,
  "ssid": "ssid4",
  "tx_bps": 113.22,
  "tx_bytes": 158.46,
  "tx_packets": 176.32,
  "tx_rate": 11.58,
  "tx_retries": 71.4,
  "uptime": 178.48,
  "username": "username2",
  "wlan_id": "0000097e-0000-0000-0000-000000000000",
  "accuracy": 64,
  "airespace_ifname": "airespace_ifname8",
  "airwatch": {
    "authorized": false
  },
  "guest": {
    "authorized": false,
    "authorized_expiring_time": 236.54,
    "authorized_time": 42.96,
    "company": "company2",
    "email": "email4",
    "field1": "field16",
    "name": "name2"
  },
  "map_id": "00001f46-0000-0000-0000-000000000000"
}
```

