# zuo

Command-line interface to Zuora

## Why the Name?

Because it's z-**u**-ora, not z*ou*ra people!

It's the same as Zendesk, not Zen*D*esk, no capital “d” my friend.

## Usage

Query a Zuora object knowing its ID

```sh
zuo query account zuora-account-id
```

Execute a ZOQL query

```sh
zuo exec "SELECT Name FROM Product"
```
