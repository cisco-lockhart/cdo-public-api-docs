# CompatibleSecureClientPlatformDto

An OS and CPU architecture combination available for a Secure Client version.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture. One of: ARM64, X86_64, UNIVERSAL. | 
**os** | **str** | The operating system. One of: LINUX, MACOS, WINDOWS. | 

## Example

```python
from scc_firewall_manager_sdk.models.compatible_secure_client_platform_dto import CompatibleSecureClientPlatformDto

# TODO update the JSON string below
json = "{}"
# create an instance of CompatibleSecureClientPlatformDto from a JSON string
compatible_secure_client_platform_dto_instance = CompatibleSecureClientPlatformDto.from_json(json)
# print the JSON string representation of the object
print(CompatibleSecureClientPlatformDto.to_json())

# convert the object into a dict
compatible_secure_client_platform_dto_dict = compatible_secure_client_platform_dto_instance.to_dict()
# create an instance of CompatibleSecureClientPlatformDto from a dict
compatible_secure_client_platform_dto_form_dict = compatible_secure_client_platform_dto.from_dict(compatible_secure_client_platform_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


