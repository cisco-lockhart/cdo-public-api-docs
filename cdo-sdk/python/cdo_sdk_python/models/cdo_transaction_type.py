# coding: utf-8

"""
    Cisco Defense Orchestrator API

    Use the interactive documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 0.0.1
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
from enum import Enum
from typing_extensions import Self


class CdoTransactionType(str, Enum):
    """
    the type of the transaction
    """

    """
    allowed enum values
    """
    ONBOARD_ASA = 'ONBOARD_ASA'
    ONBOARD_IOS = 'ONBOARD_IOS'
    ONBOARD_DUO_ADMIN_PANEL = 'ONBOARD_DUO_ADMIN_PANEL'
    CREATE_FTD = 'CREATE_FTD'
    REGISTER_FTD = 'REGISTER_FTD'
    DELETE_CDFMC_MANAGED_FTD = 'DELETE_CDFMC_MANAGED_FTD'
    RECONNECT_ASA = 'RECONNECT_ASA'
    READ_ASA = 'READ_ASA'
    DEPLOY_ASA_DEVICE_CHANGES = 'DEPLOY_ASA_DEVICE_CHANGES'
    INDEX_TENANT = 'INDEX_TENANT'
    TERMINATE_DEVICE_RA_VPN_SESSIONS = 'TERMINATE_DEVICE_RA_VPN_SESSIONS'
    REFRESH_RA_VPN_SESSIONS = 'REFRESH_RA_VPN_SESSIONS'
    TERMINATE_USER_RA_VPN_SESSIONS = 'TERMINATE_USER_RA_VPN_SESSIONS'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of CdoTransactionType from a JSON string"""
        return cls(json.loads(json_str))

