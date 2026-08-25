# CompatibleSecureClientPlatformDto

The OS and CPU architecture combinations available for this version, each with the package that serves it.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture. One of: X86_64, ARM64, UNIVERSAL. | 
**os** | **str** | The operating system. One of: LINUX, MACOS, WINDOWS. | 
**package** | [**CompatibleSecureClientPackageDto**](CompatibleSecureClientPackageDto.md) |  | 

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


