package keeper

import (
	"testing"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/iov-one/starnamed/x/wasm/keeper/wasmtesting"
	"github.com/iov-one/starnamed/x/wasm/types"
	"github.com/stretchr/testify/assert"
)

func TestConstructorOptions(t *testing.T) {
	specs := map[string]struct {
		srcOpt Option
		verify func(*testing.T, Keeper)
	}{
		"wasm engine": {
			srcOpt: WithWasmEngine(&wasmtesting.MockWasmer{}),
			verify: func(t *testing.T, k Keeper) {
				assert.IsType(t, k.wasmVM, &wasmtesting.MockWasmer{})
			},
		},
		"message handler": {
			srcOpt: WithMessageHandler(&wasmtesting.MockMessageHandler{}),
			verify: func(t *testing.T, k Keeper) {
				assert.IsType(t, k.messenger, &wasmtesting.MockMessageHandler{})
			},
		},
		"query plugins": {
			srcOpt: WithQueryHandler(&wasmtesting.MockQueryHandler{}),
			verify: func(t *testing.T, k Keeper) {
				assert.IsType(t, k.wasmVMQueryHandler, &wasmtesting.MockQueryHandler{})
			},
		},
		"coin transferrer": {
			srcOpt: WithCoinTransferrer(&wasmtesting.MockCoinTransferrer{}),
			verify: func(t *testing.T, k Keeper) {
				assert.IsType(t, k.bank, &wasmtesting.MockCoinTransferrer{})
			},
		},
		"costs": {
			srcOpt: WithCosts(1, 2, 3),
			verify: func(t *testing.T, k Keeper) {
				t.Cleanup(setApiDefaults)
				assert.Equal(t, uint64(1), k.compileCost)
				assert.Equal(t, uint64(2), k.instanceCost)
				assert.Equal(t, uint64(3), k.gasMultiplier)
				assert.Equal(t, uint64(15), costHumanize)
				assert.Equal(t, uint64(12), costCanonical)
			},
		},
		"api costs": {
			srcOpt: WithApiCosts(1, 2),
			verify: func(t *testing.T, k Keeper) {
				t.Cleanup(setApiDefaults)
				assert.Equal(t, uint64(1), costHumanize)
				assert.Equal(t, uint64(2), costCanonical)
			},
		},
	}
	for name, spec := range specs {
		t.Run(name, func(t *testing.T) {
			k := NewKeeper(nil, nil, paramtypes.NewSubspace(nil, nil, nil, nil, ""), authkeeper.AccountKeeper{}, nil, stakingkeeper.Keeper{}, distributionkeeper.Keeper{}, nil, nil, nil, nil, nil, nil, "tempDir", types.DefaultWasmConfig(), SupportedFeatures, spec.srcOpt)
			spec.verify(t, k)
		})
	}

}

func setApiDefaults() {
	costHumanize = DefaultGasCostHumanAddress * DefaultGasMultiplier
	costCanonical = DefaultGasCostCanonicalAddress * DefaultGasMultiplier
}
