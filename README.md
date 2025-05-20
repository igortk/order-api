# Order api

**Order-api** is an interface that allows users to interact with the stock exchange system. All available functionalities will be described below.

## 🛠 Tech Stack  
- **Backend:** Golang
- **REST library:** gin 
- **Message brocker:** RabbitMq

## 🚀 Features
- 📝 **Create Order** — Place a new order with volume, price, direction, and trading pair.
- 💱 **Get Exchange Rate** — Retrieve current exchange rates (min/avg/max) for selected currency pairs.
- 💰 **Emit User Balance** — Simulate or update a user’s balance with a specified currency and amount.
- 🧾 **Get User Balance** — Fetch detailed user balance information including locked and available funds.
- ❌ **Remove Order** — Cancel an existing order by ID.

## 📡 API Endpoints

### Create order
Request:
POST /v1/order/create
```sh
{
    "id":"632-request",
    "user_id":"632-user",
    "order_id":"632-order",
    "pair": "USD/EUR",
    "init_volume": 6.5,
    "init_price": 4.3,
    "direction": 1
}
```
Response:
```sh
{
  "id": "632-request",
  "order": {
    "order_id": "632-order",
    "user_id": "632-user",
    "pair": "USD/EUR",
    "init_volume": 6.5,
    "fill_volume": 6.5,
    "init_price": 4.3,
    "status": 1,
    "direction": 1,
    "updatedDate": 1710866507,
    "createdDate": 1710866507
  },
  "error": {}
}
```

### Get Exchange Rate
Request:
GET /v1/currency/exchange/rate
```sh
{
  "id": "456-exchange-rate",
  "pair": "EUR/CAD",
  "type": 2
}
```
Response:
```sh
{
  "id": "456-exchange-rate",
  "pairs": {
    "EUR/CAD": {
      "rate": [
        {
          "pair": "EUR/CAD",
          "min": "0.5733",
          "average": "0.6787",
          "max": "0.6909",
          "direction": 2
        },
        {
          "pair": "EUR/CAD",
          "min": "1.357",
          "average": "1.4735",
          "max": "1.5008",
          "direction": 1
        }
      ]
    }
  }
}
```

### User Balance Emit
Request:
GET /v1/user/balance/emit
```sh
{
  "id": "123-balance-emit",
  "user_id": "123-user_id",
  "currency": "TEST",
  "amount": 5.6
}
```
Response:
```sh
{
  "user_id": "123-user_id",
  "balances": [
    {
      "currency": "TEST",
      "balance": 5.6,
      "locked_balance": 2.8,
      "updated_date": 1718666679
    }
  ]
}
```

### Get User Balance
Request:
GET /v1/user/balance/info/{user_id}
Response:
```sh
{
  "id": "04bc8a90-f5fd-4920-937d-0d939b7a17ff",
  "user_balance": {
    "user_id": "123456789",
    "balances": [
      {
        "currency": "TEST3695",
        "balance": 50,
        "locked_balance": 25.5,
        "updated_date": 1718686048
      }
    ]
  }
}
```

### Remove order
Request:
GET /v1/order/remove
```sh
{
  "id": "632-remove-id",
  "order_id": "632-order"
}
```
Response:
```sh
{
  "id": "632-remove-id",
  "order": {
    "order_id": "632-order",
    "user_id": "632-user",
    "pair": "USD/EUR",
    "init_volume": 6.5,
    "fill_volume": 6.5,
    "init_price": 4.3,
    "status": 1,
    "direction": 1,
    "updatedDate": 1710866507,
    "createdDate": 1710866507
  },
  "error": {}
}
```

## 🐳 Running via docker
_comming soon_
