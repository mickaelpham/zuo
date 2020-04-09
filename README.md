# zuo

Command-line interface to Zuora

## Usage

### Query a Zuora object knowing its ID

_This feature is not (yet) implemented._

```sh
zuo query account zuora-account-id
```

### Execute a ZOQL query

Standard usage will print out data in a table format.

```sh
zuo exec "SELECT Name FROM Product WHERE Status = 'Active'"
```

ZOQL query string is case-insensitive.

```sh
zuo exec "select name from product where status='active'"
```

If you don't use any `WHERE` clause with a `String` or `Date` argument, you can
use single quotes.

```sh
zuo exec 'select name from product'
```

You can use the `--json` flag to print the raw JSON response from Zuora. It's
useful when piping the output to a command-line JSON processor such as
[`jq`](https://stedolan.github.io/jq/).

```sh
zuo exec --json "SELECT Id FROM Account WHERE Balance > 0" | jq '.size'
```

## Setup

1. Download the [latest release][latest-release]
2. Move the executable to `/usr/local/bin`

```sh
mv ~/Downloads/zuo /usr/local/bin
```

3. Create an [OAuth client][kc-oauth] in the Zuora UI
4. Add the client ID and secret to your environment

```sh
# if using ZSH, put this in your ~/.zshrc
export ZUO_CLIENT_ID="my-client-id"
export ZUO_CLIENT_SECRET="my-client-secret"
export ZUO_BASE_URL="https://rest.apisandbox.zuora.com"
```

## Why the Name?

Because it's Z-**u**-ora, not Z*ou*ra people!

(It's the same as Zen**d**esk, not Zen*D*esk, no capital “d” my friend.)

[latest-release]: https://github.com/mickaelpham/zuo/releases
[kc-oauth]:
  https://knowledgecenter.zuora.com/Billing/Tenant_Management/A_Administrator_Settings/Manage_Users#Create_an_OAuth_Client_for_a_User
