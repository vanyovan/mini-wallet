# Mini Wallet
Mini wallet service backend application

## Introduction

Mini Wallet is a simple wallet management system that allows users to perform operations like wallet initialization, enabling, viewing, and managing transactions.

## Prerequisites

- Go programming language (version 1.20.4 or later)
- SQLite3 database (version 3.42.0 or later)

## Installation

git clone https://github.com/your-username/mini-wallet.git

## Getting Started

1. Database using sqlite3 in db file named **database.db**

2. Configure the database connection details in the config.json file.

3. Run the application:
   ```bash
   go run ./cmd/main.go
   ```
   
   or you can use makefile
   ```bash
   make run
   ```

## ERD
You can access the ERD in `wallet.draw.io` file in, just open it in [Draw.io](https://app.diagrams.net/)

## API Endpoints

The following API endpoints are available:

- `POST /api/v1/init`: Initializes the wallet for a user.
- `POST /api/v1/wallet`: Enables the wallet for a user.
- `GET /api/v1/wallet`: Retrieves the details of the user's wallet.
- `GET /api/v1/wallet/transactions`: Retrieves the transactions of the user's wallet.
- `POST /api/v1/wallet/deposits`: Deposits funds into the user's wallet.
- `POST /api/v1/wallet/withdrawals`: Withdraws funds from the user's wallet.
- `PATCH /api/v1/wallet`: Disables the user's wallet.

## Database

The application uses an SQLite3 database named database.db. The database contains the following tables:

1. `mst_user`: Stores information about all users.

   Columns:
   - `user_id` (TEXT): User ID.
   - `token` (TEXT): User token.

2. `mst_wallet`: Stores information about wallets including the balance wallet.

   Columns:
   - `wallet_id` (TEXT): Wallet ID.
   - `owned_by` (TEXT): User ID of the wallet owner.
   - `status` (TEXT): Wallet status (enabled or disabled).
   - `enabled_at` (TIMESTAMP): Timestamp when the wallet was enabled.
   - `disabled_at` (TIMESTAMP): Timestamp when the wallet was disabled.
   - `balance` (REAL): Wallet balance.

3. `trx_wallet`: Stores information about wallet transactions.

   Columns:
   - `wallet_id` (TEXT): Wallet ID.
   - `transaction_id` (TEXT): Transaction ID.
   - `status` (TEXT): Transaction status.
   - `transacted_at` (TIMESTAMP): Timestamp of the transaction.
   - `type` (TEXT): Transaction type.
   - `amount` (REAL): Transaction amount.
   - `reference_id` (TEXT): Transaction reference ID.

## Testing

To run the unit tests, use the following command:

```bash
make test
```
