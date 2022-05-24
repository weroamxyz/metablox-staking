// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

interface ICashier {
    function stake(uint256 amount) external;

    function withdraw() external;

    function withdrawRewardPermit(
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
        ) external;

    function stakingToken() external view returns (address);

    function rewardToken() external view returns (address);

    function rewardAmount() external view returns (uint256);

    function totalStaked() external view returns (uint256);

    function getRewardTokenBalance() external view returns (uint256);

    function getStakingTokenBalance() external view returns (uint256);

    function startTime() external view returns (uint256);

    function stopTime() external view returns (uint256);

    function duration() external view returns (uint256);

    function enableWithdrawal() external;

    function roundNum() external view returns (uint256);

    function newRound(uint256 startTime,uint256 stopTime,uint256 rewardAmount) external;

}