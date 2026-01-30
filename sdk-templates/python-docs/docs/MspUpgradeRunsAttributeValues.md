# MspUpgradeRunsAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**upgrade_run_statuses** | **List[str]** | Distinct status values for the MSP upgrade runs on this MSP portal. | [optional] 
**upgrade_run_types** | **List[str]** |  | [optional] 
**users** | **List[str]** |  | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.msp_upgrade_runs_attribute_values import MspUpgradeRunsAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of MspUpgradeRunsAttributeValues from a JSON string
msp_upgrade_runs_attribute_values_instance = MspUpgradeRunsAttributeValues.from_json(json)
# print the JSON string representation of the object
print(MspUpgradeRunsAttributeValues.to_json())

# convert the object into a dict
msp_upgrade_runs_attribute_values_dict = msp_upgrade_runs_attribute_values_instance.to_dict()
# create an instance of MspUpgradeRunsAttributeValues from a dict
msp_upgrade_runs_attribute_values_form_dict = msp_upgrade_runs_attribute_values.from_dict(msp_upgrade_runs_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


