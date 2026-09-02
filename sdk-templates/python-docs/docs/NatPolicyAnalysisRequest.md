# NatPolicyAnalysisRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**List[Item]**](Item.md) | Policy analysis requests grouped by data source | 

## Example

```python
from scc_firewall_manager_sdk.models.nat_policy_analysis_request import NatPolicyAnalysisRequest

# TODO update the JSON string below
json = "{}"
# create an instance of NatPolicyAnalysisRequest from a JSON string
nat_policy_analysis_request_instance = NatPolicyAnalysisRequest.from_json(json)
# print the JSON string representation of the object
print(NatPolicyAnalysisRequest.to_json())

# convert the object into a dict
nat_policy_analysis_request_dict = nat_policy_analysis_request_instance.to_dict()
# create an instance of NatPolicyAnalysisRequest from a dict
nat_policy_analysis_request_form_dict = nat_policy_analysis_request.from_dict(nat_policy_analysis_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


