# gb-example [![Build Status](http://img.shields.io/travis/fatih/gb-example.svg?style=flat-square)](https://travis-ci.org/fatih/gb-example)

This is an example Go project which uses [gb](http://getgb.io/). It includes
internal packages, external third-party vendorized packages and travis
integration to show all possible use cases.

Download it and use it to play with the [gb](http://getgb.io/) tool.
Instructions to install `gb` is here: http://getgb.io/docs/install/

# Usage

Fetch the project to a directory (doesn't need to be under `GOPATH` at all!)

```
git clone https://github.com/fatih/gb-example.git
cd gb-example
```

To build the project:

```
$ gb build
```

Execute the newly build command:

```
$ ./bin/convert gbExample
gb_example
```

To run the tests:

```
$ gb test
```

That's it! 

# Details

Examine the directory and see how the source code is laid out in the `src`
folder. Currently, after `gb build` for Mac OSX it's in the form of:

```bash
$ tree
.
├── LICENSE.md
├── README.md
├── bin
│   └── convert
├── pkg
│   └── darwin
│       └── amd64
│           ├── github.com
│           │   └── fatih
│           │       ├── camelcase.a
│           │       └── gb-example
│           │           └── snakecase.a
│           └── notgood.a
├── src
│   ├── cmd
│   │   └── convert
│   │       └── main.go
│   ├── github.com
│   │   └── fatih
│   │       └── gb-example
│   │           └── snakecase
│   │               ├── snakecase.go
│   │               └── snakecase_test.go
│   └── notgood
│       └── notgood.go
└── vendor
    ├── manifest
    └── src
        └── github.com
            └── fatih
                └── camelcase
                    ├── LICENSE.md
                    ├── README.md
                    ├── camelcase.go
                    └── camelcase_test.go
```

* Your code is always under `src/`. The code should be organized in packages. 
* Packages should be in Go style and have proper layout. Here we use
  `github.com/fatih/gb-example` folder and all our packages are inside that
  folder. 
* We have three packages under `src/github.com/fatih/gb-example/`:
	1. `snakecase` is package, which can be imported by others. It uses the third
	   party package `github.com/fatih/camelcase`
	2. `convert` is a `main` package that can be executed once build. It
	   imports `snakecase`.
	3. `notgood` is a package that can be imported as `import "notgood"`. It's
	   a stdlib style package and is not recommended.
* `snakecase` is using a third party package, we fetched and stored it under the
  `vendor/` directory. This is done with `gb fetch github.com/fatih/camelcase`.
* vendor directory doesn't need the `manifest` file. It's up to you have you
  store the dependencies. You can even left it out and have virtual file system
  attached to it on your local server. It doesn't matter for gb.
* Main packages are put under the `bin/` directory once you build the project
  with `gb build`
* We run all tests with `gb test`, which runs the tests for all packages under
  `src/`. 
* We use `travis` for CI integration. Check the `.travis.yml` file out how to
  integrate `gb` into other CI services.
* Having a proper layout under `src/` is also handy, as you can just copy the
  folder `src/github.com/gb-example/fatih` to a `$GOPATH` and it'll compile if
  you have the necessary (and correct) dependencies in your `$GOPATH`.

# Questions/Concerns

* The path `src/github.com/fatih/gb-example` is causing to appeare the path twice because it
  changes the URL path into `https://github.com/fatih/gb-example/tree/master/src/github.com/fatih/gb-example`.
  This is very long and not pleasant to the eyes. A solution would be putting in
  the form of `src/gb-example`, so the code can import in the form of `import "gb-example/snakecase"`.

# Contribute

`gb` is under constant development. Please feel free to contribute to fix or
improve the current layout. If you have a new test case, open a PR and let us
discuss it.

## License

The MIT License (MIT) - see LICENSE.md for more details
