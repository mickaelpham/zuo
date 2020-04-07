# zuo

Command-line interface to Zuora

## Why the Name?

Because it's Z-**u**-ora, not Z*ou*ra people!

(It's the same as Zen**d**esk, not Zen*D*esk, no capital “d” my friend.)

## Usage

Query a Zuora object knowing its ID

```sh
zuo query account zuora-account-id
```

Execute a ZOQL query

```sh
zuo exec "SELECT Name FROM Product WHERE Status = 'Active'"
```
