# UpgradeCompatibilityInfoDto

Compatible upgrade versions for MSP managed devices across multiple tenants.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compatible_versions** | [**List[CompatibleVersionInfoDto]**](CompatibleVersionInfoDto.md) | List of compatible upgrade versions available for the specified devices. | [optional] 
**error_message** | **str** | Error message detailing the issues that occurred during compatibility information retrieval. This field is populated when the status is ERROR and contains more detailed error information. | [optional] 
**status** | **str** | The current status of the retrieval of upgrade versions. | 
**uid** | **str** | The unique identifier, represented as a UUID, of this list of compatible versions. | 

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_compatibility_info_dto import UpgradeCompatibilityInfoDto

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeCompatibilityInfoDto from a JSON string
upgrade_compatibility_info_dto_instance = UpgradeCompatibilityInfoDto.from_json(json)
# print the JSON string representation of the object
print(UpgradeCompatibilityInfoDto.to_json())

# convert the object into a dict
upgrade_compatibility_info_dto_dict = upgrade_compatibility_info_dto_instance.to_dict()
# create an instance of UpgradeCompatibilityInfoDto from a dict
upgrade_compatibility_info_dto_form_dict = upgrade_compatibility_info_dto.from_dict(upgrade_compatibility_info_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


