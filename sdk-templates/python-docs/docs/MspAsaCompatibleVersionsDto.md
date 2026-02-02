# MspAsaCompatibleVersionsDto

Compatible ASA upgrade versions for MSP managed devices across multiple tenants.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compatible_versions** | [**List[CompatibleVersionInfo]**](CompatibleVersionInfo.md) | List of compatible ASA upgrade versions available for the specified devices. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_asa_compatible_versions_dto import MspAsaCompatibleVersionsDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspAsaCompatibleVersionsDto from a JSON string
msp_asa_compatible_versions_dto_instance = MspAsaCompatibleVersionsDto.from_json(json)
# print the JSON string representation of the object
print(MspAsaCompatibleVersionsDto.to_json())

# convert the object into a dict
msp_asa_compatible_versions_dto_dict = msp_asa_compatible_versions_dto_instance.to_dict()
# create an instance of MspAsaCompatibleVersionsDto from a dict
msp_asa_compatible_versions_dto_form_dict = msp_asa_compatible_versions_dto.from_dict(msp_asa_compatible_versions_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


