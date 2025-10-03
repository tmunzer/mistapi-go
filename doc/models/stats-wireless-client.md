
# Stats Wireless Client

*This model accepts additional fields of type interface{}.*

## Structure

`StatsWirelessClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Accuracy` | `*int` | Optional | Estimated client location accuracy, in meter |
| `AirespaceIfname` | `*string` | Optional | - |
| `Airwatch` | [`*models.StatsWirelessClientAirwatch`](../../doc/models/stats-wireless-client-airwatch.md) | Optional | Information if airwatch enabled |
| `ApId` | `uuid.UUID` | Required | AP ID the client is connected to |
| `ApMac` | `string` | Required | AP the client is connected to |
| `Band` | [`models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Required | enum: `24`, `5`, `6` |
| `Channel` | `int` | Required | Current channel |
| `DualBand` | `*bool` | Optional | Whether the client is dual_band capable (determined by whether we’ve seen probe requests from both bands) |
| `Family` | `*string` | Optional | Device family, through fingerprinting. iPod / Nexus Galaxy / Windows Mobile or CE … |
| `Guest` | [`*models.Guest`](../../doc/models/guest.md) | Optional | Guest |
| `Hostname` | `*string` | Optional | Hostname that we learned from sniffing DHCP |
| `IdleTime` | `*float64` | Optional | How long, in seconds, has the client been idle (since the last RX packet) |
| `Ip` | `*string` | Optional | - |
| `IsGuest` | `bool` | Required | Whether this is a guest<br><br>**Default**: `false` |
| `KeyMgmt` | `string` | Required | E.g. WPA2-PSK/CCMP |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `Mac` | `string` | Required | Client mac |
| `Manufacture` | `*string` | Optional | Device manufacture, through fingerprinting or OUI |
| `MapId` | `*uuid.UUID` | Optional | Estimated client location - map_id |
| `Model` | `*string` | Optional | Device model, may be available if we can identify them |
| `NumLocatingAps` | `*int` | Optional | Number of APs used to locate this client |
| `Os` | `*string` | Optional | Device os, through fingerprinting |
| `PowerSaving` | `*bool` | Optional | If it’s currently in power-save mode |
| `Proto` | [`models.Dot11ProtoEnum`](../../doc/models/dot-11-proto-enum.md) | Required | enum: `a`, `ac`, `ax`, `b`, `be`, `g`, `n` |
| `PskId` | `*uuid.UUID` | Optional | PSK id (if multi-psk is used) |
| `Rssi` | `float64` | Required | Signal strength |
| `Rssizones` | [`[]models.StatsWirelessClientRssiZone`](../../doc/models/stats-wireless-client-rssi-zone.md) | Optional | List of rssizone_id’s where client is in and since when (if known) |
| `RxBps` | `models.Optional[int64]` | Optional | Rate of receiving traffic, bits/seconds, last known |
| `RxBytes` | `models.Optional[int64]` | Optional | Amount of traffic received since connection |
| `RxPkts` | `models.Optional[int64]` | Optional | Amount of packets received since connection |
| `RxRate` | `models.Optional[float64]` | Optional | RX Rate, Mbps |
| `RxRetries` | `models.Optional[int]` | Optional | Amount of rx retries |
| `Snr` | `float64` | Required | Signal over noise |
| `Ssid` | `string` | Required | SSID the client is connected to |
| `TxBps` | `models.Optional[int64]` | Optional | Rate of transmitting traffic, bits/seconds, last known |
| `TxBytes` | `models.Optional[int64]` | Optional | Amount of traffic sent since connection |
| `TxPkts` | `models.Optional[int64]` | Optional | Amount of packets sent since connection |
| `TxRate` | `models.Optional[float64]` | Optional | TX Rate, Mbps |
| `TxRetries` | `models.Optional[int]` | Optional | Amount of tx retries |
| `Type` | `*string` | Optional | Client’s type, regular / vip / resource / blocked (if client object is created) |
| `Uptime` | `*float64` | Optional | How long, in seconds, has the client been connected |
| `Username` | `*string` | Optional | Username that we learned from 802.1X exchange or Per_user PSK or User Portal |
| `Vbeacons` | [`[]models.StatsWirelessClientVbeacon`](../../doc/models/stats-wireless-client-vbeacon.md) | Optional | List of beacon_id’s where the client is in and since when (if known) |
| `VlanId` | `*string` | Optional | VLAN id, could be empty (from older AP) |
| `WlanId` | `uuid.UUID` | Required | WLAN ID the client is connected to |
| `WxruleId` | `*uuid.UUID` | Optional | Current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched. |
| `WxruleUsage` | [`[]models.StatsWirelessClientWxruleUsage`](../../doc/models/stats-wireless-client-wxrule-usage.md) | Optional | Current WxlanRule usage per tag_id |
| `X` | `*float64` | Optional | Estimated client location in pixels |
| `XM` | `*float64` | Optional | Estimated client location in meter |
| `Y` | `*float64` | Optional | Estimated client location in pixels |
| `YM` | `*float64` | Optional | Estimated client location in meter |
| `Zones` | [`[]models.StatsWirelessClientZone`](../../doc/models/stats-wireless-client-zone.md) | Optional | List of zone_id’s where client is in and since when (if known) |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_id": "00001902-0000-0000-0000-000000000000",
  "ap_mac": "ap_mac2",
  "band": "5",
  "channel": 152,
  "is_guest": false,
  "key_mgmt": "key_mgmt2",
  "last_seen": 1470417522,
  "mac": "mac4",
  "proto": "b",
  "rssi": 66.02,
  "rx_bps": 60003,
  "rx_bytes": 8515104416,
  "rx_pkts": 57770567,
  "snr": 64.24,
  "ssid": "ssid8",
  "tx_bps": 634301,
  "tx_bytes": 211217389682,
  "tx_pkts": 812204062,
  "wlan_id": "000004a8-0000-0000-0000-000000000000",
  "accuracy": 154,
  "airespace_ifname": "airespace_ifname0",
  "airwatch": {
    "authorized": false,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "dual_band": false,
  "family": "family2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

