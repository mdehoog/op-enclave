// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import { Clones } from "@openzeppelin/contracts/proxy/Clones.sol";
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
    address public immutable portal;
    address public immutable systemConfig;
    address public immutable l1StandardBridge;
    address public immutable l1ERC721Bridge;
    address public immutable optimismMintableERC20Factory;
    address public immutable l1CrossDomainMessenger;
    address public immutable outputOracle;
    address public immutable superchainConfig;

    constructor(
        address _portal,
        address _systemConfig,
        address _l1StandardBridge,
        address _l1ERC721Bridge,
        address _optimismMintableERC20Factory,
        address _l1CrossDomainMessenger,
        address _outputOracle,
        address _superchainConfig
    ) {
        portal = _portal;
        systemConfig = _systemConfig;
        l1StandardBridge = _l1StandardBridge;
        l1ERC721Bridge = _l1ERC721Bridge;
        optimismMintableERC20Factory = _optimismMintableERC20Factory;
        l1CrossDomainMessenger = _l1CrossDomainMessenger;
        outputOracle = _outputOracle;
        superchainConfig = _superchainConfig;
    }

    function deploy(bytes32 salt) public {
        address _outputOracle = Clones.cloneDeterministic(outputOracle, salt);
        address _systemConfig = Clones.cloneDeterministic(systemConfig, salt);
        address _portal = Clones.cloneDeterministic(portal, salt);
        address _l1CrossDomainMessenger = Clones.cloneDeterministic(l1CrossDomainMessenger, salt);
        address _l1StandardBridge = Clones.cloneDeterministic(l1StandardBridge, salt);
        address _l1ERC721Bridge = Clones.cloneDeterministic(l1ERC721Bridge, salt);
        address _optimismMintableERC20Factory = Clones.cloneDeterministic(optimismMintableERC20Factory, salt);

        OutputOracle(_outputOracle).initialize();

        SystemConfig _configTemplate = SystemConfig(systemConfig);
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
            L2OutputOracle(_outputOracle),
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
