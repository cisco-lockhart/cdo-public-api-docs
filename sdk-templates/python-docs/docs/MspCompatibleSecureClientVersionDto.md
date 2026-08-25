# MspCompatibleSecureClientVersionDto

A Secure Client version available across all requested managed tenants, with the platforms available for that version.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**platforms** | [**List[MspCompatibleSecureClientPlatformDto]**](MspCompatibleSecureClientPlatformDto.md) | The OS and CPU architecture combinations available for this Secure Client version across every requested managed tenant. | 
**version** | **str** | The Secure Client version string. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_compatible_secure_client_version_dto import MspCompatibleSecureClientVersionDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspCompatibleSecureClientVersionDto from a JSON string
msp_compatible_secure_client_version_dto_instance = MspCompatibleSecureClientVersionDto.from_json(json)
# print the JSON string representation of the object
print(MspCompatibleSecureClientVersionDto.to_json())

# convert the object into a dict
msp_compatible_secure_client_version_dto_dict = msp_compatible_secure_client_version_dto_instance.to_dict()
# create an instance of MspCompatibleSecureClientVersionDto from a dict
msp_compatible_secure_client_version_dto_form_dict = msp_compatible_secure_client_version_dto.from_dict(msp_compatible_secure_client_version_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


