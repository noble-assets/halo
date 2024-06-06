# Aggregator Messages

## Report Balance

`halo.aggregator.v1.MsgReportBalance`

A message that reports a new round.

```json
{
  "body": {
    "messages": [
      {
        "@type": "/halo.aggregator.v1.MsgReportBalance",
        "signer": "noble1owner",
        "principal": "4790909990",
        "interest": "701123",
        "total_supply": "46169872257060",
        "next_price": "103794354"
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

- `principal`
- `interest`
- `total_supply`
- `next_price`

### Requirements

- Signer must be the current `owner`.

### State Changes

- `last_round_id`
- `next_price`
- `rounds`

### Events Emitted

- [`halo.aggregator.v1.BalanceReported`](./03_events_aggregator.md#balancereported)

## Set Next Price

`halo.aggregator.v1.MsgSetNextPrice`

A message that reports a new next price.

```json
{
  "body": {
    "messages": [
      {
        "@type": "/halo.aggregator.v1.MsgSetNextPrice",
        "signer": "noble1owner",
        "next_price": "103780600"
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

- `next_price` — The new next price to report.

### Requirements

- Signer must be the current `owner`.

### State Changes

- `next_price`

### Events Emitted

- [`halo.aggregator.v1.NextPriceReported`](./03_events_aggregator.md#nextpricereported)

## Transfer Ownership

`halo.aggregator.v1.MsgTransferOwnership`

A message that transfers ownership to a provided address.

```json
{
  "body": {
    "messages": [
      {
        "@type": "/halo.aggregator.v1.MsgTransferOwnership",
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

- [`halo.aggregator.v1.OwnershipTransferred`](./03_events_aggregator.md#ownershiptransferred)
