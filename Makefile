.PHONY: bindings
bindings:
	forge clean && forge build --extra-output-files abi
	mkdir -p bindings
	abigen --abi out/OutputOracle.sol/OutputOracle.abi.json --pkg bindings --type OutputOracle --out bindings/output_oracle.go
	abigen --abi out/Portal.sol/Portal.abi.json --pkg bindings --type Portal --out bindings/portal.go
	abigen --abi out/DeployChain.sol/DeployChain.abi.json --pkg bindings --type DeployChain --out bindings/deploy_chain.go
