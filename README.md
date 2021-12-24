# smtpoauth2
Go net/smtp Auth provider for oauth2

## Usage
To get started you will need a token.json. The easiest way to generate that is with the googleworkshop/go-samples gmail api quickstart tool. I have included a modified version in this repository that includes the appropriate scope for email sending under `cmd/quickstart` 

After that just `import "github.com/jacobalberty/smtpoauth2"` and use `smtpoauth2.Oauth2(user, tokenType, token)` to return an auth object suitable for net/smtp. tokenType will most likely be `Bearer` and token is the `access_token` from your previosly generated `token.json`.

There is sample code in `cmd/sendmail` that reads `token.json` and will send an email using this auth provider.
