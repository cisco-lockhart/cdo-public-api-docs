# MspCompatibleSecureClientPlatformDto

An OS and CPU architecture combination available for a Secure Client version.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture. One of: ARM64, X86_64, UNIVERSAL. | 
**os** | **str** | The operating system. One of: LINUX, MACOS, WINDOWS. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_compatible_secure_client_platform_dto import MspCompatibleSecureClientPlatformDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspCompatibleSecureClientPlatformDto from a JSON string
msp_compatible_secure_client_platform_dto_instance = MspCompatibleSecureClientPlatformDto.from_json(json)
# print the JSON string representation of the object
print(MspCompatibleSecureClientPlatformDto.to_json())

# convert the object into a dict
msp_compatible_secure_client_platform_dto_dict = msp_compatible_secure_client_platform_dto_instance.to_dict()
# create an instance of MspCompatibleSecureClientPlatformDto from a dict
msp_compatible_secure_client_platform_dto_form_dict = msp_compatible_secure_client_platform_dto.from_dict(msp_compatible_secure_client_platform_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


