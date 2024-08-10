
# Wlan Dynamic Psk

for dynamic PSK where we get per_user PSK from Radius. dynamic_psk allows PSK to be selected at runtime depending on context (wlan/site/user/...) thus following configurations are assumed (currently)

* PSK will come from RADIUS server
* AP sends client MAC as username ans password (i.e. `enable_mac_auth` is assumed)
* AP sends BSSID:SSID as Caller-Station-ID
* `auth_servers` is required
* PSK will come from cloud WLC if source is cloud_psks
* default_psk will be used if cloud WLC is not available
* `multi_psk_only` and `psk` is ignored
* `pairwise` can only be wpa2-ccmp (for now, wpa3 support on the roadmap)

## Structure

`WlanDynamicPsk`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultPsk` | `*string` | Optional | default PSK to use if cloud WLC is not available, 8-63 characters<br>**Constraints**: *Minimum Length*: `8`, *Maximum Length*: `63` |
| `DefaultVlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `ForceLookup` | `*bool` | Optional | when 11r is enabled, we'll try to use the cached PMK, this can be disabled<br>`false` means auto<br>**Default**: `false` |
| `Source` | [`*models.DynamicPskSourceEnum`](../../doc/models/dynamic-psk-source-enum.md) | Optional | enum: `cloud_psks`, `radius`<br>**Default**: `"radius"` |
| `VlanIds` | [`[]models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |

## Example (as JSON)

```json
{
  "default_psk": "foryoureyesonly",
  "enabled": false,
  "force_lookup": false,
  "source": "cloud_psks",
  "default_vlan_id": "String3"
}
```

