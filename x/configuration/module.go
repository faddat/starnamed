package configuration

import (
	"context"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/iov-one/starnamed/x/configuration/client/cli"
	"github.com/iov-one/starnamed/x/configuration/client/rest"
	"github.com/iov-one/starnamed/x/configuration/types"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the configuration module.
type AppModuleBasic struct {
	cdc codec.Marshaler
}

// RegisterLegacyAminoCodec registers the amino codec.
func (b AppModuleBasic) RegisterLegacyAminoCodec(amino *codec.LegacyAmino) {
	RegisterCodec(amino)
}

// RegisterGRPCGatewayRoutes registers the query handler client.
func (b AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, serveMux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), serveMux, types.NewQueryClient(clientCtx))
}

// Name returns the configuration module's name.
func (AppModuleBasic) Name() string { return types.ModuleName }

// DefaultGenesis returns default genesis state as raw bytes for the configuration module.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONMarshaler) json.RawMessage {
	defaultGenesisState := DefaultGenesisState()
	return cdc.MustMarshalJSON(&defaultGenesisState)
}

// ValidateGenesis performs genesis state validation for the configuration module.
func (b AppModuleBasic) ValidateGenesis(marshaler codec.JSONMarshaler, config client.TxEncodingConfig, message json.RawMessage) error {
	var data types.GenesisState
	err := marshaler.UnmarshalJSON(message, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

// RegisterRESTRoutes registers the REST routes for the configuration module.
func (AppModuleBasic) RegisterRESTRoutes(cliCtx client.Context, rtr *mux.Router) {
	rest.RegisterRoutes(cliCtx, rtr, AvailableQueries())
}

// GetTxCmd returns the root tx command for the configuration module.
func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns no root query command for the configuration module.
func (b AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}

// RegisterInterfaces implements InterfaceModule
func (b AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

// AppModule implements an application module for the configuration module.
type AppModule struct {
	AppModuleBasic
	keeper Keeper
}

// NewAppModule creates a new AppModule object.
func NewAppModule(k Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

// Name returns the configuration module's name.
func (AppModule) Name() string { return types.ModuleName }

// RegisterServices allows a module to register services
func (a AppModule) RegisterServices(configurator module.Configurator) {
	types.RegisterQueryServer(configurator.QueryServer(), NewQuerier(&a.keeper))
}

// LegacyQuerierHandler provides an sdk.Querier object that uses the legacy amino codec.
func (a AppModule) LegacyQuerierHandler(amino *codec.LegacyAmino) sdk.Querier {
	return NewLegacyQuerier(&a.keeper)
}

// RegisterInvariants registers the configuration module invariants.
func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// Route returns the message routing key for the configuration module.
func (a AppModule) Route() sdk.Route {
	return sdk.NewRoute(RouterKey, NewHandler(a.keeper))
}

// QuerierRoute returns the configuration module's querier route name.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// BeginBlock returns the begin blocker for the configuration module.
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock returns the end blocker for the configuration module. It returns no validator updates.
func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// InitGenesis performs genesis initialization for the configuration module. It returns no validator updates.
func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONMarshaler, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, a.keeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the exported genesis state as raw bytes for the configuration module.
func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONMarshaler) json.RawMessage {
	genesisState := ExportGenesis(ctx, a.keeper)
	return cdc.MustMarshalJSON(genesisState)
}
