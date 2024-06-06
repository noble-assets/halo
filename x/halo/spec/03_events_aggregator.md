# Aggregator Events

## BalanceReported

This event is emitted whenever a new round is reported.

```json
{
  "type": "halo.aggregator.v1.BalanceReported",
  "attributes": [
    {
      "key": "balance",
      "value": "4791611113"
    },
    {
      "key": "interest",
      "value": "701123"
    },
    {
      "key": "price",
      "value": "103782204"
    },
    {
      "key": "round_id",
      "value": "187"
    },
    {
      "key": "updated_at",
      "value": "1717589739"
    }
  ]
}
```

This event is emitted by the following transactions:

- [`halo.aggregator.v1.MsgReportBalance`](./02_messages_aggregator.md#report-balance)

## NextPriceReported

This event is emitted whenever a new next price is reported.

```json
{
  "type": "halo.aggregator.v1.NextPriceReported",
  "attributes": [
    {
      "key": "price",
      "value": "103780600"
    }
  ]
}
```

This event is emitted by the following transactions:

- [`halo.aggregator.v1.MsgSetNextPrice`](./02_messages_aggregator.md#set-next-price)

## OwnershipTransferred

This event is emitted whenever ownership is transferred.

```json
{
  "type": "halo.aggregator.v1.OwnershipTransferred",
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

- [`halo.aggregator.v1.MsgTransferOwnership`](./02_messages_aggregator.md#transfer-ownership)
