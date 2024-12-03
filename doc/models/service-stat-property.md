
# Service Stat Property

*This model accepts additional fields of type interface{}.*

## Structure

`ServiceStatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AshVersion` | `*string` | Optional | - |
| `CiaVersion` | `*string` | Optional | - |
| `EmberVersion` | `*string` | Optional | - |
| `IpsecClientVersion` | `*string` | Optional | - |
| `MistAgentVersion` | `*string` | Optional | - |
| `PackageVersion` | `*string` | Optional | - |
| `TestingToolsVersion` | `*string` | Optional | - |
| `WheeljackVersion` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ash_version": "ash_version6",
  "cia_version": "cia_version8",
  "ember_version": "ember_version8",
  "ipsec_client_version": "ipsec_client_version2",
  "mist_agent_version": "mist_agent_version4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

