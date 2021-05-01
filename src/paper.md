# Initial value assignment

1 KOLXI === 2LARI
2KG PIG MEAT == 24KOLXI

# What white papers are.
Technical documentation of how blockchain behaivs.

# Genesis file
The token supply, initial user balances, and global blockchain settings you define in a Genesis file.

# Database changes.
We record all database changes and update state.json file

# Example of transaction

1. Incoming transaction event
```json

  // Move 10 token from one acoount to another
  "ghostChaser": -10,
  "ghostBuster": +10

  // Example buying product
  "ghostBuster": -1
  "ghostChaser": +1
```

2. Recalculate balances state
```json
{
  "balances": {
    "ghostchaserr": 83,
    "ghostbuster": 17
  }
}
```

# Notes
1. Genesis balance never gets updated
2. Database state changes are callled transactions. (TX)

# Transaction and db state calculation.
Blockchain transactions represent a series of events, and the database is a final aggregated, calculated state after replaying all the transactions in a specific sequence.

# Notes.
Using transactions database state can be recovered.


# Mongo Collections
1. Accounts.

- Username
- Email
- Password
- Avatar

2. Transactions.
- From 
  - UserId
  - Username

- To
  - UserId
  - Username

3. Blocks
