# What I learned

- A Go array like `[26]byte` is comparable, so it can be used directly as a map key.
- Counting characters is enough to identify an anagram group; sorting is not required.
- The input constraint says lowercase English letters only, which makes the fixed 26-size count safe.
