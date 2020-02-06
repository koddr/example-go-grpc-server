# Example gRPC server on Go

This is full project structure example of article **[Enter to gRPC with Go: first server]()**.

Published on [Dev.to]() @ 07 Feb 2020.

## Requirements

- Go `1.11+`
- Go Modules

## Usage

- Clone this repository and go to folder:

```console
git clone https://github.com/koddr/example-go-grpc-server.git
cd example-go-grpc-server
```

- Try to build binary:

```console
make build
```

- No errors? Let's run:

```console
make run
```

- Go to another console session and connect to gRPC server with [Evans](https://github.com/ktr0731/evans) (gRPC client):

```console
evans api/proto/adder.proto -p 8080
```

- That's all!

### Live demo

[![asciicast](https://asciinema.org/a/298722.svg)](https://asciinema.org/a/298722)

## Author

- [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Article assistance

If you want to say «thank you»:

1. Twit about article [on your Twitter](https://twitter.com/intent/tweet?text=).
2. Add a GitHub Star and make Fork to this repository.
3. Donate some money to project author via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).
4. Join DigitalOcean at our [referral link](https://m.do.co/c/b41859fa9b6e) (your profit is **\$100** and we get \$25).

Thanks for your support! 😘

## License

MIT
