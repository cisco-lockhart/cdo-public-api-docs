# MspSecureClientUpgradeRunsAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**upgrade_run_statuses** | **List[str]** | Distinct status values for the MSP Secure Client upgrade runs on this MSP portal. | [optional] 
**upgrade_run_types** | **List[str]** | Distinct upgrade run type values for the MSP Secure Client upgrade runs. | [optional] 
**users** | **List[str]** | Distinct usernames of users who initiated upgrade runs. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_secure_client_upgrade_runs_attribute_values import MspSecureClientUpgradeRunsAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspSecureClientUpgradeRunsAttributeValues from a JSON string
msp_secure_client_upgrade_runs_attribute_values_instance = MspSecureClientUpgradeRunsAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspSecureClientUpgradeRunsAttributeValues.to_json())

# convert the object into a dict
msp_secure_client_upgrade_runs_attribute_values_dict = msp_secure_client_upgrade_runs_attribute_values_instance.to_dict()
# create an instance of MspSecureClientUpgradeRunsAttributeValues from a dict
msp_secure_client_upgrade_runs_attribute_values_form_dict = msp_secure_client_upgrade_runs_attribute_values.from_dict(msp_secure_client_upgrade_runs_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


