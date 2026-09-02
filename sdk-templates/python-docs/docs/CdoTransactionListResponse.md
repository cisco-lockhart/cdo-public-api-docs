# CdoTransactionListResponse


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**items** | [**List[CdoTransaction]**](CdoTransaction.md) | Transactions created for the requested data sources | [optional] 

## Example

```python
from scc_firewall_manager_sdk.models.cdo_transaction_list_response import CdoTransactionListResponse

# TODO update the JSON string below
json = "{}"
# create an instance of CdoTransactionListResponse from a JSON string
cdo_transaction_list_response_instance = CdoTransactionListResponse.from_json(json)
# print the JSON string representation of the object
print(CdoTransactionListResponse.to_json())

# convert the object into a dict
cdo_transaction_list_response_dict = cdo_transaction_list_response_instance.to_dict()
# create an instance of CdoTransactionListResponse from a dict
cdo_transaction_list_response_form_dict = cdo_transaction_list_response.from_dict(cdo_transaction_list_response_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


