You said:
create readme for this project 
## Mandatory Part
### Validate

The validate feature checks if a credit card number is valid using Luhn's Algorithm.

Requirements:

    The number must be at least 13 digits long.
    If valid, print OK to stdout and exit with status 0.
    If invalid, print INCORRECT to stderr and exit with status 1.
    Support passing multiple entries.
    Support --stdin flag to pass number from stdin.

$ ./creditcard validate "4400430180300003"
OK
$ ./creditcard validate "4400430180300002"
INCORRECT
$ ./creditcard validate "4400430180300003" "4400430180300011"
OK
OK
$ echo "4400430180300003" | ./creditcard validate --stdin
OK
$ echo "4400430180300003" "4400430180300011" | ./creditcard validate --stdin
OK
OK

### Generate

The generate feature creates possible credit card numbers by replacing asterisks (*) with digits.

Requirements:

    Replace up to 4 asterisks (*) with digits. If more - it's an error. Asterisks should be at the end of the given credit card number.
    Print the generated numbers to stdout.
    Numbers must be printed in ascending order.
    Exit with status 1 if there is any error.
    Support --pick flag to randomly pick a single entry.

$ ./creditcard generate "440043018030****"
4400430180300003
4400430180300011
4400430180300029
...
4400430180309988
4400430180309996
$ ./creditcard generate --pick "440043018030****"
4400430180304385

In case of an error:

$ ./creditcard generate --pick "44004301803*****"
$ echo $?
1

### Information

The information feature provides details about the card based on data in brands.txt and issuers.txt.

Requirements:

    Output the card number, validity, brand, and issuer.
    Support --stdin flag to pass number from stdin.
    Support passing multiple entries.

$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
4400430180300003
Correct: yes
Card Brand: VISA
Card Issuer: Kaspi Gold

In case of an incorrect card number:

4400450180300003
Correct: no
Card Brand: -
Card Issuer: -

Example content of brands.txt:

VISA:4
MASTERCARD:51
MASTERCARD:52
MASTERCARD:53
MASTERCARD:54
MASTERCARD:55
AMEX:34
AMEX:37

Example content of issuers.txt:

Kaspi Gold:440043
Forte Black:404243
Forte Blue:517792
Halyk Bonus:440563
Jusan Pay:539545

Notice, that both brands.txt and issuers.txt enumerate different kinds of a card, where each line consists of two parts. First part is the name of the kind and the second is the number prefix (i.e. numbers the card starts with). The parts are separated by : (two dots) symbol.
### Issue

The issue feature generates a random valid credit card number for a specified brand and issuer.

Requirements:

    Pick a random number for the specified brand and issuer.
    Exit with status 1 if there is any error.

$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
4400430180300003

## Support

If you get stuck, try your code with the example inputs from the project. You should get the same results. If not, read the description again. Maybe you missed something, or your code is wrong. After the examples work, but your answer is still wrong, make some test cases you can check by hand. See if they work with your code. Use the complete example input.

If you are still stuck, ask a friend for help or take a break and come back later.
## Guidelines from Author

Learn how to decompose problems

When you're trying to tackle a large project or problem, it's easy to become discouraged by just how big or intimidating the task seems. Instead of being intimidated, work on breaking the problem up (decomposing the problem) into smaller and smaller subproblems, until each individual subproblem seems manageable and doable.

Essentially, you're trying to gain the mindset of an engineer. If a problem is big, don't give up. Keep breaking it up into smaller and smaller pieces, and research things that seem irreducible. Then repeat, continuing the cycle: break down, google, break down, google...

This, by the way, is a pretty effective strategy for self-teaching!
ChatGPT said:
ChatGPT
Credit Card Validator and Generator
Overview

This project provides a command-line tool to validate, generate, and gather information about credit card numbers. It supports various features including validating credit card numbers using Luhn's Algorithm, generating possible card numbers, retrieving brand and issuer information from external files, and issuing random valid card numbers based on brand and issuer. The tool works with the following functionalities:

    Validate a card number.
    Generate card numbers by replacing asterisks with digits.
    Information about a card based on brand and issuer data.
    Issue a random valid card number for a specific brand and issuer.

Table of Contents

    Features
        Validate
        Generate
        Information
        Issue
    Usage Examples
    Requirements
    Support
    Author's Guidelines

Features
1. Validate

The validate feature checks if a credit card number is valid using Luhn's Algorithm.
Requirements:

    The number must be at least 13 digits long.
    If the number is valid, it prints OK and exits with status 0.
    If the number is invalid, it prints INCORRECT and exits with status 1.
    Supports multiple card numbers.
    Supports the --stdin flag to input numbers via standard input.

Example:

$ ./creditcard validate "4400430180300003"
OK

$ ./creditcard validate "4400430180300002"
INCORRECT

$ ./creditcard validate "4400430180300003" "4400430180300011"
OK
OK

$ echo "4400430180300003" | ./creditcard validate --stdin
OK

$ echo "4400430180300003" "4400430180300011" | ./creditcard validate --stdin
OK
OK

2. Generate

The generate feature generates possible credit card numbers by replacing asterisks (*) in the input with digits.
Requirements:

    Replace up to 4 asterisks with digits; any more will result in an error.
    The asterisks should be at the end of the card number.
    Numbers are printed in ascending order.
    Use --pick flag to randomly pick one generated number.

Example:

$ ./creditcard generate "440043018030****"
4400430180300003
4400430180300011
4400430180300029
...
4400430180309988
4400430180309996

$ ./creditcard generate --pick "440043018030****"
4400430180304385

In case of an error (e.g., more than 4 asterisks):

$ ./creditcard generate --pick "44004301803*****"
$ echo $?
1

3. Information

The information feature provides details about the card (validity, brand, and issuer) based on external files (brands.txt and issuers.txt).
Requirements:

    Output the card number, validity, brand, and issuer.
    Supports --stdin flag to pass number(s) via standard input.
    Supports multiple card entries.

Example:

$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
4400430180300003
Correct: yes
Card Brand: VISA
Card Issuer: Kaspi Gold

$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400450180300003"
4400450180300003
Correct: no
Card Brand: -
Card Issuer: -

Example brands.txt:

VISA:4
MASTERCARD:51
MASTERCARD:52
MASTERCARD:53
MASTERCARD:54
MASTERCARD:55
AMEX:34
AMEX:37

Example issuers.txt:

Kaspi Gold:440043
Forte Black:404243
Forte Blue:517792
Halyk Bonus:440563
Jusan Pay:539545

4. Issue

The issue feature generates a random valid credit card number based on a specified brand and issuer.
Requirements:

    Pick a random valid card number for the specified brand and issuer.
    Exits with status 1 if there is an error.

Example:

$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
4400430180300003

Usage Examples
Validate Command

To validate a single card:

$ ./creditcard validate "4400430180300003"
OK

To validate multiple cards:

$ ./creditcard validate "4400430180300003" "4400430180300002"
OK
INCORRECT

To validate from stdin:

$ echo "4400430180300003" | ./creditcard validate --stdin
OK

Generate Command

Generate possible cards from a template:

$ ./creditcard generate "440043018030****"
4400430180300003
4400430180300007
...

Pick a single random card:

$ ./creditcard generate --pick "440043018030****"
4400430180304536

Information Command

Get information about a card:

$ ./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"
4400430180300003
Correct: yes
Card Brand: VISA
Card Issuer: Kaspi Gold

Issue Command

Generate a random valid card number for a specific brand and issuer:

$ ./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"
4400430180300003

Requirements

To run this project, you need the following:

    A C/C++ compiler (GCC, Clang, etc.)
    Basic familiarity with handling file I/O and command-line arguments.
    brands.txt and issuers.txt files containing card brands and issuers data.

Support

If you encounter any issues or need help with the code, refer to the example inputs provided. Ensure your output matches the expected results, and double-check your understanding of the task.

If you’re still stuck, here are some tips:

    Break down the problem into smaller subproblems.
    Test your code with simple cases first.
    Review the project’s specifications and requirements.

If you continue to face issues, reach out to others for help or take a break before trying again.
Author's Guidelines
Decompose Problems

When tackling a complex problem, break it down into smaller, manageable parts. Solving smaller problems one at a time is an effective strategy for success. Research and iteration are key to understanding and solving complex problems.

This approach is not only useful for the project at hand but is also a great way to improve your problem-solving skills for future projects. Keep breaking down, researching, and iterating until you reach a solution!


ChatGPT can make mistakes. Check important info.
