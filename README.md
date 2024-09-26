# upgen

The Universal Password GENerator (`upgen`).

`upgen` can generates passwords with _at least_ the specified amount of entropy.
The default is 100 bits.

## Usage

```go install github.com/49pctber/upgen/cmd/...@latest```

This will install the `upgen` utility.

Specify the amount of entropy you would like with the `-e` flag, and the number of passwords to generate with `-n`.
Use `-h` for help.

Specify the types of characters/words to use using the following flags:

- `--alphanum`: Include alphanumeric characters
- `--b32`:      Include base32 characters
- `--b64`:      Include base64 characters
- `--bip39`:    Include BIP-39 words
- `--hex`:      Include hex characters
- `--letter`:   Include letter characters
- `--lower`:    Include lowercase characters
- `--num`:      Include numeric characters
- `--special`:  Include special characters
- `--upper`:    Include uppercase characters
