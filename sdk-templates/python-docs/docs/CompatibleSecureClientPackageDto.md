# CompatibleSecureClientPackageDto

The Secure Client package available for this platform.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**download_url** | **str** | Direct download URL for the package. Requires no authentication and does not expire. | 
**file_name** | **str** | The Secure Client package filename. | 

## Example

```python
from scc_firewall_manager_sdk.models.compatible_secure_client_package_dto import CompatibleSecureClientPackageDto

# TODO update the JSON string below
json = "{}"
# create an instance of CompatibleSecureClientPackageDto from a JSON string
compatible_secure_client_package_dto_instance = CompatibleSecureClientPackageDto.from_json(json)
# print the JSON string representation of the object
print(CompatibleSecureClientPackageDto.to_json())

# convert the object into a dict
compatible_secure_client_package_dto_dict = compatible_secure_client_package_dto_instance.to_dict()
# create an instance of CompatibleSecureClientPackageDto from a dict
compatible_secure_client_package_dto_form_dict = compatible_secure_client_package_dto.from_dict(compatible_secure_client_package_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


