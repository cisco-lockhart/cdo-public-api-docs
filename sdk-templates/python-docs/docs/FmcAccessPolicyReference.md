# FmcAccessPolicyReference


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**link** | **str** | The endpoint to access this resource from. | [optional] 
**uid** | **str** | The unique identifier, represented as a UUID, of the FMC Access Policy in Security Cloud Control. | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.fmc_access_policy_reference import FmcAccessPolicyReference

# TODO update the JSON string below
json = "{}"
# create an instance of FmcAccessPolicyReference from a JSON string
fmc_access_policy_reference_instance = FmcAccessPolicyReference.from_json(json)
# print the JSON string representation of the object
print(FmcAccessPolicyReference.to_json())

# convert the object into a dict
fmc_access_policy_reference_dict = fmc_access_policy_reference_instance.to_dict()
# create an instance of FmcAccessPolicyReference from a dict
fmc_access_policy_reference_form_dict = fmc_access_policy_reference.from_dict(fmc_access_policy_reference_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


