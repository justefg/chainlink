pragma solidity 0.6.2;

import "./AggregatorFacade.sol";
import "./HistoricAggregatorInterface.sol";
import "./Whitelisted.sol";

/**
 * @title A whitelisted facade for Historic Aggregator versions to conform to the new v0.6
 * Aggregator interface.
 */
contract WhitelistedFacade is AggregatorFacade, Whitelisted {

  constructor(address _aggregator, uint8 _decimals)
    public
    AggregatorFacade(_aggregator, _decimals)
    Whitelisted()
  {}

  /**
   * @notice get the latest completed round where the answer was updated
   */
  function latestRound()
    external
    view
    virtual
    override
    isWhitelisted()
    returns (uint256)
  {
    return _latestRound();
  }

  /**
   * @notice Reads the current answer from aggregator delegated to.
   */
  function latestAnswer()
    external
    view
    virtual
    override
    isWhitelisted()
    returns (int256)
  {
    return _latestAnswer();
  }

  /**
   * @notice Reads the last updated height from aggregator delegated to.
   */
  function latestTimestamp()
    external
    view
    virtual
    override
    isWhitelisted()
    returns (uint256)
  {
    return _latestTimestamp();
  }

  /**
   * @notice get past rounds answers
   * @param _roundId the answer number to retrieve the answer for
   */
  function getAnswer(uint256 _roundId)
    external
    view
    virtual
    override
    isWhitelisted()
    returns (int256)
  {
    return _getAnswer(_roundId);
  }

  /**
   * @notice get block timestamp when an answer was last updated
   * @param _roundId the answer number to retrieve the updated timestamp for
   */
  function getTimestamp(uint256 _roundId)
    external
    view
    virtual
    override
    isWhitelisted()
    returns (uint256)
  {
    return _getTimestamp(_roundId);
  }

  function getRoundData(uint256 _roundId)
    external
    view
    virtual
    override
    isWhitelisted()
    returns (
      uint256 roundId,
      int256 answer,
      uint64 startedAt,
      uint64 updatedAt,
      uint256 answeredInRound
    )
  {
    return _getRoundData(_roundId);
  }

  /**
   * @notice get all details about the latest round.
   * @return roundId is the round ID for which details were retrieved
   * @return answer is the answer for the given round
   * @return startedAt is the timestamp when the round was started. This is 0
   * if the round hasn't been started yet.
   * @return updatedAt is the timestamp when the round last was updated (i.e.
   * answer was last computed)
   * @return answeredInRound is the round ID of the round in which the answer
   * was computed. answeredInRound may be smaller than roundId when the round
   * timed out. answerInRound is equal to roundId when the round didn't time out
   * and was completed regularly.
   * @dev Note that for in-progress rounds (i.e. rounds that haven't yet received
   * maxSubmissions) answer and updatedAt may change between queries.
   */
  function latestRoundData()
    external
    view
    virtual
    override
    isWhitelisted()
    returns (
      uint256 roundId,
      int256 answer,
      uint64 startedAt,
      uint64 updatedAt,
      uint256 answeredInRound
    )
  {
    return _latestRoundData();
  }

}
