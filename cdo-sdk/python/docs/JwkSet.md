# JwkSet


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**keys** | [**List[JsonWebKey]**](JsonWebKey.md) | The JSON Web Key Set. | [optional] 

## Example

```python
from cdo_python_sdk.models.jwk_set import JwkSet

# TODO update the JSON string below
json = "{}"
# create an instance of JwkSet from a JSON string
jwk_set_instance = JwkSet.from_json(json)
# print the JSON string representation of the object
print(JwkSet.to_json())

# convert the object into a dict
jwk_set_dict = jwk_set_instance.to_dict()
# create an instance of JwkSet from a dict
jwk_set_form_dict = jwk_set.from_dict(jwk_set_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


