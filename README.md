# td

Telegram client implementation in go.

* [Security policy](.github/SECURITY.md)
* [Support](.github/SUPPORT.md)

## Status

Work in progress. Only go1.15 is supported and no backward compatibility is provided.

Goal of this project is to implement Telegram client while
providing building blocks for the other client or even server
implementation without performance bottlenecks.

## Features

* Full MTProto 2.0 implementation in go
* Code for Telegram types automatically generated by `./cmd/gotdgen` (based on [ernado/tl](https://github.com/ernado/tl) parser)
* Pluggable session storage
* Vendored Telegram public keys that are kept up-to-date
* Testing and fuzzing applied to improve durability

## Example

You can see `cmd/gotdecho` for echo bot example.

## Reference

The MTProto protocol description is [hosted](https://core.telegram.org/mtproto#general-description) by Telegram.

Most important parts for client impelemtations:
* [Security guidelines](https://core.telegram.org/mtproto/security_guidelines) for client software developers

Current implementation does not conform to security guidelines and should be used only
as reference or for testing.

## Prior art

* [Lonami/grammers](https://github.com/Lonami/grammers) (Great Telegram client in Rust, many test vectors were used as reference)
* [sdidyk/mtproto](https://github.com/sdidyk/mtproto), [cjongseok/mtproto](https://github.com/cjongseok/mtproto), [xelaj/mtproto](https://github.com/xelaj/mtproto)  (MTProto 1.0 in go)

## License

MIT License, created by Aleksandr (ernado) Razumov in 2020
