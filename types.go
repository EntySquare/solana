package solana

import (
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/solplaydev/solana/utils"
)

// Predefined Solana account sizes
const (
	AccountSize       uint64 = 165 // 165 bytes
	FeeCalculatorSize uint64 = 8   // 8 bytes
	NonceAccountSize  uint64 = 80  // 80 bytes
	StakeAccountSize  uint64 = 200 // 200 bytes
	TokenAccountSize  uint64 = 165 // 165 bytes
	MintAccountSize   uint64 = 82  // 82 bytes
)

// Lookup table sizes
const (
	LookupTableMetaSize     uint64 = 56  // 56 bytes
	LookupTableMaxAddresses uint   = 256 // 256 addresses
)

const (
	// 1 SOL = 1e9 lamports
	SOL uint64 = 1e9

	// SPL token default decimals
	SPLTokenDefaultDecimals uint8 = 9

	// SPL token default multiplier for decimals
	SPLTokenDefaultMultiplier uint64 = 1e9

	// Solana Devnet RPC URL
	SolanaDevnetRPCURL = "https://api.devnet.solana.com"

	// Solana Mainnet RPC URL
	SolanaMainnetRPCURL = "https://api.mainnet-beta.solana.com"

	// Solana Testnet RPC URL
	SolanaTestnetRPCURL = "https://api.testnet.solana.com"
)

// TransactionStatus represents the status of a transaction.
type TransactionStatus uint8

// TransactionStatus enum.
const (
	TransactionStatusUnknown TransactionStatus = iota
	TransactionStatusSuccess
	TransactionStatusInProgress
	TransactionStatusFailure
)

// TransactionStatusStrings is a map of TransactionStatus to string.
var transactionStatusStrings = map[TransactionStatus]string{
	TransactionStatusUnknown:    "unknown",
	TransactionStatusSuccess:    "success",
	TransactionStatusInProgress: "in_progress",
	TransactionStatusFailure:    "failure",
}

// String returns the string representation of the transaction status.
func (s TransactionStatus) String() string {
	return transactionStatusStrings[s]
}

// ParseTransactionStatus parses the transaction status from the given string.
func ParseTransactionStatus(s rpc.Commitment) TransactionStatus {
	switch s {
	case rpc.CommitmentFinalized:
		return TransactionStatusSuccess
	case rpc.CommitmentConfirmed, rpc.CommitmentProcessed:
		return TransactionStatusInProgress
	default:
		return TransactionStatusUnknown
	}
}

// TokenAmount represents the amount of a token.
type TokenAmount struct {
	Amount         uint64  `json:"amount"`                     // amount in token lamports
	Decimals       uint8   `json:"decimals,omitempty"`         // number of decimals for the token; max 9
	UIAmount       float64 `json:"ui_amount,omitempty"`        // amount in token units, e.g. 1 SOL
	UIAmountString string  `json:"ui_amount_string,omitempty"` // amount in token units as a string, e.g. "1.000000000"
}

// TokenAmountFromLamports converts the given lamports to a token amount.
func NewTokenAmountFromLamports(lamports uint64, decimals uint8) TokenAmount {
	uiAmount := utils.AmountToFloat64(lamports, decimals)
	return TokenAmount{
		Amount:         lamports,
		Decimals:       decimals,
		UIAmount:       uiAmount,
		UIAmountString: utils.Float64ToString(uiAmount),
	}
}
