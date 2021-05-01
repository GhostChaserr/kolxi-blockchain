# Initial value assignment
ტოკენების მონეტიზაცია, რაში იცვლება ტოკენი. ტოკენის გაცვლა ხდება ვალიუში.

1 KOLXI === 2LARI
2KG PIG MEAT == 12KOLXI TOKEN

# mutating global state
ყოველი ცვლილება აისაახება ტრანზაქციიტ, რომელიც თავისმხრივ ახდენს იუზერის ბალანსის რეკალკულაციას.

მაგგალითი, ბაზისს ცვლილების
ghostchaser -3
ghostchaser +3

ყოველი ტრანზაქცია update ს აკეთებს გლობალური state ის.

ტრანზაქცია -> state update


! account
! transaction


მაგალითი, პროდუქტის შეძენის

1. განხორციელდა ტრანზაქცია
ghostchaser -20
ghostbuster +20

2. განახლდა state
```json
  {
    "balances": {
      "ghostchaser": 72,
      "ghostbuster": 28
    },
    "transactions": ["transactionID"]
  }
```


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
