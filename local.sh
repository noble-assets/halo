alias halod=./simapp/build/simd

for arg in "$@"
do
    case $arg in
        -r|--reset)
        rm -rf .halo
        shift
        ;;
    esac
done

if ! [ -f .halo/data/priv_validator_state.json ]; then
  halod init validator --chain-id "halo-1" --home .halo &> /dev/null

  halod keys add validator --home .halo --keyring-backend test &> /dev/null
  halod add-genesis-account validator 1000000ustake --home .halo --keyring-backend test
  OWNER=$(halod keys add owner --home .halo --keyring-backend test --output json | jq .address)
  halod add-genesis-account owner 10000000uusdc --home .halo --keyring-backend test
  AGGREGATOR_OWNER=$(halod keys add aggregator-owner --home .halo --keyring-backend test --output json | jq .address)
  halod add-genesis-account aggregator-owner 10000000uusdc --home .halo --keyring-backend test
  halod keys add aggregator-pending-owner --home .halo --keyring-backend test &> /dev/null
  halod add-genesis-account aggregator-pending-owner 10000000uusdc --home .halo --keyring-backend test
  ENTITLEMENTS_OWNER=$(halod keys add entitlements-owner --home .halo --keyring-backend test --output json | jq .address)
  halod add-genesis-account entitlements-owner 10000000uusdc --home .halo --keyring-backend test
  halod keys add entitlements-pending-owner --home .halo --keyring-backend test &> /dev/null
  halod add-genesis-account entitlements-pending-owner 10000000uusdc --home .halo --keyring-backend test

  TEMP=.halo/genesis.json
  touch $TEMP && jq '.app_state.bank.denom_metadata = [{ "description": "Circle USD Coin", "denom_units": [{ "denom": "uusdc", "exponent": 0, "aliases": ["microusdc"] }, { "denom": "usdc", "exponent": 6 }], "base": "uusdc", "display": "usdc", "name": "Circle USD Coin", "symbol": "USDC" }]' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.staking.params.bond_denom = "ustake"' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state."fiat-tokenfactory".mintingDenom = { "denom": "uusdc" }' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.halo.aggregator_state.owner = '$AGGREGATOR_OWNER'' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.halo.entitlements_state.owner = '$ENTITLEMENTS_OWNER'' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.halo.owner = '$OWNER'' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json

  halod gentx validator 1000000ustake --chain-id "halo-1" --home .halo --keyring-backend test &> /dev/null
  halod collect-gentxs --home .halo &> /dev/null

  sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' .halo/config/config.toml
fi

halod start --home .halo
