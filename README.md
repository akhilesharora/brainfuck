# Brainfuck App

Simple interpreter for [brainfuck](https://en.wikipedia.org/wiki/Brainfuck) language

## Usage

### Build

Run `make build`. It will create a binary file in `build` directory

### Tests

Run `make test`. It will run all the tests

To test it with test payload, pass the filename as an arg to the binary: `./build/app test/helloworld.bf`
or `make run-example-payload`

### Help

Run `make help` for different commands


### Scope of improvement

* Error handling needs to be improved
* More test cases could be added
