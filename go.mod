module github.com/EntySquare/solana

go 1.19

require (
	filippo.io/edwards25519 v1.0.0
	github.com/fatih/color v1.15.0
	github.com/joho/godotenv v1.5.1
	github.com/mr-tron/base58 v1.2.0
	github.com/near/borsh-go v0.3.2-0.20220516180422-1ff87d108454
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.7.0
	github.com/stretchr/testify v1.8.3
	github.com/tyler-smith/go-bip39 v1.1.0
)

require github.com/EntySquare/solana-go-sdk v0.0.0-20240515143930-e6ccc5f75684

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/crypto v0.9.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/blocto/solana-go-sdk => github.com/EntySquare/solana-go-sdk v1.23.8
