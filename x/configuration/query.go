package configuration

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/iov-one/starnamed/pkg/queries"
	abci "github.com/tendermint/tendermint/abci/types"
)

// QueryHandlerFunc defines the query handler for this module
type QueryHandlerFunc func(ctx sdk.Context, path []string, query abci.RequestQuery, k Keeper) ([]byte, error)

// AvailableQueries returns the list of available queries in the module
func AvailableQueries() []queries.QueryHandler {
	queries := []queries.QueryHandler{
		&QueryConfiguration{},
		&QueryFees{},
	}
	return queries
}

// queryRouter defines a router for domain queries
type queryRouter map[string]QueryHandlerFunc

func buildRouter(qrs []queries.QueryHandler) queryRouter {
	// queryHandler extends the default query handler
	// to provide also an handler function required to
	// build a router
	type queryHandler interface {
		queries.QueryHandler
		Handler() QueryHandlerFunc
	}
	// build router
	router := make(queryRouter, len(qrs))
	for _, query := range qrs {
		queryAndHandler, ok := query.(queryHandler)
		// if interface is not implemented then the query type formation is invalid
		if !ok {
			panic(fmt.Sprintf("invalid query type: %T", query))
		}
		router[queryAndHandler.QueryPath()] = queryAndHandler.Handler()
	}
	// return
	return router
}

// QueryConfiguration is the request model used to get the configuration
type QueryConfiguration struct{}

// Use is a placeholder
func (q *QueryConfiguration) Use() string {
	return "query-config"
}

// Description is a placeholder
func (q *QueryConfiguration) Description() string {
	return "return the current configuration"
}

// Handler implements QueryHandler
func (q *QueryConfiguration) Handler() QueryHandlerFunc {
	return queryConfigurationHandler
}

// Validate implements QueryHandler
func (q *QueryConfiguration) Validate() error {
	return nil
}

// QueryPath implements QueryHandler
func (q *QueryConfiguration) QueryPath() string {
	return "configuration"
}

// queryAccountsInDomainHandler returns all accounts in aliceAddr domain
func queryConfigurationHandler(ctx sdk.Context, _ []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	cfg := k.GetConfiguration(ctx)
	// return response
	respBytes, err := queries.DefaultQueryEncode(QueryConfigurationResponse{Config: cfg})
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return respBytes, nil
}

// QueryConfigurationResponse is the response returned after querying the configuration
type QueryConfigurationResponse struct {
	// Config represents the current configurations
	Config Config `json:"configuration"`
}

// QueryFees is the request model used to get the current fees
type QueryFees struct{}

// Use is a placeholder
func (q *QueryFees) Use() string {
	return "query-fees"
}

// Description is a placeholder
func (q *QueryFees) Description() string {
	return "return the current fees"
}

// Handler implements QueryHandler
func (q *QueryFees) Handler() QueryHandlerFunc {
	return queryFeesHandler
}

// Validate implements QueryHandler
func (q *QueryFees) Validate() error {
	return nil
}

// QueryPath implements QueryHandler
func (q *QueryFees) QueryPath() string {
	return "fees"
}

func queryFeesHandler(ctx sdk.Context, _ []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	fees := k.GetFees(ctx)
	// return response
	respBytes, err := queries.DefaultQueryEncode(QueryFeesResponse{Fees: *fees})
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return respBytes, nil
}

// QueryFeesResponse is returned after querying fees
type QueryFeesResponse struct {
	// Fees represents the current fees of the network
	Fees Fees `json:"fees"`
}
