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
import { Hashing } from "@eth-optimism-bedrock/src/libraries/Hashing.sol";
import { Types } from "@eth-optimism-bedrock/src/libraries/Types.sol";
import { Constants } from "@eth-optimism-bedrock/src/libraries/Constants.sol";

contract DeployChain {
    struct DeployAddresses {
        address l2OutputOracle;
        address systemConfig;
        address optimismPortal;
        address l1CrossDomainMessenger;
        address l1StandardBridge;
        address l1ERC721Bridge;
        address optimismMintableERC20Factory;
    }

    event Deploy(
        uint256 indexed chainID,
        bytes32 configHash,
        bytes32 outputRoot,
        address batchInbox,
        DeployAddresses addresses
    );

    bytes32 public constant MESSAGE_PASSER_STORAGE_HASH = 0x8ed4baae3a927be3dea54996b4d5899f8c01e7594bf50b17dc1e741388ce3d12;

    address public immutable proxyAdmin;
    address public immutable optimismPortal;
    address public immutable systemConfig;
    address public immutable l1StandardBridge;
    address public immutable l1ERC721Bridge;
    address public immutable optimismMintableERC20Factory;
    address public immutable l1CrossDomainMessenger;
    address public immutable l2OutputOracle;
    address public immutable superchainConfig;
    address public immutable protocolVersions;

    constructor(
        address _proxyAdmin,
        address _optimismPortal,
        address _systemConfig,
        address _l1StandardBridge,
        address _l1ERC721Bridge,
        address _optimismMintableERC20Factory,
        address _l1CrossDomainMessenger,
        address _l2OutputOracle,
        address _superchainConfig,
        address _protocolVersions
    ) {
        proxyAdmin = _proxyAdmin;
        optimismPortal = _optimismPortal;
        systemConfig = _systemConfig;
        l1StandardBridge = _l1StandardBridge;
        l1ERC721Bridge = _l1ERC721Bridge;
        optimismMintableERC20Factory = _optimismMintableERC20Factory;
        l1CrossDomainMessenger = _l1CrossDomainMessenger;
        l2OutputOracle = _l2OutputOracle;
        superchainConfig = _superchainConfig;
        protocolVersions = _protocolVersions;
    }

    function deployAddresses(uint256 chainID) external view returns (DeployAddresses memory) {
        bytes32 salt = keccak256(abi.encodePacked(chainID));
        return DeployAddresses({
            l2OutputOracle: proxyAddress(l2OutputOracle, salt),
            systemConfig: proxyAddress(systemConfig, salt),
            optimismPortal: proxyAddress(optimismPortal, salt),
            l1CrossDomainMessenger: proxyAddress(l1CrossDomainMessenger, salt),
            l1StandardBridge: proxyAddress(l1StandardBridge, salt),
            l1ERC721Bridge: proxyAddress(l1ERC721Bridge, salt),
            optimismMintableERC20Factory: proxyAddress(optimismMintableERC20Factory, salt)
        });
    }

    function deploy(
        uint256 chainID,
        uint64 genesisL1Number,
        bytes32 genesisL2Hash,
        bytes32 genesisL2StateRoot,
        uint64 genesisL2Time,
        uint32 basefeeScalar,
        uint32 blobbasefeeScalar,
        uint64 gasLimit,
        address batcherAddress,
        address unsafeBlockSigner
    ) external {
        bytes32 salt = keccak256(abi.encodePacked(chainID));
        address _l2OutputOracle = setupProxy(l2OutputOracle, salt);
        address _systemConfig = setupProxy(systemConfig, salt);
        address _optimismPortal = setupProxy(optimismPortal, salt);
        address _l1CrossDomainMessenger = setupProxy(l1CrossDomainMessenger, salt);
        address _l1StandardBridge = setupProxy(l1StandardBridge, salt);
        address _l1ERC721Bridge = setupProxy(l1ERC721Bridge, salt);
        address _optimismMintableERC20Factory = setupProxy(optimismMintableERC20Factory, salt);

        bytes32 scalar = bytes32((uint256(0x01) << 248) | (uint256(blobbasefeeScalar) << 32) | basefeeScalar);
        bytes32 genesisL1Hash = blockhash(uint256(genesisL1Number));
        bytes32 configHash = keccak256(abi.encodePacked(
            uint64(0), // version
            chainID,
            genesisL1Hash,
            genesisL1Number,
            genesisL2Hash,
            genesisL2Time,
            batcherAddress,
            scalar,
            gasLimit,
            _optimismPortal,
            _systemConfig
        ));

        bytes32 genesisOutputRoot = Hashing.hashOutputRootProof(Types.OutputRootProof({
            version: 0,
            stateRoot: genesisL2StateRoot,
            messagePasserStorageRoot: MESSAGE_PASSER_STORAGE_HASH,
            latestBlockhash: genesisL2Hash
        }));

        OutputOracle(_l2OutputOracle).initialize(configHash, genesisOutputRoot);

        address batchInbox = calculateBatchInbox(chainID);
        SystemConfigOwnable(_systemConfig).initialize({
            _basefeeScalar: basefeeScalar,
            _blobbasefeeScalar: blobbasefeeScalar,
            _batcherHash: bytes32(uint256(uint160(batcherAddress))),
            _gasLimit: gasLimit,
            _unsafeBlockSigner: unsafeBlockSigner,
            _config: Constants.DEFAULT_RESOURCE_CONFIG(),
            _batchInbox: batchInbox,
            _addresses: SystemConfig.Addresses({
                l1CrossDomainMessenger: _l1CrossDomainMessenger,
                l1ERC721Bridge: _l1ERC721Bridge,
                l1StandardBridge: _l1StandardBridge,
                disputeGameFactory: address(0),
                optimismPortal: _optimismPortal,
                optimismMintableERC20Factory: _optimismMintableERC20Factory,
                gasPayingToken: address(0)
            })
        });

        Portal(payable(_optimismPortal)).initialize(
            OutputOracle(_l2OutputOracle),
            SystemConfig(_systemConfig),
            SuperchainConfig(superchainConfig)
        );

        L1CrossDomainMessenger(_l1CrossDomainMessenger).initialize(
            SuperchainConfig(superchainConfig),
            OptimismPortal(payable(_optimismPortal)),
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

        emit Deploy({
            chainID: chainID,
            configHash: configHash,
            outputRoot: genesisOutputRoot,
            batchInbox: batchInbox,
            addresses: DeployAddresses({
                l2OutputOracle: _l2OutputOracle,
                systemConfig: _systemConfig,
                optimismPortal: _optimismPortal,
                l1CrossDomainMessenger: _l1CrossDomainMessenger,
                l1StandardBridge: _l1StandardBridge,
                l1ERC721Bridge: _l1ERC721Bridge,
                optimismMintableERC20Factory: _optimismMintableERC20Factory
            })
        });
    }

    function calculateBatchInbox(uint256 chainID) public pure returns (address) {
        uint256 inbox = 0;
        for (; chainID > 0; chainID /= 10) {
            inbox = (inbox << 4) | (chainID % 10);
        }
        return address(uint160(inbox | (0xff << 152)));
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

    function proxyAddress(address proxy, bytes32 salt) internal view returns (address predicted) {
        address _proxyAdmin = proxyAdmin;
        address deployer = address(this);
        /// @solidity memory-safe-assembly
        assembly {
            let ptr := mload(0x40)
            mstore(ptr, 0x60678060095f395ff363204e1c7a60e01b5f5273000000000000000000000000)
            mstore(add(ptr, 0x14), shl(0x60, proxy))
            mstore(add(ptr, 0x28), 0x6004525f5f60245f730000000000000000000000000000000000000000000000)
            mstore(add(ptr, 0x31), shl(0x60, _proxyAdmin))
            mstore(add(ptr, 0x45), 0x5afa3d5f5f3e3d60201416604d573d5ffd5b5f5f365f5f51365f5f375af43d5f)
            mstore(add(ptr, 0x65), 0x5f3e5f3d91606557fd5bf3ff0000000000000000000000000000000000000000)
            mstore(add(ptr, 0x71), shl(0x60, deployer))
            mstore(add(ptr, 0x85), salt)
            mstore(add(ptr, 0xa5), keccak256(ptr, 0x70))
            predicted := keccak256(add(ptr, 0x70), 0x55)
        }
    }
}
