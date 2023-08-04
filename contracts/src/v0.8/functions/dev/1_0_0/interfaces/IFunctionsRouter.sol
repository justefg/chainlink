// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {FunctionsResponse} from "../libraries/FunctionsResponse.sol";

// @title Chainlink Functions Router interface.
interface IFunctionsRouter {
  struct Config {
    // Maximum number of consumers which can be added to a single subscription
    // This bound ensures we are able to loop over all subscription consumers as needed,
    // without exceeding gas limits.
    // Should a user require more consumers, they can use multiple subscriptions.
    uint16 maxConsumersPerSubscription;
    // Flat fee (in Juels of LINK) that will be paid to the Router owner for operation of the network
    uint96 adminFee;
    // The function selector that is used when calling back to the Client contract
    bytes4 handleOracleFulfillmentSelector;
    // List of max callback gas limits used by flag with GAS_FLAG_INDEX
    uint32[] maxCallbackGasLimits;
  }

  // @notice The identifier of the route to retrieve the address of the access control contract
  // The access control contract controls which accounts can manage subscriptions
  // @return id - bytes32 id that can be passed to the "getContractById" of the Router
  function getAllowListId() external pure returns (bytes32);

  // @notice The router configuration
  function getConfig() external view returns (Config memory);

  // @notice Sends a request (encoded as data) using the provided subscriptionId
  // @param subscriptionId A unique subscription ID allocated by billing system,
  // a client can make requests from different contracts referencing the same subscription
  // @param data Encoded Chainlink Functions request data, use FunctionsClient API to encode a request
  // @param dataVersion Gas limit for the fulfillment callback
  // @param callbackGasLimit Gas limit for the fulfillment callback
  // @param donId An identifier used to determine which route to send the request along
  // @return requestId A unique request identifier
  function sendRequest(
    uint64 subscriptionId,
    bytes calldata data,
    uint16 dataVersion,
    uint32 callbackGasLimit,
    bytes32 donId
  ) external returns (bytes32);

  // @notice Sends a request (encoded as data) to the proposed contracts
  // @param subscriptionId A unique subscription ID allocated by billing system,
  // a client can make requests from different contracts referencing the same subscription
  // @param data Encoded Chainlink Functions request data, use FunctionsClient API to encode a request
  // @param dataVersion Gas limit for the fulfillment callback
  // @param callbackGasLimit Gas limit for the fulfillment callback
  // @param donId An identifier used to determine which route to send the request along
  // @return requestId A unique request identifier
  function sendRequestToProposed(
    uint64 subscriptionId,
    bytes calldata data,
    uint16 dataVersion,
    uint32 callbackGasLimit,
    bytes32 donId
  ) external returns (bytes32);

  // @notice Fulfill the request by:
  // - calling back the data that the Oracle returned to the client contract
  // - pay the DON for processing the request
  // @dev Only callable by the Coordinator contract that is saved in the commitment
  // @param response response data from DON consensus
  // @param err error from DON consensus
  // @param juelsPerGas - current rate of juels/gas
  // @param costWithoutFulfillment - The cost of processing the request (in Juels of LINK ), without fulfillment
  // @param transmitter - The Node that transmitted the OCR report
  // @param commitment - The parameters of the request that must be held consistent between request and response time
  // @return fulfillResult -
  // @return callbackGasCostJuels -
  function fulfill(
    bytes memory response,
    bytes memory err,
    uint96 juelsPerGas,
    uint96 costWithoutFulfillment,
    address transmitter,
    FunctionsResponse.Commitment memory commitment
  ) external returns (FunctionsResponse.FulfillResult, uint96);

  // @notice Validate requested gas limit is below the subscription max.
  // @param subscriptionId subscription ID
  // @param callbackGasLimit desired callback gas limit
  function isValidCallbackGasLimit(uint64 subscriptionId, uint32 callbackGasLimit) external view;

  // @notice Get the current contract given an ID
  // @param id A bytes32 identifier for the route
  // @return contract The current contract address
  function getContractById(bytes32 id) external view returns (address);

  // @notice Get the proposed next contract given an ID
  // @param id A bytes32 identifier for the route
  // @return contract The current or proposed contract address
  function getProposedContractById(bytes32 id) external view returns (address);

  // @notice Return the latest proprosal set
  // @return timelockEndBlock The block number that the proposal is able to be merged at
  // @return ids The identifiers of the contracts to update
  // @return to The addresses of the contracts that will be updated to
  function getProposedContractSet() external view returns (uint, bytes32[] memory, address[] memory);

  // @notice Proposes one or more updates to the contract routes
  // @dev Only callable by owner
  function proposeContractsUpdate(bytes32[] memory proposalSetIds, address[] memory proposalSetAddresses) external;

  // @notice Updates the current contract routes to the proposed contracts
  // @dev Only callable once timelock has passed
  // @dev Only callable by owner
  function updateContracts() external;

  // @notice Proposes new configuration data for the Router contract itself
  // @dev Only callable by owner
  function proposeConfigUpdateSelf(bytes calldata config) external;

  // @notice Updates configuration data for the Router contract itself
  // @dev Only callable once timelock has passed
  // @dev Only callable by owner
  function updateConfigSelf() external;

  // @notice Proposes new configuration data for the current (not proposed) contract
  // @dev Only callable by owner
  function proposeConfigUpdate(bytes32 id, bytes calldata config) external;

  // @notice Sends new configuration data to the contract along a route route
  // @dev Only callable once timelock has passed
  // @dev Only callable by owner
  function updateConfig(bytes32 id) external;

  // @dev Puts the system into an emergency stopped state.
  // @dev Only callable by owner
  function pause() external;

  // @dev Takes the system out of an emergency stopped state.
  // @dev Only callable by owner
  function unpause() external;
}
