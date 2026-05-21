# SecureClientUpgradeRunAttributeValues


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**upgrade_run_statuses** | **List[str]** | Distinct status values for the Secure Client upgrade runs in this tenant. | [optional] 
**upgrade_run_types** | **List[str]** | Distinct upgrade run type values for the Secure Client upgrade runs. | [optional] 
**users** | **List[str]** | Distinct usernames of users who initiated upgrade runs. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.secure_client_upgrade_run_attribute_values import SecureClientUpgradeRunAttributeValues

# TODO update the JSON string below
json = "{}"
# create an instance of SecureClientUpgradeRunAttributeValues from a JSON string
secure_client_upgrade_run_attribute_values_instance = SecureClientUpgradeRunAttributeValues.from_json(json)
# print the JSON string representation of the object
print(SecureClientUpgradeRunAttributeValues.to_json())

# convert the object into a dict
secure_client_upgrade_run_attribute_values_dict = secure_client_upgrade_run_attribute_values_instance.to_dict()
# create an instance of SecureClientUpgradeRunAttributeValues from a dict
secure_client_upgrade_run_attribute_values_form_dict = secure_client_upgrade_run_attribute_values.from_dict(secure_client_upgrade_run_attribute_values_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


