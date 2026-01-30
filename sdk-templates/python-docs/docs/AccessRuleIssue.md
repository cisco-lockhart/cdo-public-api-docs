# AccessRuleIssue

Represents an access rule issue, including issue type and causes.

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**causes** | [**List[Cause]**](Cause.md) | A list of the causes for the rule being marked with an issue. | [optional] 
**issue_type** | **str** | Indicates if the rule has any known issue. For unknown issues, UNKNOWN value will be used. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.access_rule_issue import AccessRuleIssue

# TODO update the JSON string below
json = "{}"
# create an instance of AccessRuleIssue from a JSON string
access_rule_issue_instance = AccessRuleIssue.from_json(json)
# print the JSON string representation of the object
print(AccessRuleIssue.to_json())

# convert the object into a dict
access_rule_issue_dict = access_rule_issue_instance.to_dict()
# create an instance of AccessRuleIssue from a dict
access_rule_issue_form_dict = access_rule_issue.from_dict(access_rule_issue_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


