// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { ISemver } from "@eth-optimism-bedrock/src/universal/interfaces/ISemver.sol";
import { INitroValidator } from "./INitroValidator.sol";

contract SystemConfigGlobal is OwnableUpgradeable, ISemver {
    /// @notice The AWS Nitro validator. 
    INitroValidator public nitroValidator;

    /// @notice The address of the proposer.
    address public proposer;

    /// @notice Mapping of valid PCR0s attested from AWS Nitro.
    mapping(bytes32 => bool) public validPCR0s;

    /// @notice Mapping of valid signers attested from AWS Nitro.
    mapping(address => bool) public validSigners;

    /// @notice Semantic version.
    /// @custom:semver 0.0.1
    function version() public pure virtual returns (string memory) {
        return "0.0.1";
    }

    constructor(INitroValidator _nitroValidator) {
        nitroValidator = _nitroValidator;
        initialize({
            _owner: address(0xdEaD)
        });
    }

    function initialize(address _owner) public initializer {
        __Ownable_init();
        transferOwnership(_owner);
    }

    function setProposer(address _proposer) external onlyOwner {
        proposer = _proposer;
    }

    function registerPCR0(bytes calldata pcr0) external onlyOwner {
        validPCR0s[keccak256(pcr0)] = true;
    }

    function deregisterPCR0(bytes calldata pcr0) external onlyOwner {
        delete validPCR0s[keccak256(pcr0)];
    }

    function registerSigner(bytes calldata attestation) external onlyOwner {
        (bytes memory enclavePublicKey, bytes memory pcr0) = nitroValidator.validateAttestation(attestation, 1 days); //todo maxAge setting
        require (validPCR0s[keccak256(pcr0)], "invalid pcr0 in attestation");

        bytes32 publicKeyHash = keccak256(enclavePublicKey);
        address enclaveAddress = address(uint160(uint256(publicKeyHash)));
        validSigners[enclaveAddress] = true; 
    }

    function deregisterSigner(address signer) external onlyOwner {
        delete validSigners[signer];
    }
}
