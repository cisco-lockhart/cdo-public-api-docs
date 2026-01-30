# AccessRuleDetailsContent


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**elements** | **List[str]** | The list of elements. | [optional] 
**name** | **str** | The name | 
**type** | **str** | The type of the policy object. | [optional] 
**uid** | **str** | The unique identifier. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.access_rule_details_content import AccessRuleDetailsContent

# TODO update the JSON string below
json = "{}"
# create an instance of AccessRuleDetailsContent from a JSON string
access_rule_details_content_instance = AccessRuleDetailsContent.from_json(json)
# print the JSON string representation of the object
print(AccessRuleDetailsContent.to_json())

# convert the object into a dict
access_rule_details_content_dict = access_rule_details_content_instance.to_dict()
# create an instance of AccessRuleDetailsContent from a dict
access_rule_details_content_form_dict = access_rule_details_content.from_dict(access_rule_details_content_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


