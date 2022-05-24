// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import { Address } from "@openzeppelin/contracts/utils/Address.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { Pausable } from "@openzeppelin/contracts/security/Pausable.sol";
import { EIP712,ECDSA } from "@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol";
import { Counters } from "@openzeppelin/contracts/utils/Counters.sol";
import { ICashier } from "./ICashier.sol";

contract Cashier is ICashier, Ownable, Pausable, EIP712 {
    using Address for address;
    using Counters for Counters.Counter;

    enum State {
        Preparing,
        Staking,
        Withdrawing
    }

    struct Stakeholder {
        uint256 staked;
        uint256 timestamp;
    }

    State private _state;
    Counters.Counter private _roundNum;
    address public override stakingToken;
    address public override rewardToken;
    uint256 public override rewardAmount;
    uint256 public override totalStaked;
    uint256 public override startTime;
    uint256 public override stopTime;
    uint256 public override duration;
    bool public rewardLoaded;

    mapping(address => Stakeholder) public stakeholders;

    mapping(address => Counters.Counter) private _nonces;
    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _PERMIT_TYPEHASH =
        keccak256("WithdrawRewardPermit(address staker,uint256 value,uint256 nonce,uint256 deadline)");

    event Staked(address indexed staker, uint256 amount);
    event Withdrawal(address indexed staker, uint256 rewardAmount);
    event Reward(address indexed staker, uint256 rewardAmount);

    event StakingPreparing(uint256 startTime,uint256 stopTime,uint256 rewardAmount,uint256 roundNum);
    event StakingEnabled(uint256 startTime,uint256 stopTime,uint256 rewardAmount,uint256 roundNum);
    event WithdrawalEnabled(uint256 startTime,uint256 stopTime,uint256 totalStaked,uint256 roundNum);

    constructor(
        address _stakingToken,
        address _rewardToken,
        uint256 _rewardAmount,
        uint256 _startTime,
        uint256 _stopTime
    ) EIP712("MetaBlox","1"){
        require( _stakingToken.isContract(), "Staking: stakingToken not a contract address");
        require( _rewardToken.isContract(),"Staking: rewardToken not a contract address");
        require( _rewardAmount > 0,"Staking: rewardAmount must be greater than zero");
        require(block.timestamp <= _startTime, "Staking: incorrect timestamp");

        stakingToken = _stakingToken;
        rewardToken = _rewardToken;
        rewardAmount = _rewardAmount;
        startTime = _startTime;
        stopTime = _stopTime;
        duration = _stopTime - _startTime;
        _roundNum.increment();
    }

    /**
     * @dev after the contract deployed , recharge spec token for rewarding in future
     *  Requirement: 
     *     Approve rewardToken to this contract
     */
    function loadReward() external {
        require(!rewardLoaded,"Staking: Rewards has loaded into the contract");
        IERC20(rewardToken).transferFrom( _msgSender(),address(this),rewardAmount);
        rewardLoaded = true;
        _state = State.Staking;
        emit StakingEnabled(startTime,stopTime,rewardAmount,roundNum());
    }

    /**
     * @dev start staking 
    */
    function stake(uint256 _amount) public virtual override whenNotPaused {
        require(state() == State.Staking,"ChangeState: can only enable staking while staking");
        require(rewardLoaded,"Staking: Rewards not loaded into the contract yet");
        require(_amount > 0, "Staking: amount can't be 0");
        require(block.timestamp >= startTime, "Staking: staking not started");
        require(block.timestamp < stopTime, "Staking: staking has stoped");

        Stakeholder storage stakeholder = stakeholders[_msgSender()];
        stakeholder.staked += _amount;
        if (stakeholder.timestamp == 0) {
            stakeholder.timestamp = block.timestamp;
        }
        totalStaked += _amount;

        IERC20(stakingToken).transferFrom(_msgSender(), address(this), _amount);

        emit Staked(_msgSender(), _amount);
    }

     /**
      * @dev withdraw principal only
      */
    function withdraw() public virtual override whenNotPaused {
        Stakeholder storage stakeholder = stakeholders[_msgSender()];
        require(stakeholder.staked > 0,"Withdrawing: you have not participated in staking");
        require(state() == State.Withdrawing, "Withdrawing: illegal state");
        _withdrawStaked(_msgSender(), stakeholder.staked);
        stakeholder.staked = 0;
        stakeholder.timestamp = 0;
    }

    function withdrawRewardPermit(
        uint256 value,
        uint256 deadline,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public virtual override whenNotPaused {
        Stakeholder storage stakeholder = stakeholders[_msgSender()];
        require(  stakeholder.staked > 0, "WithdrawRewardPermit: you have not participated in staking" );
        require( block.timestamp <= deadline, "WithdrawRewardPermit: expired deadline" );

        bytes32 structHash = keccak256(
            abi.encode(
                _PERMIT_TYPEHASH,
                owner(),
                _msgSender(),
                value,
                _useNonce(_msgSender()),
                deadline
            )
        );

        bytes32 hash = _hashTypedDataV4(structHash);

        address signer = ECDSA.recover(hash, v, r, s);
        require(signer == owner(), "WithdrawRewardPermit: invalid signature");

        _withdrawReward(_msgSender(), value);
    }

    function getRewardTokenBalance() public view override returns (uint256) {
        return IERC20(rewardToken).balanceOf(address(this));
    }

    function getStakingTokenBalance() public view override returns (uint256) {
        return IERC20(stakingToken).balanceOf(address(this));
    }

    function getStaked(address _stakeholder) public view returns (uint256) {
        return stakeholders[_stakeholder].staked;
    }

    function _withdrawStaked(address _to, uint256 _amount) internal {
        IERC20(stakingToken).transfer(_to, _amount);

        emit Withdrawal(_msgSender(), _amount);
    }

    function _withdrawReward(address _to, uint256 _reward) internal {
        require(rewardAmount > _reward,"WithdrawingReward: rewardAmount not enough");
        IERC20(rewardToken).transfer(_to, _reward);

        emit Reward(_msgSender(), _reward);
    }


    /**
     * @return The current state of the escrow.
     */
    function state() public view virtual returns (State) {
        return _state;
    }

    function enableWithdrawal() public virtual override onlyOwner {
        require(state() == State.Staking,"Staking: can only enable withdrawal while staking");
        _state = State.Withdrawing;
        emit WithdrawalEnabled(startTime,stopTime,totalStaked,roundNum());
    }

    /**
     * @dev Returns whether stakers can withdraw their principal + rewards.
     */
    function withdrawalAllowed() public view returns (bool) {
        return state() == State.Withdrawing;
    }

    

    /**
     * @dev "Consume a nonce": return the current value and increment.
     *
     * _Available since v4.1._
     */
    function _useNonce(address staker)internal virtual returns (uint256 current){
        Counters.Counter storage nonce = _nonces[staker];
        current = nonce.current();
        nonce.increment();
    }

    /**
     * @dev See {IERC20Permit-nonces}.
     */
    function nonces(address staker) public view virtual returns (uint256) {
        return _nonces[staker].current();
    }


     /**
      * @dev record current round number
      */
     function roundNum() public view override returns (uint256){
       return _roundNum.current();
     }


    /**
     * @dev open new round for staking
     */
    function newRound(uint256 _startTime,uint256 _stopTime,uint256 _rewardAmount) external virtual override onlyOwner{
        require(state() == State.Withdrawing,"Staking: can only start new round while state withdrawing");
        require(block.timestamp > stopTime + 60*60*24,"Staking: can only start the day after the last round");
        
         _roundNum.increment();
        rewardLoaded = false;
        rewardAmount = _rewardAmount;
        startTime = _startTime;
        stopTime = _stopTime;
        duration = _stopTime - _startTime;

        emit StakingPreparing(_startTime, _stopTime, _rewardAmount, roundNum());
    }




    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }

}