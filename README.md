# AWB Stock Allocation (Simulation)
## There's a feature in Belli we use to allocate the AWB number.

## Quickstart
### Run one of these commands on your command prompt:
```shell
  make run_http
```
or
```shell
  go mod tidy; go run cmd/http/*.go
```

### To generate the API docs using swagger, run this command:
```shell
  make swag
```
or
```shell
  swag init -g cmd/http/main.go ./docs; swag fmt
```

## The goals:
### AWB Stock Allocation Creation:
    Terms:
        > AWB stock data consists of AWB number, status
        > status: "in_use", "not_in_use" (or can use the integer too)

    1. AWB number validation:
        a. Is it a valid AWB number based on digit length?
        b. Is it a valid AWB number based on check digit?
    2. Does the AWB number already exist in the DB?

### Order Data CRUD
    Terms:
        > the order data consists of AWB number, sender, receiver, total weight,
        total price, and status
        > the price per kg is set with progressive price with these conditions:
            - 0kg - <10kg => IDR 5000
            - 10kg - <20kg => IDR 4500
            - 20kg - <25kg => IDR 4000
            - >25kg => IDR 3500
        > status list: "PENDING", "CONFIRM", "SHIPPED", "COMPLETED, "CANCELLED"

    1. Order creation:
        a. if the AWB already exists in the DB:
            - if status == "not_in_use", input order and set the status = "in_use"
            - if status == "in_use", return a proper error that the AWB number is not available or been in used
        b. if doesn't exist in the DB:
            - do the AWB number validation
            - if valid, create the AWB number on the AWB stocks and set the status = "in_use"
    2. Order update (just the status):
        a. add the status flow validation:
            - PENDING can only become CONFIRM or CANCELLED
            - CONFIRM can only become SHIPPED or CANCELLED
            - SHIPPED can only become COMPLETED
            - COMPLETED and CANCELLED can't be updated
        b. if the order is cancelled, set the AWB number on the stocks = "not_in_use"
    3. Order view (list):
        a. add pagination = 5 data per page
        b. add search by AWB number
    4. Order view (detail per order) -> fetch by order ID
    5. No deletion

### Notes:
- find the most suitable HTTP response
- use a clear message but not too sensitive
- you can adjust the status to be a label or an integer (it's up to you to decide, as long as user can see the status clearly)
- the dummy data is only stored in the list
- <b>[nice to have]</b> handle race condition for AWB number usage
- <b>[nice to have]</b> proper commit message (using conventional commits)

