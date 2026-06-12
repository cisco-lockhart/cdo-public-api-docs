# MspSecureClientPackageDto

The list of items retrieved.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cpu_architecture** | **str** | The CPU architecture the package targets. | [optional] 
**device_uid** | **str** | The unique identifier, represented as a UUID, of the device in Security Cloud Control that the package is installed on. | [optional] 
**file_name** | **str** | The file name of the Secure Client webdeploy package. | [optional] 
**os** | **str** | The operating system the package targets. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the Secure Client package. | [optional] 
**version** | **str** | The version of the Secure Client package. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_secure_client_package_dto import MspSecureClientPackageDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSecureClientPackageDto from a JSON string
msp_secure_client_package_dto_instance = MspSecureClientPackageDto.from_json(json)
# print the JSON string representation of the object
print(MspSecureClientPackageDto.to_json())

# convert the object into a dict
msp_secure_client_package_dto_dict = msp_secure_client_package_dto_instance.to_dict()
# create an instance of MspSecureClientPackageDto from a dict
msp_secure_client_package_dto_form_dict = msp_secure_client_package_dto.from_dict(msp_secure_client_package_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


