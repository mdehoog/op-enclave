guard-%:
	@ if [ "${${*}}" = "" ]; then echo "Environment variable $* not set" && exit 1; fi

.PHONY: bindings
bindings:
	go install github.com/ethereum/go-ethereum/cmd/abigen@v1.14.11
	forge clean && forge build --extra-output-files abi
	mkdir -p bindings
	abigen --abi out/OutputOracle.sol/OutputOracle.abi.json --pkg bindings --type OutputOracle --out bindings/output_oracle.go
	abigen --abi out/Portal.sol/Portal.abi.json --pkg bindings --type Portal --out bindings/portal.go
	abigen --abi out/DeployChain.sol/DeployChain.abi.json --pkg bindings --type DeployChain --out bindings/deploy_chain.go

.PHONY: deploy
deploy: guard-IMPL_SALT guard-DEPLOY_CONFIG_PATH guard-DEPLOY_PRIVATE_KEY guard-RPC_URL
	@forge script DeploySystem --sig deploy --rpc-url $(RPC_URL) \
		--private-key $(DEPLOY_PRIVATE_KEY) --broadcast

.PHONY: testnet
testnet: guard-L1_URL guard-DEPLOY_PRIVATE_KEY
	DEPLOY_CHAIN_ADDRESS=${DEPLOY_CHAIN_ADDRESS:-$(jq -r ".DeployChain" deployments/84532-deploy.json)} \
		go run ./testnet

.PHONY: verify
verify:
	deploy=broadcast/DeploySystem.s.sol/84532/deploy-latest.json; \
	addresses=$$(jq -r '.transactions[] | select(.transactionType=="CREATE" or .transactionType=="CREATE2") | .contractAddress' $$deploy); \
	for address in $$addresses; do \
		name=$$(jq -r --arg address "$$address" '.transactions[] | select((.transactionType=="CREATE" or .transactionType=="CREATE2") and .contractAddress==$$address) | .contractName' $$deploy); \
		arguments=$$(jq -r --arg address "$$address" '.transactions[] | select((.transactionType=="CREATE" or .transactionType=="CREATE2") and .contractAddress==$$address) | .arguments // [] | join(" ")' $$deploy); \
		constructor=$$(jq '.[] | select(.type=="constructor")' out/$$name.sol/$$name.abi.json | jq -r '.inputs | map(.type) | join(",")'); \
		echo "Verifying $$name @ $$address using constructor($$constructor) $$arguments"; \
		constructor_args=$$(cast abi-encode "constructor($$constructor)" $$arguments); \
		forge verify-contract --watch --verifier-url https://api-sepolia.basescan.org/api --constructor-args $$constructor_args $$address $$name ; \
	done
