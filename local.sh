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
  halod genesis add-genesis-account validator 1000000ustake --home .halo --keyring-backend test
  OWNER=$(halod keys add owner --home .halo --keyring-backend test --output json | jq .address)
  halod genesis add-genesis-account owner 10000000uusdc --home .halo --keyring-backend test
  REPORTER=$(halod keys add reporter --home .halo --keyring-backend test --output json | jq .address)
  halod genesis add-genesis-account reporter 10000000uusdc --home .halo --keyring-backend test
  halod keys add pending-reporter --home .halo --keyring-backend test &> /dev/null
  halod genesis add-genesis-account pending-reporter 10000000uusdc --home .halo --keyring-backend test
  ENTITLEMENTS_OWNER=$(halod keys add entitlements-owner --home .halo --keyring-backend test --output json | jq .address)
  halod genesis add-genesis-account entitlements-owner 10000000uusdc --home .halo --keyring-backend test
  halod keys add entitlements-pending-owner --home .halo --keyring-backend test &> /dev/null
  halod genesis add-genesis-account entitlements-pending-owner 10000000uusdc --home .halo --keyring-backend test

  TEMP=.halo/genesis.json
  touch $TEMP && jq '.app_state.staking.params.bond_denom = "ustake"' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.halo.aggregator_state.reporter = '$REPORTER'' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.halo.entitlements_state.owner = '$ENTITLEMENTS_OWNER'' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json
  touch $TEMP && jq '.app_state.halo.owner = '$OWNER'' .halo/config/genesis.json > $TEMP && mv $TEMP .halo/config/genesis.json

  halod genesis gentx validator 1000000ustake --chain-id "halo-1" --home .halo --keyring-backend test &> /dev/null
  halod genesis collect-gentxs --home .halo &> /dev/null

  sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' .halo/config/config.toml
fi

halod start --home .halo
