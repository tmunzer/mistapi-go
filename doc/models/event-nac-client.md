
# Event Nac Client

*This model accepts additional fields of type interface{}.*

## Structure

`EventNacClient`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ap` | `*string` | Optional | - |
| `AuthType` | [`*models.NacAuthTypeEnum`](../../doc/models/nac-auth-type-enum.md) | Optional | enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap` |
| `Bssid` | `*string` | Optional | BSSID |
| `DeviceMac` | `*string` | Optional | MAC Address of the device (AP, Switch) the client is connected to |
| `DryrunNacruleId` | `*uuid.UUID` | Optional | NAC Policy Dry Run Rule ID, if present and matched |
| `DryrunNacruleMatched` | `*bool` | Optional | `true` if dryrun rule present and matched with priority, `false` if not matched or not present |
| `IdpId` | `*uuid.UUID` | Optional | If IDP is used, the id of the IDP configuration used |
| `IdpRole` | `[]string` | Optional | - |
| `IdpUsername` | `*string` | Optional | If IDP is used, the username presented to the Identity Provider |
| `Mac` | `*string` | Optional | Client MAC address |
| `MxedgeId` | `*string` | Optional | Mist Edge ID used to connect to cloud |
| `NacruleId` | `*uuid.UUID` | Optional | NAC Policy Rule ID, if matched |
| `NacruleMatched` | `*bool` | Optional | NAC Policy Rule Matched |
| `NasVendor` | `*string` | Optional | Vendor name of the NAS |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PortId` | `[]string` | Optional | Port-ids the client was connected to  for the specified duration |
| `PortType` | [`*models.NacAccessTypeEnum`](../../doc/models/nac-access-type-enum.md) | Optional | Type of network access. enum: `wireless`, `wired` |
| `RandomMac` | `*bool` | Optional | Whether the client is using randomized MAC Address or not |
| `RespAttrs` | `[]string` | Optional | List of Radius AVP returned by the Authentication Server<br><br>**Constraints**: *Unique Items Required* |
| `SiteId` | `*uuid.UUID` | Optional | - |
| `Ssid` | `*string` | Optional | SSIDs the client was connecting to |
| `Timestamp` | `*float64` | Optional | Epoch (seconds) |
| `Type` | `*string` | Optional | Event type, e.g. NAC_CLIENT_PERMIT. Use the [List NAC Events Definitions](../../doc/controllers/constants-events.md#list-nac-events-definitions) endpoint to get the full list of available values. |
| `UsermacLabel` | `[]string` | Optional | Labels derived from usermac entry |
| `Username` | `*string` | Optional | username assigned to the client |
| `Vlan` | `*string` | Optional | vlan that assigned to the client |
| `VlanSource` | `*string` | Optional | Vlan source, e.g. "nactag", "usermac" |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap": "5c5b35513227",
  "auth_type": "eap-tls",
  "bssid": "5c5b355fafcc",
  "device_mac": "60c78d8c7f6f",
  "dryrun_nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264012",
  "idp_id": "912ef72e-2239-4996-b81e-469e87a27cd6",
  "idp_role": [
    "itsuperusers",
    "vip"
  ],
  "idp_username": "user@deaflyz.net",
  "mac": "ac3eb179e535",
  "nacrule_id": "32f27e7d-ff26-4a9b-b3d1-ff9bcb264c62",
  "nas_vendor": "juniper-mist",
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "port_id": [
    "ge-0/0/17.0"
  ],
  "port_type": "wireless",
  "resp_attrs": [
    "Tunnel-Type=VLAN",
    "Tunnel-Medium-Type=IEEE-802",
    "Tunnel-Private-Group-Id=750",
    "User-Name=anonymous"
  ],
  "site_id": "441a1214-6928-442a-8e92-e1d34b8ec6a6",
  "ssid": "MyCorp-NAC",
  "type": "NAC_CLIENT_PERMIT",
  "usermac_label": [
    "bldg5",
    "printer"
  ],
  "vlan_source": "nactag",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

