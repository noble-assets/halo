# CHANGELOG

## v2.0.1

*Nov 13, 2024*

This is a non-consensus breaking patch to the `v2` release line.

### BUG FIXES

- Switch entitlements method input to a query string. ([#13](https://github.com/noble-assets/halo/pull/13))

## v2.0.0

*Nov 10, 2024*

Second release of the `x/halo` module, aligning it with the latest developments in Cosmos and Noble.

### DEPENDENCIES

- Upgrade Cosmos SDK to the latest `v0.50.x` release.

### FEATURES

- Utilize [collections](https://docs.cosmos.network/v0.50/build/packages/collections) for managing module state.
- Support [app wiring](https://docs.cosmos.network/v0.50/build/building-apps/app-go-v2) for compatibility with Noble's core codebase.

### IMPROVEMENTS

- Reorganize repository to align with Noble's standards.
- Update module path for v2 release line.

## v1.0.1

*Sep 12, 2024*

This is a consensus breaking patch to the `v1` release line.

### BUG FIXES

- Ensure the recipient is a liquidity provider when trading to fiat. ([#8](https://github.com/noble-assets/halo/pull/8))

## v1.0.0

*Aug 26, 2024*

Initial release of the `x/halo` module, allowing [USYC](https://usyc.hashnote.com) native issuance on Noble.

