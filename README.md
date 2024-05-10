# SMTP Example

## Brief

This is a simple example of how to send an email using the `net/smtp` package in Go. The example reads the SMTP server host, port, sender email address, and password from environment variables. The email recipients are specified as command-line arguments.

## How to run

1. Build the executable

    ```Console
    go build -o smtp-example
    ```

1. Run the executable

    ```Console
    ./smtp-example -h
    ```

## How to use

```Console
$ ./smtp-example -h
usage: ./smtp-example [-h|--help] -r|--recipient "<value>" [-r|--recipient
                      "<value>" ...]

                      This is a simple example of how to send an email using
                      the `net/smtp` package in Go. The example reads the SMTP
                      server host, port, sender email address, and password
                      from environment variables. The email recipients are
                      specified as command-line arguments.

Arguments:

  -h  --help       Print help information
  -r  --recipient  Recipient email address; can be specified multiple times
```

Several environment variables are required to be set in order to send an email. These are:

- `HOST`: The SMTP server host, e.g. `smtp.gmail.com`.
- `PORT`: The SMTP server port, e.g. the port for Gmail when using TLS is `587`.
- `SENDER`: The email address of the sender.
- `PASSWORD`: The password of the sender. For Gmail, you need to [generate an app password](https://support.google.com/accounts/answer/185833).

You can set these environment variables in a `.env` file in the root of the project. Rename the `sample.env` file to `.env` and fill in the values.

## Known limitations

- The recipients are set as blind carbon copy (BCC) recipients.
