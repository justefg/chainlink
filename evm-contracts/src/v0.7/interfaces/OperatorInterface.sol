// SPDX-License-Identifier: MIT
pragma solidity ^0.7.0;

interface OperatorInterface {
  function fulfillOracleRequest2(
    bytes32 requestId,
    uint256 payment,
    address callbackAddress,
    bytes4 callbackFunctionId,
    uint256 expiration,
    bytes calldata data
  ) external returns (bool);
  function operatorTransferAndCall(address to, uint256 value, bytes calldata data) external returns (bool success);
  function distributeFunds(address payable[] calldata receivers,uint[] calldata amounts) external payable;
}
