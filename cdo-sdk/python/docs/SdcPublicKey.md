# SdcPublicKey

Information on the public key used to encrypt credentials sent to the SDC.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**base64_encoded_key** | **str** | Base64 RSA public key to use to encrypt device credentials sent to the SDC. | [optional] 
**key_id** | **str** | The identifier of the RSA public key. This identifier is used by the SDC to know which private key to use to decrypt a string. | [optional] 

## Example

```python
from cdo-python-sdk.models.sdc_public_key import SdcPublicKey

# TODO update the JSON string below
json = "{}"
# create an instance of SdcPublicKey from a JSON string
sdc_public_key_instance = SdcPublicKey.from_json(json)
# print the JSON string representation of the object
print(SdcPublicKey.to_json())

# convert the object into a dict
sdc_public_key_dict = sdc_public_key_instance.to_dict()
# create an instance of SdcPublicKey from a dict
sdc_public_key_form_dict = sdc_public_key.from_dict(sdc_public_key_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


