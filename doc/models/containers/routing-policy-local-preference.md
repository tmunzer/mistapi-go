
# Routing Policy Local Preference

Optional, for an import policy, local_preference can be changed, value in range 1-4294967294. Can be a Variable (e.g. `{{bgp_as}}`)

## Class Name

`RoutingPolicyLocalPreference`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.RoutingPolicyLocalPreferenceContainer.FromString(string mString) |
| `int` | models.RoutingPolicyLocalPreferenceContainer.FromNumber(int number) |

## string

### Initialization Code

#### Example

```go
value := models.RoutingPolicyLocalPreferenceContainer.FromString("String0")
```

## int

### Initialization Code

#### Example

```go
value := models.RoutingPolicyLocalPreferenceContainer.FromNumber(1)
```

