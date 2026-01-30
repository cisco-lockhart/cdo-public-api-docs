# Cause

Represents a specific cause for an access rule issue.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**cause_types** | **List[str]** | A set of cause types for the issue. | [optional] 
**causing_rule** | **str** | The rule that is causing the issue. | [optional] 
**causing_rule_idx** | **int** | The index of the causing rule. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cause import Cause

# TODO update the JSON string below
json = "{}"
# create an instance of Cause from a JSON string
cause_instance = Cause.from_json(json)
# print the JSON string representation of the object
print(Cause.to_json())

# convert the object into a dict
cause_dict = cause_instance.to_dict()
# create an instance of Cause from a dict
cause_form_dict = cause.from_dict(cause_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


