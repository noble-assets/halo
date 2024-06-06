# Entitlements Messages

## Pause

`halo.entitlements.v1.MsgPause`

A message that pauses USYC.

```json
{
  "body": {
    "messages": [
      {
        "@type": "/halo.entitlements.v1.MsgPause",
        "signer": "noble1owner"
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": []
}
```

### Arguments

This message takes no arguments.

### Requirements

- Signer must be the current `owner`.

### State Changes

- `paused`

### Events Emitted

- [`halo.entitlements.v1.Paused`](./03_events_entitlements.md#paused)

## Unpause

`halo.entitlements.v1.MsgUnpause`

A message that unpauses USYC.

```json
{
  "body": {
    "messages": [
      {
        "@type": "/halo.entitlements.v1.MsgUnpause",
        "signer": "noble1owner"
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": []
}
```

### Arguments

This message takes no arguments.

### Requirements

- Signer must be the current `owner`.

### State Changes

- `paused`

### Events Emitted

- [`halo.entitlements.v1.Unpaused`](./03_events_entitlements.md#unpaused)

## Transfer Ownership

`halo.entitlements.v1.MsgTransferOwnership`

A message that transfers ownership to a provided address.

```json
{
  "body": {
    "messages": [
      {
        "@type": "/halo.entitlements.v1.MsgTransferOwnership",
        "signer": "noble1signer",
        "new_owner": "noble1owner"
      }
    ],
    "memo": "",
    "timeout_height": "0",
    "extension_options": [],
    "non_critical_extension_options": []
  },
  "auth_info": {
    "signer_infos": [],
    "fee": {
      "amount": [],
      "gas_limit": "200000",
      "payer": "",
      "granter": ""
    }
  },
  "signatures": []
}
```

### Arguments

- `new_owner` — The Noble address to transfer ownership to.

### Requirements

- Signer must be the current `owner`.

### State Changes

- `owner`

### Events Emitted

- [`halo.entitlements.v1.OwnershipTransferred`](./03_events_entitlements.md#ownershiptransferred)
