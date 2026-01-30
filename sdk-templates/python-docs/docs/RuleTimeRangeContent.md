# RuleTimeRangeContent

The optional time range for which the access rule is active.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.rule_time_range_content import RuleTimeRangeContent

# TODO update the JSON string below
json = "{}"
# create an instance of RuleTimeRangeContent from a JSON string
rule_time_range_content_instance = RuleTimeRangeContent.from_json(json)
# print the JSON string representation of the object
print(RuleTimeRangeContent.to_json())

# convert the object into a dict
rule_time_range_content_dict = rule_time_range_content_instance.to_dict()
# create an instance of RuleTimeRangeContent from a dict
rule_time_range_content_form_dict = rule_time_range_content.from_dict(rule_time_range_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


