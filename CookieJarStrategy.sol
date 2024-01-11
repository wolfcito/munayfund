// SPDX-License-Identifier: AGPL-3.0-only
pragma solidity 0.8.19;

// External Libraries
import {ReentrancyGuard} from "openzeppelin-contracts/contracts/security/ReentrancyGuard.sol";
//Interfaces
import {IAllo} from "../../../core/interfaces/IAllo.sol";
import {IRegistry} from "../../../core/interfaces/IRegistry.sol";
// Core Contracts
import {BaseStrategy} from "../../BaseStrategy.sol";
// Internal Libraries
import {Metadata} from "../../../core/libraries/Metadata.sol";

contract CookieJarStrategy is BaseStrategy, ReentrancyGuard {

    struct InitializeData {
        bool metadataRequired;
        uint256 distributionInterval; // Allocated Recipients can take funds from the jar each distributionInterval number of days
    }

    struct Recipient {
        address recipientAddress;
        uint256 totalDelieverables; // proo
        uint256 grantAmount;
        Metadata metadata;
        Status recipientStatus;
        uint256 applicationId;
        uint256 totalVotesReceived;
    }

    bool public metadataRequired;
    uint256 public distributionInterval;
    IRegistry private _registry;
    mapping(address => Recipient) private _recipients;

    event UpdatedRegistration(
        address indexed recipientId, uint256 applicationId, bytes data, address sender, Status status
    );


    constructor(address _allo, string memory _name) BaseStrategy(_allo, _name) {}

    function initialize(uint256 _poolId, bytes memory _data) external virtual override {
        (InitializeData memory initData) = abi.decode(_data, (InitializeData));
        __CookieJarSimpleStrategy_init(_poolId, initData);
        emit Initialized(_poolId, _data);
    }

    function __CookieJarSimpleStrategy_init(uint256 _poolId, InitializeData memory _initData) internal {
        // Initialize the BaseStrategy
        __BaseStrategy_init(_poolId);

        // Set the strategy specific variables
        metadataRequired = _initData.metadataRequired;
        distributionInterval = _initData.distributionInterval;
        _registry = allo.getRegistry();

        // Set the pool to active - this is required for the strategy to work and distribute funds
        _setPoolActive(true);
    }

    function _getRecipientStatus(address _recipientId) internal view override returns (Status) {
        return _getRecipient(_recipientId).recipientStatus;
    }

    function _isValidAllocator(address _allocator) internal view override returns (bool) {
        return allo.isPoolManager(poolId, _allocator);
    }

    function _registerRecipient(bytes memory _data, address _sender)
        internal
        override
        onlyActivePool
        returns (address recipientId)
    {
        address recipientAddress;
        Metadata memory metadata;

        (recipientAddress, metadata) = abi.decode(_data, (address, Metadata));

        // make sure that if metadata is required, it is provided
        if (metadataRequired && (bytes(metadata.pointer).length == 0 || metadata.protocol == 0)) {
            revert INVALID_METADATA();
        }

        // make sure the recipient address is not the zero address
        if (recipientAddress == address(0)) revert RECIPIENT_ERROR(recipientId);

        Recipient storage recipient = _recipients[recipientId];

        // update the recipients data
        recipient.recipientAddress = recipientAddress;
        recipient.metadata = metadata;
        ++recipient.applicationId;

        Status currentStatus = recipient.recipientStatus;

        if (currentStatus == Status.None) {
            // recipient registering new application
            recipient.recipientStatus = Status.Pending;
            emit Registered(recipientId, _data, _sender);
        } else {
            // recipient updating rejected/pending/appealed/accepted application
            if (currentStatus == Status.Rejected) {
                recipient.recipientStatus = Status.Appealed;
            } else if (currentStatus == Status.Accepted) {
                // recipient updating already accepted application
                recipient.recipientStatus = Status.Pending;
            }

            // emit the new status with the '_data' that was passed in
            emit UpdatedRegistration(recipientId, recipient.applicationId, _data, _sender, recipient.recipientStatus);
        }
    }

    function _allocate(bytes memory _data, address _sender)
        internal
        virtual
        override
        nonReentrant
        onlyPoolManager(_sender)
    {
        
    }

    function _distribute(address[] memory _recipientIds, bytes memory, address _sender)
        internal
        virtual
        override
        onlyPoolManager(_sender)
    {
        
    }

    function _getPayout(address _recipientId, bytes memory) internal view override returns (PayoutSummary memory) {
        Recipient memory recipient = _getRecipient(_recipientId);
        return PayoutSummary(recipient.recipientAddress, recipient.grantAmount);
    }

    function _getRecipient(address _recipientId) internal view returns (Recipient memory recipient) {
        recipient = _recipients[_recipientId];
    }

    function _isProfileMember(address _anchor, address _sender) internal view returns (bool) {
        IRegistry.Profile memory profile = _registry.getProfileByAnchor(_anchor);
        return _registry.isOwnerOrMemberOfProfile(profile.id, _sender);
    }
}