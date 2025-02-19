
# Wired Client Response

*This model accepts additional fields of type interface{}.*

## Structure

`WiredClientResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `[]string` | Optional | MAC Address of the switch the client is connected to |
| `DeviceMacPort` | [`[]models.WiredClientResponseDeviceMacPortItem`](../../doc/models/wired-client-response-device-mac-port-item.md) | Optional | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |
| `DhcpClientIdentifier` | `*string` | Optional | - |
| `DhcpClientOptions` | [`[]models.DhcpClientOption`](../../doc/models/dhcp-client-option.md) | Optional | - |
| `DhcpFqdn` | `*string` | Optional | - |
| `DhcpHostname` | `*string` | Optional | - |
| `DhcpRequestParams` | `*string` | Optional | - |
| `DhcpVendorClassIdentifier` | `*string` | Optional | - |
| `Ip` | `[]string` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `[]string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Vlan` | `[]int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "dhcp_client_identifier": "MAC address 00155df6d500",
  "dhcp_fqdn": "ITS-VMMT0-D1N02.mgthub.local",
  "dhcp_hostname": "ITS-VMMT0-D1N02",
  "dhcp_request_params": "1 3 6 15 31 33 43 44 46 47 119 121 249 252",
  "dhcp_vendor_class_identifier": "MSFT 5.0",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "device_mac": [
    "device_mac1"
  ],
  "device_mac_port": [
    {
      "device_mac": "device_mac8",
      "ip": "ip8",
      "port_id": "port_id4",
      "port_parent": "port_parent6",
      "start": "start8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "device_mac": "device_mac8",
      "ip": "ip8",
      "port_id": "port_id4",
      "port_parent": "port_parent6",
      "start": "start8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "dhcp_client_options": [
    {
      "code": "code2",
      "data": "data4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "code": "code2",
      "data": "data4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    {
      "code": "code2",
      "data": "data4",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

