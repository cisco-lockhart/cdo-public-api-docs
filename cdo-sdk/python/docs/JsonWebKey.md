# JsonWebKey

The JSON Web Key Set.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**kty** | **str** | The family of cryptographic algorithms used with the key. | [optional] 
**e** | **str** | The exponent for the RSA public key. | [optional] 
**use** | **str** | How the key was meant to be used. | [optional] 
**kid** | **str** | The unique identifier for the key. | [optional] 
**alg** | **str** | The specific cryptographic algorithm used with the key. | [optional] 
**n** | **str** | The modulus for the RSA public key. | [optional] 

## Example

```python
from cdo_sdk_python.models.json_web_key import JsonWebKey

# TODO update the JSON string below
json = "{}"
# create an instance of JsonWebKey from a JSON string
json_web_key_instance = JsonWebKey.from_json(json)
# print the JSON string representation of the object
print(JsonWebKey.to_json())

# convert the object into a dict
json_web_key_dict = json_web_key_instance.to_dict()
# create an instance of JsonWebKey from a dict
json_web_key_form_dict = json_web_key.from_dict(json_web_key_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


