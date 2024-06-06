# Entitlements Events

## Paused

This event is emitted whenever USYC is paused.

```json
{
  "type": "halo.entitlements.v1.Paused",
  "attributes": [
    {
      "key": "account",
      "value": "noble1owner"
    }
  ]
}
```

This event is emitted by the following transactions:

- [`halo.entitlements.v1.MsgPause`](./02_messages_entitlements.md#pause)

## Unpaused

This event is emitted whenever USYC is unpaused.

```json
{
  "type": "halo.entitlements.v1.Unpaused",
  "attributes": [
    {
      "key": "account",
      "value": "noble1owner"
    }
  ]
}
```

This event is emitted by the following transactions:

- [`halo.entitlements.v1.MsgUnpause`](./02_messages_entitlements.md#unpause)

## OwnershipTransferred

This event is emitted whenever ownership is transferred.

```json
{
  "type": "halo.entitlements.v1.OwnershipTransferred",
  "attributes": [
    {
      "key": "new_owner",
      "value": "noble1owner"
    },
    {
      "key": "previous_owner",
      "value": "noble1signer"
    }
  ]
}
```

This event is emitted by the following transactions:

- [`halo.entitlements.v1.MsgTransferOwnership`](./02_messages_entitlements.md#transfer-ownership)
