
# Setting Ssr

*This model accepts additional fields of type interface{}.*

## Structure

`SettingSsr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConductorHosts` | `[]string` | Optional | List of Conductor IP Addresses or Hosts to be used by the SSR Devices |
| `ConductorToken` | `*string` | Optional | Token to be used by the SSR Devices to connect to the Conductor |
| `DisableStats` | `*bool` | Optional | Disable stats collection on SSR devices |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "conductor_hosts": [
    "conductor_hosts4",
    "conductor_hosts5"
  ],
  "conductor_token": "conductor_token6",
  "disable_stats": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

