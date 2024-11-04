
# Client Wireless

## Structure

`ClientWireless`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `[]string` | Optional | List of AP MAC Addresses the client was connected to |
| `AppVersion` | `[]string` | Optional | only when client has the Marvis Client app running. List of the versions of the Marvis Client |
| `Band` | `*string` | Optional | Wi-Fi Radio band |
| `Device` | `[]string` | Optional | only when client has the Marvis Client app running. List of the type of device type detected |
| `Ftc` | `*bool` | Optional | - |
| `Hardware` | `*string` | Optional | only when client has the Marvis Client app running. Type of Wi-Fi adapter |
| `Hostname` | `[]string` | Optional | list of hostname detected for this client |
| `Ip` | `[]string` | Optional | list if the ip addresses detected for this client |
| `LastAp` | `*string` | Optional | latest AP where the client is/was connected to |
| `LastDevice` | `*string` | Optional | latest type of device we identified (e.g. iPhone, Mac, …) |
| `LastFirmware` | `*string` | Optional | only when client has the Marvis Client app running. Same as "firmware" |
| `LastHostname` | `*string` | Optional | latest hostname we detected for the client |
| `LastIp` | `*string` | Optional | latest ip address we detected for the client |
| `LastModel` | `*string` | Optional | only when client has the Marvis Client app running. latest client hardware model we detected for the client |
| `LastOs` | `*string` | Optional | only when client has the Marvis Client app running. Latest version of OS Type we detected for the client |
| `LastOsVersion` | `*string` | Optional | only when client has the Marvis Client app running. Latest version of OS Version we detected for the client |
| `LastPskId` | `*uuid.UUID` | Optional | only for PPSK authentication. Latest PPSK ID used by the client |
| `LastPskName` | `*string` | Optional | only for PPSK authentication. Latest PPSK Name used by the client |
| `LastSsid` | `*string` | Optional | Name of the latest SSID (WLAN) the client is/was connected to client |
| `LastUsername` | `*string` | Optional | only for 802.1X authentication. Latest username used by the client |
| `LastVlan` | `*int` | Optional | latest VLAN ID assigned to the client |
| `LastWlanId` | `*uuid.UUID` | Optional | ID of the latest SSID (WLAN) the client is/was connected to |
| `Mac` | `*string` | Optional | Client MAC Address |
| `Mfg` | `*string` | Optional | manufacturer of the client hardware (MAC OUI based) |
| `Model` | `*string` | Optional | only when client has the Marvis Client app running. Client hardware model |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `Os` | `[]string` | Optional | only when client is having the Marvis Client app running. List of OS detected for the client |
| `OsVersion` | `[]string` | Optional | only when client is having the Marvis Client app running. List of OS version detected for the client |
| `Protocol` | `*string` | Optional | 802.11 amendment |
| `PskId` | `[]uuid.UUID` | Optional | list of IDs of the PPSK used by the client |
| `PskName` | `[]string` | Optional | list of names of the PPSK used by the client |
| `RandomMac` | `*bool` | Optional | whether the client is using randomized MAC Address or not |
| `SdkVersion` | `[]string` | Optional | only when client has the Marvis Client app running. List of Marvis Client SDK version detected for the client |
| `SiteId` | `*uuid.UUID` | Optional | Mist Site ID where the client is connected |
| `SiteIds` | `[]uuid.UUID` | Optional | list of Mist Site IDs where the client was connected |
| `Ssid` | `[]string` | Optional | list of the WLAN names the client was connected to |
| `Timestamp` | `*float64` | Optional | when the data has been updated |
| `Username` | `[]string` | Optional | only for 802.1X authentication. List of usernames used by the client |
| `Vlan` | `[]int` | Optional | list of vlans that have been assigned to the client |
| `WlanId` | `[]uuid.UUID` | Optional | list of IDs of WLANs the client was connected to |

## Example (as JSON)

```json
{
  "ap": [
    "a83a79a947ee",
    "003e73170b4c"
  ],
  "app_version": [
    "0.100.3"
  ],
  "band": "5",
  "device": [
    "Mac"
  ],
  "hardware": "Apple Wi-Fi adapter",
  "hostname": [
    "hostname-a",
    "hostname-b"
  ],
  "ip": [
    "10.5.23.43",
    "192.168.0.2"
  ],
  "last_ap": "a83a79a947ee",
  "last_device": "Mac",
  "last_firmware": "wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675",
  "last_hostname": "hostname-a",
  "last_ip": "10.5.23.43",
  "last_model": "MBP 16\\\" M1 2021",
  "last_os": "Sonoma",
  "last_os_version": "14.4.1 (Build 23E224)",
  "last_psk_id": "abf7dc5c-bb51-4bb7-93b6-5547400ffe11",
  "last_psk_name": "iot",
  "last_ssid": "IoT SSID",
  "last_username": "user@corp.com",
  "last_vlan": 10,
  "last_wlan_id": "e5d67b07-aae8-494b-8584-cbc20c8110aa",
  "mac": "bcd074000000",
  "mfg": "Apple",
  "model": "MBP 16\\\" M1 2021",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "os": [
    "Sonoma"
  ],
  "os_version": [
    "14.4.1 (Build 23E224)"
  ],
  "protocol": "ax",
  "psk_id": [
    "abf7dc5c-bb51-4bb7-93b6-5547400ffe11"
  ],
  "psk_name": [
    "iot"
  ],
  "sdk_version": [
    "0.100.3"
  ],
  "site_id": "25ff5219-9be7-4db9-907d-0c9b60445147",
  "site_ids": [
    "25ff5219-9be7-4db9-907d-0c9b60445147"
  ],
  "ssid": [
    "IoT SSID"
  ],
  "timestamp": 1714124722.113,
  "username": [
    "user@corp.com"
  ],
  "vlan": [
    10
  ],
  "wlan_id": [
    "e5d67b07-aae8-494b-8584-cbc20c8110aa"
  ],
  "ftc": false
}
```

