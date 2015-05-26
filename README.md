# gb-example [![Build Status](http://img.shields.io/travis/fatih/gb-example.svg?style=flat-square)](https://travis-ci.org/fatih/gb-example)

This is an example Go project which uses [gb](http://getgb.io/). It includes
internal packages, external third-party vendorized packages and travis
integration to show all possible use cases.

Download it and use it to play with the [gb](http://getgb.io/) tool.

# Usage

```sh
git clone https://github.com/fatih/gb-example.git
```

To build the project with `gb`

```sh
gb build all
```
Execute the newly build command:

```sh
$ ./bin/cmdname
```
To run the tests:

```sh
gb test all
```

That's it! 

# Details

Examine the directory and see how the source code is laid out in the `src`
folder. Currently, after `gb build all` it's for Mac OSX it's in the form of:

```bash
~/Code/gb-example (master) tree
.
├── LICENSE.md
├── README.md
├── bin
│   └── cmdname
├── pkg
│   └── darwin
│       └── amd64
│           ├── github.com
│           │   └── fatih
│           │       └── camelcase.a
│           └── util.a
├── src
│   ├── cmdname
│   │   └── main.go
│   └── util
│       ├── util.go
│       └── util_test.go
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
* We have two packages under `src/`:
	1. `util` is package, which can be imported by others. It uses the third party package `github.com/fatih/camelcase`
    2. `cmdname` is a `main` package that can be executed once build. It imports `util`.
* `util` is using a third party package, we fetched and stored it under the `vendor/` directory. This is done with `gb fetch github.com/fatih/camelcase`.
* `gb` put main packages under the `bin/` directory once you build the project with `gb build`
* `gb test all` runs the tests for all packages under `src/`. You'll see `testing: warning: no tests to run` because there are not tests under `src/cmdname`.
* `travis` is used for CI integration. Check the `.travis.yml` file out how to integrate `gb` into other CI services.

# Contribute

`gb` is under constant development. Please feel free to contribute to fix or
improve the current layout. If you have a new test case, open a PR and let us
discuss it.

## License

The MIT License (MIT) - see LICENSE.md for more details
