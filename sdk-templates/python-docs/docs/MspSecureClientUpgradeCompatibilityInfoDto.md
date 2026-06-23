# MspSecureClientUpgradeCompatibilityInfoDto

Compatible Secure Client upgrade versions for MSP managed ASA devices across multiple tenants.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**compatible_versions** | [**List[CompatibleSecureClientVersionDto]**](CompatibleSecureClientVersionDto.md) | List of Secure Client versions compatible with the specified ASA devices across managed tenants. A version is included only if it is compatible with all of the requested devices in every managed tenant they belong to. The platforms listed for each version are likewise the intersection of platforms available in every managed tenant. | [optional] 
**error_message** | **str** | Error message detailing the issues that occurred during compatibility information retrieval. This field is populated when the status is ERROR. | [optional] 
**status** | **str** | The current status of the retrieval of compatible Secure Client versions. | 
**uid** | **str** | The unique identifier, represented as a UUID, of this list of compatible Secure Client versions. | 

## Example

```python
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_compatibility_info_dto import MspSecureClientUpgradeCompatibilityInfoDto

# TODO update the JSON string below
json = "{}"
# create an instance of MspSecureClientUpgradeCompatibilityInfoDto from a JSON string
msp_secure_client_upgrade_compatibility_info_dto_instance = MspSecureClientUpgradeCompatibilityInfoDto.from_json(json)
# print the JSON string representation of the object
print(MspSecureClientUpgradeCompatibilityInfoDto.to_json())

# convert the object into a dict
msp_secure_client_upgrade_compatibility_info_dto_dict = msp_secure_client_upgrade_compatibility_info_dto_instance.to_dict()
# create an instance of MspSecureClientUpgradeCompatibilityInfoDto from a dict
msp_secure_client_upgrade_compatibility_info_dto_form_dict = msp_secure_client_upgrade_compatibility_info_dto.from_dict(msp_secure_client_upgrade_compatibility_info_dto_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


