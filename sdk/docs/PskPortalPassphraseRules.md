# PskPortalPassphraseRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlphabertsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Length** | Pointer to **int32** |  | [optional] 
**MaxLength** | Pointer to **int32** | for valid &#x60;max_length&#x60; and &#x60;min_length&#x60;, passphrase size is set randomly from that range. - if &#x60;max_length&#x60; and/or &#x60;min_length&#x60; are invalid, passphrase size is equal to &#x60;length&#x60; parameter - if &#x60;length&#x60; is not set or is invalid, default passphrase size is 8. valid &#x60;max_length&#x60;, &#x60;min_length&#x60;, &#x60;length&#x60; should be an integer between 8 to 63. Also, &#x60;max_length&#x60; &gt; &#x60;min_length&#x60; | [optional] 
**MinLength** | Pointer to **int32** | for valid &#x60;max_length&#x60; and &#x60;min_length&#x60;, passphrase size is set randomly from that range. - if &#x60;max_length&#x60; and/or &#x60;min_length&#x60; are invalid, passphrase size is equal to &#x60;length&#x60; parameter - if &#x60;length&#x60; is not set or is invalid, default passphrase size is 8. valid &#x60;max_length&#x60;, &#x60;min_length&#x60;, &#x60;length&#x60; should be an integer between 8 to 63. Also, &#x60;max_length&#x60; &gt; &#x60;min_length&#x60; | [optional] 
**NumericsEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Symbols** | Pointer to **string** |  | [optional] 
**SymbolsEnabled** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewPskPortalPassphraseRules

`func NewPskPortalPassphraseRules() *PskPortalPassphraseRules`

NewPskPortalPassphraseRules instantiates a new PskPortalPassphraseRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPskPortalPassphraseRulesWithDefaults

`func NewPskPortalPassphraseRulesWithDefaults() *PskPortalPassphraseRules`

NewPskPortalPassphraseRulesWithDefaults instantiates a new PskPortalPassphraseRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlphabertsEnabled

`func (o *PskPortalPassphraseRules) GetAlphabertsEnabled() bool`

GetAlphabertsEnabled returns the AlphabertsEnabled field if non-nil, zero value otherwise.

### GetAlphabertsEnabledOk

`func (o *PskPortalPassphraseRules) GetAlphabertsEnabledOk() (*bool, bool)`

GetAlphabertsEnabledOk returns a tuple with the AlphabertsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlphabertsEnabled

`func (o *PskPortalPassphraseRules) SetAlphabertsEnabled(v bool)`

SetAlphabertsEnabled sets AlphabertsEnabled field to given value.

### HasAlphabertsEnabled

`func (o *PskPortalPassphraseRules) HasAlphabertsEnabled() bool`

HasAlphabertsEnabled returns a boolean if a field has been set.

### GetLength

`func (o *PskPortalPassphraseRules) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PskPortalPassphraseRules) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PskPortalPassphraseRules) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *PskPortalPassphraseRules) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *PskPortalPassphraseRules) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PskPortalPassphraseRules) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PskPortalPassphraseRules) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PskPortalPassphraseRules) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *PskPortalPassphraseRules) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PskPortalPassphraseRules) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PskPortalPassphraseRules) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PskPortalPassphraseRules) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetNumericsEnabled

`func (o *PskPortalPassphraseRules) GetNumericsEnabled() bool`

GetNumericsEnabled returns the NumericsEnabled field if non-nil, zero value otherwise.

### GetNumericsEnabledOk

`func (o *PskPortalPassphraseRules) GetNumericsEnabledOk() (*bool, bool)`

GetNumericsEnabledOk returns a tuple with the NumericsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumericsEnabled

`func (o *PskPortalPassphraseRules) SetNumericsEnabled(v bool)`

SetNumericsEnabled sets NumericsEnabled field to given value.

### HasNumericsEnabled

`func (o *PskPortalPassphraseRules) HasNumericsEnabled() bool`

HasNumericsEnabled returns a boolean if a field has been set.

### GetSymbols

`func (o *PskPortalPassphraseRules) GetSymbols() string`

GetSymbols returns the Symbols field if non-nil, zero value otherwise.

### GetSymbolsOk

`func (o *PskPortalPassphraseRules) GetSymbolsOk() (*string, bool)`

GetSymbolsOk returns a tuple with the Symbols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbols

`func (o *PskPortalPassphraseRules) SetSymbols(v string)`

SetSymbols sets Symbols field to given value.

### HasSymbols

`func (o *PskPortalPassphraseRules) HasSymbols() bool`

HasSymbols returns a boolean if a field has been set.

### GetSymbolsEnabled

`func (o *PskPortalPassphraseRules) GetSymbolsEnabled() bool`

GetSymbolsEnabled returns the SymbolsEnabled field if non-nil, zero value otherwise.

### GetSymbolsEnabledOk

`func (o *PskPortalPassphraseRules) GetSymbolsEnabledOk() (*bool, bool)`

GetSymbolsEnabledOk returns a tuple with the SymbolsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbolsEnabled

`func (o *PskPortalPassphraseRules) SetSymbolsEnabled(v bool)`

SetSymbolsEnabled sets SymbolsEnabled field to given value.

### HasSymbolsEnabled

`func (o *PskPortalPassphraseRules) HasSymbolsEnabled() bool`

HasSymbolsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


