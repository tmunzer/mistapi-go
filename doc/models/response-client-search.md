
# Response Client Search

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseClientSearch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `float64` | Required | - |
| `Limit` | `int` | Required | - |
| `Next` | `*string` | Optional | - |
| `Results` | [`[]models.ClientWireless`](../../doc/models/client-wireless.md) | Required | **Constraints**: *Unique Items Required* |
| `Start` | `float64` | Required | - |
| `Total` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "end": 84.88,
  "limit": 126,
  "results": [
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
      "last_device": "Zebra",
      "last_firmware": "wl0: Jan 20 2024 04:08:41 version 20.103.12.0.8.7.171 FWID 01-e09d2675",
      "last_hostname": "hostname-a",
      "last_ip": "10.100.0.157",
      "last_model": "MBP 16\\\" M1 2021",
      "last_os": "Sonoma",
      "last_os_version": "14.4.1 (Build 23E224)",
      "last_psk_id": "abf7dc5c-bb51-4bb7-93b6-5547400ffe11",
      "last_psk_name": "iot",
      "last_ssid": "john@mycorp.net",
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
      "username": [
        "user@corp.com"
      ],
      "vlan": [
        10
      ],
      "wlan_id": [
        "e5d67b07-aae8-494b-8584-cbc20c8110aa"
      ],
      "ftc": false,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "start": 40.94,
  "total": 32,
  "next": "next8",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

