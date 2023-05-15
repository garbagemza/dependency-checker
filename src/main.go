/**
  * This is free and unencumbered software released into the public domain.
  *
  * Anyone is free to copy, modify, publish, use, compile, sell, or
  * distribute this software, either in source code form or as a compiled
  * binary, for any purpose, commercial or non-commercial, and by any
  * means.
  *
  * In jurisdictions that recognize copyright laws, the author or authors
  * of this software dedicate any and all copyright interest in the
  * software to the public domain. We make this dedication for the benefit
  * of the public at large and to the detriment of our heirs and
  * successors. We intend this dedication to be an overt act of
  * relinquishment in perpetuity of all present and future rights to this
  * software under copyright law.
  *
  * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
  * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
  * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
  * IN NO EVENT SHALL THE AUTHORS BE LIABLE FOR ANY CLAIM, DAMAGES OR
  * OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE,
  * ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
  * OTHER DEALINGS IN THE SOFTWARE.
  *
  * For more information, please refer to <https://unlicense.org>
  **/


package main

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "gopkg.in/yaml.v3"
)

type Dependency struct {
  Name       string
	Type       string
  Repository string
	Version    string
}

type Dependencies struct {
  Dependencies []Dependency
}

func main() {
	// Load the file; returns []byte
	f, err := os.ReadFile("build.yaml")
	if err != nil {
		log.Fatal(err)
	}

	var dependencies Dependencies
	if err := yaml.Unmarshal(f, &dependencies); err != nil {
		log.Fatal(err)
	}

  const workDirectory = "./build/dependencies"

  printGitVersion()
  createDirectory(workDirectory)
  fetchDependencies(dependencies.Dependencies, workDirectory)
}

func printGitVersion() {
  cmd := exec.Command(
    "git",
    "--version")
    out, _ := cmd.CombinedOutput()
    fmt.Printf("%+v\n", string(out))
}

func createDirectory(path string)  {
  err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}

func fetchDependencies(dependencies []Dependency, directory string) {
  for _, d := range dependencies {
    fetchDependency(d, directory)
  }
}

func fetchDependency(dependency Dependency, directory string)  {
  fmt.Printf("%+v\n", dependency.Name)

  cmd := exec.Command(
    "git",
    "clone",
    "--branch",
    dependency.Version,
    "--depth",
    "1",
    dependency.Repository)

  cmd.Dir = directory
  out, err := cmd.CombinedOutput()
  if err != nil {
    fmt.Printf("%+v\n", string(out))
    log.Fatal(err)
  }
}

