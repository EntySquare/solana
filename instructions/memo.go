package instructions

import (
	"context"

	"github.com/EntySquare/solana-go-sdk/common"
	"github.com/EntySquare/solana-go-sdk/program/memo"
	"github.com/EntySquare/solana-go-sdk/types"
)

// Memo is the memo instruction.
func Memo(str string, signers ...common.PublicKey) InstructionFunc {
	return func(ctx context.Context, c Client) ([]types.Instruction, error) {
		return []types.Instruction{
			memo.BuildMemo(memo.BuildMemoParam{
				SignerPubkeys: signers,
				Memo:          []byte(str),
			}),
		}, nil
	}
}
