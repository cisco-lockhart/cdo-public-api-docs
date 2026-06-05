# CompatibleSecureClientVersionDto


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**platforms** | [**List[SecureClientPlatformDto]**](SecureClientPlatformDto.md) | The OS and CPU architecture combinations available for this version. | 
**version** | **str** | The Secure Client version string (e.g. \&quot;5.1.15.287\&quot;, \&quot;4.10.05095\&quot;). | 

## Example

```python
from scc_firewall_manager_sdk.models.compatible_secure_client_version_dto import CompatibleSecureClientVersionDto

# TODO update the JSON string below
json = "{}"
# create an instance of CompatibleSecureClientVersionDto from a JSON string
compatible_secure_client_version_dto_instance = CompatibleSecureClientVersionDto.from_json(json)
# print the JSON string representation of the object
print(CompatibleSecureClientVersionDto.to_json())

# convert the object into a dict
compatible_secure_client_version_dto_dict = compatible_secure_client_version_dto_instance.to_dict()
# create an instance of CompatibleSecureClientVersionDto from a dict
compatible_secure_client_version_dto_form_dict = compatible_secure_client_version_dto.from_dict(compatible_secure_client_version_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


