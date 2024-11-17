# Credit Card Validator and Generator

## Overview

This tool provides a set of command-line utilities for validating, generating, and retrieving information about credit card numbers. It supports:

- **Validating** card numbers using Luhn’s Algorithm.
- **Generating** card numbers by replacing asterisks with digits.
- **Providing card information** based on external brand and issuer files.
- **Issuing** random valid credit card numbers for specific brands and issuers.

## Features

### 1. Validate

Validates a credit card number using **Luhn's Algorithm**.

#### Requirements:
- The number must be **at least 13 digits** long.
- Prints `OK` if the number is valid, or `INCORRECT` if invalid.
- Supports multiple card numbers and `--stdin` flag for input via stdin.

#### Example:
```bash
$ ./creditcard validate "4400430180300003"
OK

$ ./creditcard validate "4400430180300002"
INCORRECT
```

### 2. Generate

Generates credit card numbers by replacing **up to 4 asterisks** at the end of the number with digits.

#### Requirements:
- Asterisks must be **at the end** of the number.
- Numbers are printed in ascending order.
- The `--pick` flag randomly picks one generated number.

#### Example:
```bash
$ ./creditcard generate "440043018030****"
4400430180300003
4400430180300007
...

$ ./creditcard generate --pick "440043018030****"
4400430180304385
```

### 3. Information

Retrieves detailed information about a card (brand, issuer, validity) using external data files (`brands.txt` and `issuers.txt`).

#### Requirements:
- Supports `--stdin` to input card numbers via stdin.
- Looks up brand and issuer based on card number prefixes.

#### Example:
```bash
$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
4400430180300003
Correct: yes
Card Brand: VISA
Card Issuer: Kaspi Gold
```

### 4. Issue

Generates a **random valid card number** for a specific brand and issuer.

#### Example:
```bash
$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
4400430180300003
```

## Usage

### Validate Command
```bash
$ ./creditcard validate <card_number>
```
Example:
```bash
$ ./creditcard validate "4400430180300003"
OK
```

### Generate Command
```bash
$ ./creditcard generate <card_template>
```
Example:
```bash
$ ./creditcard generate "440043018030****"
4400430180300003
4400430180300007
...
```

### Information Command
```bash
$ ./creditcard information --brands=<brands_file> --issuers=<issuers_file> <card_number>
```
Example:
```bash
$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
```

### Issue Command
```bash
$ ./creditcard issue --brands=<brands_file> --issuers=<issuers_file> --brand=<brand> --issuer=<issuer>
```
Example:
```bash
$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
4400430180300003
```

## Files

- **brands.txt**: Contains card brand prefixes (e.g., `VISA:4`, `MASTERCARD:51`).
- **issuers.txt**: Contains issuer names and prefixes (e.g., `Kaspi Gold:440043`).

## Requirements

- A C/C++ compiler (GCC, Clang, etc.)
- Files: `brands.txt` and `issuers.txt` for card brand and issuer information.

## Support

For issues or troubleshooting:
- Test with the provided example inputs.
- Review the tool’s error messages for guidance.
- If still stuck, try breaking down the problem into smaller steps or seek help from others.

