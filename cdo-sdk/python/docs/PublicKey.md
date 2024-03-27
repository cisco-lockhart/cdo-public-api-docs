# PublicKey

Information on the public key used to encrypt credentials sent to the SDC.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**base64_encoded_key** | **str** | Base64 RSA public key to use to encrypt device credentials sent to the SDC. | [optional] 
**key_id** | **str** | The identifier of the RSA public key. This identifier is used by the SDC to know which private key to use to decrypt a string. | [optional] 

## Example

```python
from cdo_sdk_python.models.public_key import PublicKey

# TODO update the JSON string below
json = "{}"
# create an instance of PublicKey from a JSON string
public_key_instance = PublicKey.from_json(json)
# print the JSON string representation of the object
print(PublicKey.to_json())

# convert the object into a dict
public_key_dict = public_key_instance.to_dict()
# create an instance of PublicKey from a dict
public_key_form_dict = public_key.from_dict(public_key_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


