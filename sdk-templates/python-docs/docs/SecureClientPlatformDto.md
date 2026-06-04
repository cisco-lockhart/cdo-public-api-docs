# SecureClientPlatformDto

The OS and CPU architecture combinations available for this version.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture. One of: X86_64, ARM64, UNIVERSAL. | 
**os** | **str** | The operating system. One of: LINUX, MACOS, WINDOWS. | 

## Example

```python
from scc_firewall_manager_sdk.models.secure_client_platform_dto import SecureClientPlatformDto

# TODO update the JSON string below
json = "{}"
# create an instance of SecureClientPlatformDto from a JSON string
secure_client_platform_dto_instance = SecureClientPlatformDto.from_json(json)
# print the JSON string representation of the object
print(SecureClientPlatformDto.to_json())

# convert the object into a dict
secure_client_platform_dto_dict = secure_client_platform_dto_instance.to_dict()
# create an instance of SecureClientPlatformDto from a dict
secure_client_platform_dto_form_dict = secure_client_platform_dto.from_dict(secure_client_platform_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


