package types

import (
	"bytes"
	"encoding/json"
	fmt "fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	commontypes "github.com/lavanet/lava/common/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/mitchellh/mapstructure"
)

const WILDCARD_CHAIN_POLICY = "*"

// gets the chainPolicy if exists, null safe
func (policy *Policy) ChainPolicy(chainID string) (chainPolicy ChainPolicy, allowed bool) {
	// empty policy | chainPolicies -> support all chains
	if policy == nil || len(policy.ChainPolicies) == 0 {
		return ChainPolicy{ChainId: chainID}, true
	}
	wildcard := false
	for _, chain := range policy.ChainPolicies {
		if chain.ChainId == chainID {
			return chain, true
		}
		if chain.ChainId == WILDCARD_CHAIN_POLICY {
			wildcard = true
		}
	}
	if wildcard {
		return ChainPolicy{ChainId: chainID}, true
	}
	return ChainPolicy{}, false
}

func (policy *Policy) GetSupportedAddons(specID string) (addons []string, err error) {
	chainPolicy, allowed := policy.ChainPolicy(specID)
	if !allowed {
		return nil, fmt.Errorf("specID %s not allowed by current policy", specID)
	}
	for _, collection := range chainPolicy.Collections {
		addons = append(addons, collection.AddOn)
	}
	return addons, nil
}

func (policy Policy) ValidateBasicPolicy(isPlanPolicy bool) error {
	// plan policy checks
	if isPlanPolicy {
		if policy.EpochCuLimit == 0 || policy.TotalCuLimit == 0 {
			return sdkerrors.Wrapf(ErrInvalidPolicyCuFields, `plan's compute units fields can't be zero 
				(EpochCuLimit = %v, TotalCuLimit = %v)`, policy.EpochCuLimit, policy.TotalCuLimit)
		}

		if policy.SelectedProvidersMode == SELECTED_PROVIDERS_MODE_DISABLED && len(policy.SelectedProviders) != 0 {
			return sdkerrors.Wrap(ErrInvalidSelectedProvidersConfig, `cannot configure mode = 3 (selected 
				providers feature is disabled) and non-empty list of selected providers`)
		}

		// non-plan policy checks
	} else if policy.SelectedProvidersMode == SELECTED_PROVIDERS_MODE_DISABLED {
		return sdkerrors.Wrap(ErrInvalidSelectedProvidersConfig, `cannot configure mode = 3 (selected 
				providers feature is disabled) for a policy that is not plan policy`)
	}

	// general policy checks
	if policy.EpochCuLimit > policy.TotalCuLimit {
		return sdkerrors.Wrapf(ErrInvalidPolicyCuFields, "EpochCuLimit can't be larger than TotalCuLimit (EpochCuLimit = %v, TotalCuLimit = %v)", policy.EpochCuLimit, policy.TotalCuLimit)
	}

	if policy.MaxProvidersToPair <= 1 {
		return sdkerrors.Wrapf(ErrInvalidPolicyMaxProvidersToPair, "invalid policy's MaxProvidersToPair fields (MaxProvidersToPair = %v)", policy.MaxProvidersToPair)
	}

	if policy.SelectedProvidersMode == SELECTED_PROVIDERS_MODE_ALLOWED && len(policy.SelectedProviders) != 0 {
		return sdkerrors.Wrap(ErrInvalidSelectedProvidersConfig, `cannot configure mode = 0 (no 
			providers restrictions) and non-empty list of selected providers`)
	}

	seen := map[string]bool{}
	for _, addr := range policy.SelectedProviders {
		_, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid selected provider address (%s)", err)
		}

		if seen[addr] {
			return sdkerrors.Wrapf(ErrInvalidSelectedProvidersConfig, "found duplicate provider address %s", addr)
		}
		seen[addr] = true
	}

	return nil
}

func GetStrictestChainPolicyForSpec(chainID string, policies []*Policy) (chainPolicyRet ChainPolicy, allowed bool) {
	allowedCollectionData := []spectypes.CollectionData{}
	for _, policy := range policies {
		chainPolicy, allowdChain := policy.ChainPolicy(chainID)
		if !allowdChain {
			return ChainPolicy{}, false
		}
		// get the strictest collection specification, while empty is allowed
		chainPolicyCollectionData := chainPolicy.Collections
		// if no collection data is specified in the policy previous allowed is stricter and no update is necessary
		if len(chainPolicyCollectionData) == 0 {
			continue
		}
		// this policy is limiting collection data so overwrite what is allowed
		if len(allowedCollectionData) == 0 {
			allowedCollectionData = chainPolicyCollectionData
			continue
		}
		// previous policies and current policy change collection data, we need the union of both
		allowedCollectionData = commontypes.Union(chainPolicyCollectionData, allowedCollectionData)
	}

	return ChainPolicy{ChainId: chainID, Collections: allowedCollectionData}, true
}

func VerifyTotalCuUsage(policies []*Policy, cuUsage uint64) bool {
	for _, policy := range policies {
		if policy != nil {
			if cuUsage >= policy.GetTotalCuLimit() {
				return false
			}
		}
	}

	return true
}

// allows unmarshaling parser func
func (s SELECTED_PROVIDERS_MODE) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(SELECTED_PROVIDERS_MODE_name[int32(s)])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (s *SELECTED_PROVIDERS_MODE) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	// Note that if the string cannot be found then it will be set to the zero value, 'Created' in this case.
	*s = SELECTED_PROVIDERS_MODE(SELECTED_PROVIDERS_MODE_value[j])
	return nil
}

// hook function to allow correct SELECTED_PROVIDERS_MODE enum read from yaml
func SelectedProvidersModeHookFunc() mapstructure.DecodeHookFuncType {
	return DecodeSelectedProvidersMode
}

func DecodeSelectedProvidersMode(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
	// Check that the data is string
	if f.Kind() != reflect.String {
		return data, nil
	}

	// Check that the target type is policy
	if t != reflect.TypeOf(SELECTED_PROVIDERS_MODE(0)) {
		return data, nil
	}

	dataStr, ok := data.(string)
	if ok {
		mode, found := SELECTED_PROVIDERS_MODE_value[dataStr]
		if found {
			return SELECTED_PROVIDERS_MODE(mode), nil
		} else {
			return 0, fmt.Errorf("invalid selected providers mode: %s", dataStr)
		}
	}

	return data, nil
}
