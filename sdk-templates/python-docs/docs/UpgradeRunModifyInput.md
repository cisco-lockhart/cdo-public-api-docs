# UpgradeRunModifyInput


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **str** | A human-readable name for the upgrade run. | 

## Example

```python
from scc_firewall_manager_sdk.models.upgrade_run_modify_input import UpgradeRunModifyInput

# TODO update the JSON string below
json = "{}"
# create an instance of UpgradeRunModifyInput from a JSON string
upgrade_run_modify_input_instance = UpgradeRunModifyInput.from_json(json)
# print the JSON string representation of the object
print(UpgradeRunModifyInput.to_json())

# convert the object into a dict
upgrade_run_modify_input_dict = upgrade_run_modify_input_instance.to_dict()
# create an instance of UpgradeRunModifyInput from a dict
upgrade_run_modify_input_form_dict = upgrade_run_modify_input.from_dict(upgrade_run_modify_input_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


