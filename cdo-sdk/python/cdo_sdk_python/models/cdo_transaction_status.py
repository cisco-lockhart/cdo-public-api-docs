# coding: utf-8

"""
    CDO API

    Use the documentation to explore the endpoints CDO has to offer

    The version of the OpenAPI document: 1.4.0
    Contact: cdo.tac@cisco.com
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
from enum import Enum
from typing_extensions import Self


class CdoTransactionStatus(str, Enum):
    """
    The status of the CDO transaction
    """

    """
    allowed enum values
    """
    PENDING = 'PENDING'
    IN_PROGRESS = 'IN_PROGRESS'
    DONE = 'DONE'
    ERROR = 'ERROR'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of CdoTransactionStatus from a JSON string"""
        return cls(json.loads(json_str))


