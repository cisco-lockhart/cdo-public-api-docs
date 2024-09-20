# CdoTransaction


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**tenant_uid** | **str** | The unique identifier, represented as a UUID, of the tenant that asynchronous transaction triggered on. | [optional] 
**sort_key** | **str** | DynamoDB sort key to provide us with efficient query capabilities. | [optional] 
**transaction_uid** | **str** | The unique identifier, represented as a UUID, of the asynchronous transaction triggered. | [optional] 
**entity_uid** | **str** | The unique identifier, represented as a UUID, of the entity that the asynchronous transaction is triggered on. | [optional] 
**entity_url** | **str** | A URL to access the entity that the asynchronous transaction is triggered on. | [optional] 
**transaction_polling_url** | **str** | The URL to poll to track the progress of the transaction. | [optional] 
**submission_time** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the transaction was triggered | [optional] 
**last_updated_time** | **datetime** | The time (UTC; represented using the RFC-3339 standard) at which the transaction status was last updated | [optional] 
**transaction_type** | [**CdoTransactionType**](CdoTransactionType.md) |  | [optional] 
**cdo_transaction_status** | [**CdoTransactionStatus**](CdoTransactionStatus.md) |  | [optional] 
**transaction_details** | **Dict[str, str]** | Transaction details, if any | [optional] 
**error_message** | **str** | Transaction error message, if any | [optional] 
**error_details** | **Dict[str, str]** | Transaction error details, if any | [optional] 
**expire_at** | **int** | TTL attribute detailing the expiry time this item should be deleted | [optional] 

## Example

```python
from cdo_sdk_python.models.cdo_transaction import CdoTransaction

# TODO update the JSON string below
json = "{}"
# create an instance of CdoTransaction from a JSON string
cdo_transaction_instance = CdoTransaction.from_json(json)
# print the JSON string representation of the object
print(CdoTransaction.to_json())

# convert the object into a dict
cdo_transaction_dict = cdo_transaction_instance.to_dict()
# create an instance of CdoTransaction from a dict
cdo_transaction_form_dict = cdo_transaction.from_dict(cdo_transaction_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


