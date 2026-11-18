# Why the length-prefix approach is better

- **Handles any characters:** values may contain `#`, commas, spaces, or other delimiters. Since the decoder trusts the length instead of searching for a delimiter, those characters are never ambiguous.
- **Handles empty strings:** `0#` decodes correctly, while delimiter-based approaches often lose or merge empty entries.
- **No escaping needed:** there is no need to escape or scan for special characters.
- **Simple and linear:** `encode` and `decode` both run in **O(total characters)** time, with `decode` skipping each value in one jump.
- **Generalized:** it satisfies the follow-up because it does not depend on the valid character set.

The main improvement left is using `strings.Builder` in `Encode` instead of repeated `+=` concatenation.
