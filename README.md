# zuo

Command-line interface to Zuora

## Usage

Query a Zuora object knowing its ID

```sh
zuo query account zuora-account-id
```

Execute a ZOQL query

```sh
zuo exec "SELECT Name FROM Product WHERE Status = 'Active'"
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
