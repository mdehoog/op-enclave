// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Portal } from "./Portal.sol";
import { OutputOracle } from "./OutputOracle.sol";
import { SystemConfigOwnable } from "./SystemConfigOwnable.sol";
import { SystemConfig } from "@eth-optimism-bedrock/src/L1/SystemConfig.sol";
import { OptimismPortal } from "@eth-optimism-bedrock/src/L1/OptimismPortal.sol";
import { SuperchainConfig } from "@eth-optimism-bedrock/src/L1/SuperchainConfig.sol";
import { L2OutputOracle } from "@eth-optimism-bedrock/src/L1/L2OutputOracle.sol";
import { L1StandardBridge } from "@eth-optimism-bedrock/src/L1/L1StandardBridge.sol";
import { L1CrossDomainMessenger } from "@eth-optimism-bedrock/src/L1/L1CrossDomainMessenger.sol";
import { L1ERC721Bridge } from "@eth-optimism-bedrock/src/L1/L1ERC721Bridge.sol";
import { OptimismMintableERC20Factory } from "@eth-optimism-bedrock/src/universal/OptimismMintableERC20Factory.sol";
import { ResourceMetering } from "@eth-optimism-bedrock/src/L1/ResourceMetering.sol";

contract DeployChain {
    struct PerChainConfig {
        uint256 chainID;
        bytes32 genesisL1Hash;
        uint64 genesisL1Number;
        bytes32 genesisL2Hash;
        uint64 genesisL2Time;
        address depositContractAddress;
        address l1SystemConfigAddress;
    }

    address public immutable proxyAdmin;
    address public immutable portal;
    address public immutable systemConfig;
    address public immutable l1StandardBridge;
    address public immutable l1ERC721Bridge;
    address public immutable optimismMintableERC20Factory;
    address public immutable l1CrossDomainMessenger;
    address public immutable outputOracle;
    address public immutable superchainConfig;

    constructor(
        address _proxyAdmin,
        address _portal,
        address _systemConfig,
        address _l1StandardBridge,
        address _l1ERC721Bridge,
        address _optimismMintableERC20Factory,
        address _l1CrossDomainMessenger,
        address _outputOracle,
        address _superchainConfig
    ) {
        proxyAdmin = _proxyAdmin;
        portal = _portal;
        systemConfig = _systemConfig;
        l1StandardBridge = _l1StandardBridge;
        l1ERC721Bridge = _l1ERC721Bridge;
        optimismMintableERC20Factory = _optimismMintableERC20Factory;
        l1CrossDomainMessenger = _l1CrossDomainMessenger;
        outputOracle = _outputOracle;
        superchainConfig = _superchainConfig;
    }

    function setupProxy(address proxy, bytes32 salt) internal returns (address instance) {
        address _proxyAdmin = proxyAdmin;
        /// @solidity memory-safe-assembly
        assembly {
            let ptr := mload(0x40)
            mstore(ptr, 0x60678060095f395ff363204e1c7a60e01b5f5273000000000000000000000000)
            mstore(add(ptr, 0x14), shl(0x60, proxy))
            mstore(add(ptr, 0x28), 0x6004525f5f60245f730000000000000000000000000000000000000000000000)
            mstore(add(ptr, 0x31), shl(0x60, _proxyAdmin))
            mstore(add(ptr, 0x45), 0x5afa3d5f5f3e3d60201416604d573d5ffd5b5f5f365f5f51365f5f375af43d5f)
            mstore(add(ptr, 0x65), 0x5f3e5f3d91606557fd5bf3000000000000000000000000000000000000000000)
            instance := create2(0, ptr, 0x70, salt)
        }
        require(instance != address(0), "Proxy: create2 failed");
    }

    function deploy(bytes32 salt, uint256 chainID, bytes32 genesisHash) external {
        address _outputOracle = setupProxy(outputOracle, salt);
        address _systemConfig = setupProxy(systemConfig, salt);
        address _portal = setupProxy(portal, salt);
        address _l1CrossDomainMessenger = setupProxy(l1CrossDomainMessenger, salt);
        address _l1StandardBridge = setupProxy(l1StandardBridge, salt);
        address _l1ERC721Bridge = setupProxy(l1ERC721Bridge, salt);
        address _optimismMintableERC20Factory = setupProxy(optimismMintableERC20Factory, salt);

        PerChainConfig memory config = PerChainConfig({
            chainID: chainID,
            genesisL1Hash: blockhash(block.number-1),
            genesisL1Number: uint64(block.number-1),
            genesisL2Hash: genesisHash,
            genesisL2Time: uint64(block.timestamp-2),
            depositContractAddress: _portal,
            l1SystemConfigAddress: _systemConfig
        });
        bytes32 configHash = keccak256(abi.encodePacked(
            uint64(0), // version
            config.chainID,
            config.genesisL1Hash,
            config.genesisL1Number,
            config.genesisL2Hash,
            config.genesisL2Time,
            config.depositContractAddress,
            config.l1SystemConfigAddress
        ));

        OutputOracle(_outputOracle).initialize(configHash);

        SystemConfigOwnable _configTemplate = SystemConfigOwnable(systemConfig);
        SystemConfigOwnable(_systemConfig).initialize({
            _basefeeScalar: _configTemplate.basefeeScalar(),
            _blobbasefeeScalar: _configTemplate.blobbasefeeScalar(),
            _batcherHash: _configTemplate.batcherHash(),
            _gasLimit: _configTemplate.gasLimit(),
            _unsafeBlockSigner: _configTemplate.unsafeBlockSigner(),
            _config: _configTemplate.resourceConfig(),
            _batchInbox: _configTemplate.batchInbox(),
            _addresses: SystemConfig.Addresses({
                l1CrossDomainMessenger: _l1CrossDomainMessenger,
                l1ERC721Bridge: _l1ERC721Bridge,
                l1StandardBridge: _l1StandardBridge,
                disputeGameFactory: address(0),
                optimismPortal: _portal,
                optimismMintableERC20Factory: _optimismMintableERC20Factory,
                gasPayingToken: address(0)
            })
        });

        Portal(payable(_portal)).initialize(
            OutputOracle(_outputOracle),
            SystemConfig(_systemConfig),
            SuperchainConfig(superchainConfig)
        );

        L1CrossDomainMessenger(_l1CrossDomainMessenger).initialize(
            SuperchainConfig(superchainConfig),
            OptimismPortal(payable(_portal)),
            SystemConfig(_systemConfig)
        );

        L1StandardBridge(payable(_l1StandardBridge)).initialize(
            L1CrossDomainMessenger(_l1CrossDomainMessenger),
            SuperchainConfig(superchainConfig),
            SystemConfig(_systemConfig)
        );

        L1ERC721Bridge(_l1ERC721Bridge).initialize(
            L1CrossDomainMessenger(_l1CrossDomainMessenger),
            SuperchainConfig(superchainConfig)
        );

        OptimismMintableERC20Factory(_optimismMintableERC20Factory).initialize(
            _l1StandardBridge
        );
    }
}
