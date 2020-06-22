# SSB Validation Dataset

A small dataset of messages that help you build a validator for Secure
Scuttlebutt messages. Having a validation dataset lowers the barrier to entry
for writing software that validates SSB messages, and helps us identify and
resolve edge-cases where validators differ.

All messages follow the 'ed25519' scheme, sometimes called the 'legacy' or
'classic' encoding of Secure Scuttlebutt.

## API

This module has a file called `data.json` which contains a pretty-printed JSON
array where each array item is an object. Each object has a few entries:

- **`state` (Object | null):** ID and sequence number of previous message, if it exists.
- **`message` (Any):** The message to be inspected, which is just a plain message
  (sometimes called `value` in implementations that use `{ key, value }` to
  represent messages).
- **`valid` (Boolean):** Whether the message is valid.
- **`error` (String | null):** Why the message is invalid, if it is.
- **`hmacKey` (String | null):**": The base64-encoded HMAC key, if it exists.
- **`id` (String | null):** The ID of the message.

## Contributing

If you find a tricky message that might be handled inconsistently by different
validators, please add it to the list. 

## License

AGPL-3.0
