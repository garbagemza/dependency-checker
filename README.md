[![cdcheck](https://github.com/garbagemza/dependency-checker/actions/workflows/go.yml/badge.svg)](https://github.com/garbagemza/dependency-checker/actions/workflows/go.yml)

# cdcheck
A dependency checker and resolver for C / C++ repositories.

## dependencies

You need the following software installed on your machine:

* `git` for downloading other dependencies.

## build

`go build -v ./...`


`cdcheck` is created for you.

## run

This tool is intended for use with command line on global scope

Put this binary on your bin/ directory for use.

## usage

1. `cd YourWorkingDir`
2. Create `build.yaml` file. Here's an example of a build.yaml file


```
dependencies:
  - name: rnd
    type: Library
    repository: "https://github.com/garbagemza/rnd.git"
    version: 0.0.1
```

3. dependencies list must contain all objects you want to download.
  * `name` name of the dependency.
  * `type` type of dependency. Library is the only valid type.
  * `repository` full path of the git repository to clone.
  * `version` tag version to download.

4. Save `build.yaml` file.
5. Run `cdcheck`
6. `cdcheck` will create the following directories for you:

  `build` > `dependencies` > `rnd` > `All rnd directories and files from repository's version 0.0.1`
  
