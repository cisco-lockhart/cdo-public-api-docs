# SecureClientPackageDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture. One of: X86_64, ARM64, UNIVERSAL. null if the filename could not be parsed. | [optional] 
**file_name** | **str** | The filename of the package as stored on the device. | 
**os** | **str** | The operating system. One of: LINUX, MACOS, WINDOWS. null if the filename could not be parsed. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the Secure Client package in Security Cloud Control. | 
**version** | **str** | The version string extracted from the filename. null if the filename could not be parsed. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.secure_client_package_dto import SecureClientPackageDto

# TODO update the JSON string below
json = "{}"
# create an instance of SecureClientPackageDto from a JSON string
secure_client_package_dto_instance = SecureClientPackageDto.from_json(json)
# print the JSON string representation of the object
print(SecureClientPackageDto.to_json())

# convert the object into a dict
secure_client_package_dto_dict = secure_client_package_dto_instance.to_dict()
# create an instance of SecureClientPackageDto from a dict
secure_client_package_dto_form_dict = secure_client_package_dto.from_dict(secure_client_package_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


