# cdo_sdk_python.AIAssistantApi

All URIs are relative to *https://edge.us.cdo.cisco.com/api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ask_ai_assistant_in_existing_conversation**](AIAssistantApi.md#ask_ai_assistant_in_existing_conversation) | **POST** /v1/ai-assistant/conversations/{conversationUid} | Ask AI Assistant (Existing Conversation)
[**ask_ai_assistant_in_new_conversation**](AIAssistantApi.md#ask_ai_assistant_in_new_conversation) | **POST** /v1/ai-assistant/conversations | Ask AI Assistant (New Conversation)
[**get_ai_assistant_conversation_messages**](AIAssistantApi.md#get_ai_assistant_conversation_messages) | **GET** /v1/ai-assistant/conversations/{conversationUid}/messages | Get Messages
[**get_conversation**](AIAssistantApi.md#get_conversation) | **GET** /v1/ai-assistant/conversations/{conversationUid} | Get Conversation
[**get_conversations**](AIAssistantApi.md#get_conversations) | **GET** /v1/ai-assistant/conversations | Get Conversations


# **ask_ai_assistant_in_existing_conversation**
> CdoTransaction ask_ai_assistant_in_existing_conversation(conversation_uid, ai_question)

Ask AI Assistant (Existing Conversation)

Ask the AI Assistant a question in the context of an existing conversation with the AI Assistant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_question import AiQuestion
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    conversation_uid = 'conversation_uid_example' # str | 
    ai_question = cdo_sdk_python.AiQuestion() # AiQuestion | 

    try:
        # Ask AI Assistant (Existing Conversation)
        api_response = api_instance.ask_ai_assistant_in_existing_conversation(conversation_uid, ai_question)
        print("The response of AIAssistantApi->ask_ai_assistant_in_existing_conversation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->ask_ai_assistant_in_existing_conversation: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_uid** | **str**|  | 
 **ai_question** | [**AiQuestion**](AiQuestion.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | CDO Transaction object that can be used to track the status of the question. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ask_ai_assistant_in_new_conversation**
> CdoTransaction ask_ai_assistant_in_new_conversation(ai_question)

Ask AI Assistant (New Conversation)

Start a new conversation with the AI Assistant by asking a question.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_question import AiQuestion
from cdo_sdk_python.models.cdo_transaction import CdoTransaction
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    ai_question = cdo_sdk_python.AiQuestion() # AiQuestion | 

    try:
        # Ask AI Assistant (New Conversation)
        api_response = api_instance.ask_ai_assistant_in_new_conversation(ai_question)
        print("The response of AIAssistantApi->ask_ai_assistant_in_new_conversation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->ask_ai_assistant_in_new_conversation: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ai_question** | [**AiQuestion**](AiQuestion.md)|  | 

### Return type

[**CdoTransaction**](CdoTransaction.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**202** | CDO Transaction object that can be used to track the status of the question. |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_ai_assistant_conversation_messages**
> List[AiMessage] get_ai_assistant_conversation_messages(conversation_uid)

Get Messages

Get a list of messages in a single AI Assistant conversation. Note: this endpoint is not paginated, and returns the full list of messages.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_message import AiMessage
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    conversation_uid = 'conversation_uid_example' # str | The unique identifier of the conversation in CDO.

    try:
        # Get Messages
        api_response = api_instance.get_ai_assistant_conversation_messages(conversation_uid)
        print("The response of AIAssistantApi->get_ai_assistant_conversation_messages:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->get_ai_assistant_conversation_messages: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_uid** | **str**| The unique identifier of the conversation in CDO. | 

### Return type

[**List[AiMessage]**](AiMessage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Conversation Messages |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_conversation**
> AiConversation get_conversation(conversation_uid)

Get Conversation

Get an AI Assistant conversation by UID in the CDO tenant.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_conversation import AiConversation
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    conversation_uid = 'conversation_uid_example' # str | The unique identifier of the conversation in CDO.

    try:
        # Get Conversation
        api_response = api_instance.get_conversation(conversation_uid)
        print("The response of AIAssistantApi->get_conversation:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->get_conversation: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **conversation_uid** | **str**| The unique identifier of the conversation in CDO. | 

### Return type

[**AiConversation**](AiConversation.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Conversation object |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**404** | Entity not found. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **get_conversations**
> AiAssistantConversationPage get_conversations(limit=limit, offset=offset, sort=sort)

Get Conversations

Get a list of AI Assistant Conversations. Note: the total number of conversations is set to -1 as this information is currently unavailable.

### Example

* Bearer (JWT) Authentication (bearerAuth):

```python
import cdo_sdk_python
from cdo_sdk_python.models.ai_assistant_conversation_page import AiAssistantConversationPage
from cdo_sdk_python.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://edge.us.cdo.cisco.com/api/rest
# See configuration.py for a list of all supported configuration parameters.
configuration = cdo_sdk_python.Configuration(
    host = "https://edge.us.cdo.cisco.com/api/rest"
)

# The client must configure the authentication and authorization parameters
# in accordance with the API server security policy.
# Examples for each auth method are provided below, use the example that
# satisfies your auth use case.

# Configure Bearer authorization (JWT): bearerAuth
configuration = cdo_sdk_python.Configuration(
    access_token = os.environ["BEARER_TOKEN"]
)

# Enter a context with an instance of the API client
with cdo_sdk_python.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = cdo_sdk_python.AIAssistantApi(api_client)
    limit = '50' # str | The number of results to retrieve. (optional) (default to '50')
    offset = '0' # str | The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve `limit` results from the offset specified. (optional) (default to '0')
    sort = ['name:DESC'] # List[str] | The fields to sort results by. (optional)

    try:
        # Get Conversations
        api_response = api_instance.get_conversations(limit=limit, offset=offset, sort=sort)
        print("The response of AIAssistantApi->get_conversations:\n")
        pprint(api_response)
    except Exception as e:
        print("Exception when calling AIAssistantApi->get_conversations: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **str**| The number of results to retrieve. | [optional] [default to &#39;50&#39;]
 **offset** | **str**| The offset of the results retrieved. The CDO API uses the offset field to determine the index of the first result retrieved, and will retrieve &#x60;limit&#x60; results from the offset specified. | [optional] [default to &#39;0&#39;]
 **sort** | [**List[str]**](str.md)| The fields to sort results by. | [optional] 

### Return type

[**AiAssistantConversationPage**](AiAssistantConversationPage.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | List of Conversations |  -  |
**400** | Invalid input provided. Check the response for details. |  -  |
**401** | Request not authorized. |  -  |
**403** | User does not have sufficient privileges to perform this operation. |  -  |
**500** | Internal server error. |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

