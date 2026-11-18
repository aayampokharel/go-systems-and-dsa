# What I learned

- Length-prefixed encoding (`len#value`) is unambiguous: the decoder reads the number, then jumps exactly that many bytes.
- It works for empty strings, `#` inside values, and any ASCII/UTF-8 content, because `#` is never searched inside the value.
- A naive delimiter join/split would break as soon as a string contains the delimiter.
- `Encode` currently builds the result with `+=`; using `strings.Builder` would avoid repeated string copying for large inputs.
