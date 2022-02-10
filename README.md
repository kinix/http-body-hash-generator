# HTTP Body Hash Generator

This CLI (Command Line Interface) tool sends HTTP GET requests and print MD5 hash values of the response's body.

## Usage

You need [Go](https://go.dev/dl/) to build the app.

After build tha app, you can use it in command line. The app get URL list from command line arguments.

Example:
```
$ ./myapp google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com

yandex.com 38edd673423f7dc9fbe919d7ea7b49cb
baroquemusiclibrary.com dadd85ea2820b3fea639660ddc055c6e
facebook.com 2eb639434b15eac3a9e1d8c2dd9dba65
twitter.com 34e80bb64eed21911f3c9d0fcf965a9c
yahoo.com d961c6dafbed56004547c44723683e07
reddit.com/r/notfunny 1567d70393a19bdef283883ab85b7ddf
reddit.com/r/funny 54f3de925270c32fb23c82c014f304ad
google.com 99f21f345907233e72a6cae7256afce5
```

In order to reduce waiting time, the app sends parallel requests. The default count of parallel workers is 10, but you can use `-parallel` flag to change that value.

Example:
```
./myapp -parallel 3 google.com facebook.com yahoo.com yandex.com twitter.com reddit.com/r/funny reddit.com/r/notfunny baroquemusiclibrary.com

facebook.com 3d059183b731490ebf3c4241bed6a2ca
google.com e57eb17df6c4b94f8203b06bd07055a6
yandex.com 8314e5d500f91a88f9931794fdd4c19f
twitter.com e9bcad7f8d2f96ba72888861160dab39
reddit.com/r/funny c9d447483d13cc3a7be5ddeca7d60544
yahoo.com 40c7a55ed6a5c4a17aa6889eca77241b
baroquemusiclibrary.com 990a7d6bf2963d76e7e6b15635e21b11
reddit.com/r/notfunny 08e88e3a61901a47d7361baf9146ebab
```

## Development

If you want to use a different method to get URL list, another hashing algorithm, or printing to file etc, the only thing you need to do is implementing an adapter for it and use it in `main.go`. You don't need to change buisiness logic for these type of changes.

## Unit Tests

There are some unit tests in the project. If you need to add new test cases, you can easily add them into `testCases` array in related test file.
