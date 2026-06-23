
# Next Hop Via

Next-hop IP address. Can be a single IP address or an array of IP addresses for ECMP (Equal-Cost Multi-Path) load balancing across multiple next-hops.

## Class Name

`NextHopVia`

## Cases

| Type | Factory Method |
|  --- | --- |
| `string` | models.NextHopViaContainer.FromString(string mString) |
| `[]string` | models.NextHopViaContainer.FromArrayOfString([]string arrayOfString) |

## string

### Initialization Code

#### Example

```go
value := models.NextHopViaContainer.FromString("String0")
```

## []string

### Initialization Code

#### Example

```go
value := models.NextHopViaContainer.FromArrayOfString([]string{
    "String1",
})
```

