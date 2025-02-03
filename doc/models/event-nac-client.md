
# Event Nac Client

*This model accepts additional fields of type interface{}.*

## Structure

`EventNacClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `AuthType` | `*string` | Optional | Authentication type, e.g. "eap-tls", "peap-tls", "eap-ttls", "eap-teap", "mab", "psk", "device-auth" |
| `Bssid` | `*string` | Optional | - |
| `DeviceMac` | `*string` | Optional | - |
| `DryrunNacruleId` | `*uuid.UUID` | Optional | NAC Policy Dry Run Rule ID, if present and matched |
| `DryrunNacruleMatched` | `*bool` | Optional | - |
| `IdpId` | `*uuid.UUID` | Optional | - |
| `IdpRole` | `[]string` | Optional | - |
| `IdpUsername` | `*string` | Optional | Username presented to the Identity Provider |
| `Mac` | `*string` | Optional | Client MAC address |
| `MxedgeId` | `*string` | Optional | Mist Edge ID used to connect to cloud |
| `NacruleId` | `*uuid.UUID` | Optional | NAC Policy Rule ID, if matched |
| `NacruleMatched` | `*bool` | Optional | - |
| `NasVendor` | `*string` | Optional | - |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `*string` | Optional | - |
| `PortType` | `*string` | Optional | Type of client i.e wired, wireless |
| `RandomMac` | `*bool` | Optional | - |
| `RespAttrs` | `[]string` | Optional | - |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | - |
| `Timestamp` | `*float64` | Optional | - |
| `Type` | `*string` | Optional | Event type, e.g. NAC_CLIENT_PERMIT |
| `UsermacLabel` | `[]string` | Optional | Labels derived from usermac entry |
| `Username` | `*string` | Optional | Username presented by the client |
| `Vlan` | `*string` | Optional | Vlan ID |
| `VlanSource` | `*string` | Optional | Vlan source, e.g. "nactag", "usermac" |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "5c5b35513227",
  "auth_type": "eap-tls",
  "bssid": "5c5b355fafcc",
  "dryrun_nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264012",
  "idp_id": "912ef72e-2239-4996-b81e-469e87a27cd6",
  "idp_role": [
    "itsuperusers",
    "vip"
  ],
  "idp_username": "user@deaflyz.net",
  "mac": "ac3eb179e535",
  "nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62",
  "nacrule_matched": true,
  "nas_vendor": "juniper-mist",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "random_mac": false,
  "resp_attrs": [
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous"
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "mist_nac",
  "timestamp": 1691512031.35819,
  "type": "NAC_CLIENT_PERMIT",
  "usermac_label": [
    "bldg5",
    "printer"
  ],
  "username": "user@deaflyz.net",
  "vlan": "750",
  "vlan_source": "nactag",
  "device_mac": "device_mac6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

