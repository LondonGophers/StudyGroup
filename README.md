# Go London Study Group

<p align="center">
  <img width="209" height="300" src="https://raw.githubusercontent.com/go-london-user-group/study-group/master/resources/gopher-study.png">
</p>

```
func main() {
	fmt.Println("Hello, London Gophers!")
}
```

## Code of Conduct

We follow the [Go Code of Conduct](https://golang.org/conduct) and the [Contributor Covenant](https://github.com/go-london-user-group/study-group/blob/master/CONTRIBUTING.md), in summary:

- Treat everyone with respect and kindness.
- Be thoughtful in how you communicate.
- Don’t be destructive or inflammatory.

## The Workshop

Visit the [Go London Study Group Meetup](https://www.meetup.com/Go-London-Study-Group/) for details of our IRL study group meetings.

Find us on Slack in the `#london-study-group` channel of the `gophers.slack.com` workspace.

Our study material is:
[The Go Programming Language by Alan Donovan and Brian W. Kernighan](https://www.gopl.io/)

Work through the material at a speed that suits you. We also have _open space_
meetups to meet to chat about Go and the book etc. So just join in any
time! :)

## 1. Begin

To begin, create your own directory in the [workspaces](workspaces) directory,
work through the book exercises, and add your code there. When committing to the 
repo just commit directly to `master` but before you do so make sure you run 
`git pull -r` (`-r` to rebase) before you push, to avoid _merge commits_. It may be 
easier to instead run:
```
git config branch.autosetuprebase always
```
...in the repository directory to ensure that git pull always performs a rebase, 
rather than a merge, on `git pull`.

To do this for all repositories (a global setting) use:
```
git config --global branch.autosetuprebase always
```

You can add your code in separate directories if you're using 
[Go Modules](#b-running-go-with-go-modules) or if you're 
[using a GOPATH](#a-running-go-with-a-gopath) you could structure your code as 
in the [GOPATH Project Structure](#gopath-project-structure) section.

If you need the exercises from the book, they're available in
[exercises.tar.gz](exercises.tar.gz) zipped archive but these are already in the
book so you might not need them.

To obtain all of the source code examples used in the book (many of which are code 
examples to be amended) you can get this by running:
```
git clone https://github.com/adonovan/gopl.io.git
```

## 2. Installing Go

###  a. Home/Linuxbrew

It may be convenient to install the latest version of Go through the
[Homebrew](https://brew.sh/) and [Linuxbrew](http://linuxbrew.sh/) package
managers.

```
brew install go
```

### b. Install with Binary Distributions

The [https://golang.org/dl/](https://golang.org/dl/) page contains distros for
Windows, MacOS, Linux, and source. The
[installation instructions](https://golang.org/doc/install) explains how to
install them.

## 3. Getting the Workshop Code

You can grab the code samples used in the book (which you update for a bunch of
the exercises) by running such as:

```
go get gopl.io/ch1/helloworld
```

(This will get the `helloworld` code, plus the other examples).

## 4. Building and Running Go

There are two main options, by setting a `GOPATH` and having all of your code in one
location or using _Go Modules_ which allow you to split your code into separate
locations. It might be easier to use a GOPATH pointing to all of your workshop
code but you might find modules more appropriate.

### a. Running Go with Go Modules)

_Go modules_ are available in Go 1.11 onwards which removes the need to
have a `$GOPATH` set.

To use modules in your project directory, run:

```
go mod init $MODULE
```
...where `$MODULE` is the name of your module. For these exercises this could be something 
simple but in production code would be a domain name and path, such as 
`github.com/somebody/component`. This will create a file called `go.mod`.

To include an external dependency to your project just add the import, such as:
```
package main

import "github.com/some/dependency"

func main() {
    dependency.f()
}
```

When running your code this module should be automatically downloaded and added to your
`go.mod` file which could look like:
```
module github.com/somebody/component

require (
    github.com/some/dependency v1.2.3
)
```

Because we have a number of separate components in one place we probably want to run things 
individually rather than all samples together. Something like this will run one exercise:

```
go run ./$MODULE/ch1/ex1_1
```
...assuming your code is in the path `./$MODULE/ch1/ex1_1` underneath your `go.mod` file.

### b. Running Go with a GOPATH

The traditional way to run/build Go code (prior to _Go Modules_) is using a
GOPATH. You should set your `$GOPATH` to your current directory, such as:
```
export GOPATH=/home/gopherg/eng-golang-workshop/workspaces/gogopher
```

To run some code you can then use:
```
go run gopl.io/ch1/helloworld
```
_(which actually builds then executes `/home/gopherg/eng-golang-workshop/workspaces/gogopher/src/gopl.io/ch1/helloworld/main.go`)_

To build it and output in your `$GOPATH\bin` directory:
```
go build -o $GOPATH/bin/helloworld gopl.io/ch1/helloworld
```

To get another module (such as the imaginary `some/dependency`):
```
go get github.com/some/dependency
```
...and this will then be downloaded to `$GOPATH/src/github.com/some/dependency` and
imported with `import "github.com/some/dependency"`

#### GOPATH Project Structure

When using GOPATH your project structure may be something like:
```
workspaces
    gogopher\
        bin\
            helloworld
            ...
        src\
            gogopher.io\
                ch1\
                    ex1_1\
                        main.go
                ...
```

## 5. Development Environments

### a. Delve (Debugger)

[Delve](https://github.com/derekparker/delve) is is a debugger for Go. To install run:
```
go get -u github.com/derekparker/delve/cmd/dlv
```
To see the available commands, run `dlv` then `help` at the `(dlv)` prompt.


### b. GoLand

GoLand is a new commercial [Go IDE by JetBrains]((https://www.jetbrains.com/go/))
aimed at providing an ergonomic environment for Go development.

The new IDE extends the IntelliJ platform with coding assistance and tool
integrations specific for the Go language.


### c. Emacs

If you follow similar instructions to get go support for emacs (OS X) as below
http://tleyden.github.io/blog/2014/05/22/configure-emacs-as-a-go-editor-from-scratch/

And you run into the following error when trying to get auto-complete to work.

```
Error running timer ‘ac-update-greedy’: (file-missing "Searching for program" "No such file or directory" "gocode")
Error running timer ‘ac-show-menu’: (file-missing "Searching for program" "No such file or directory" "gocode")
Error running timer ‘ac-update-greedy’: (file-missing "Searching for program" "No such file or directory" "gocode"
```

Then the problem is probably down to `gocode` not being available in your path:
https://emacs.stackexchange.com/questions/10722/emacs-and-command-line-path-disagreements-on-osx

So if you edit `/etc/paths.d/go` and add the path to the bin directory of your project it should fix problem.

### d. Visual Studio Code

Visual Studio Code is a lightweight but powerful source code editor with support for debugging, embedded Git control, syntax highlighting, intelligent code completion, snippets, and code refactoring. Visual Studio Code is based on Electron and uses the Blink layout engine. Vscode uses the same editor component as [Atom](https://github.com/go-london-user-group/study-group#f-atom) (codenamed "Monaco").

The [Go extension for Visual Studio Code](https://code.visualstudio.com/docs/languages/go), provides language features such as IntelliSense, code navigation, symbol search, bracket matching, snippets etc.

### e. acme

There are three kinds of IDEs:
* A character driven IDE such as unix, emacs or vi
* A closed environment with its own bespoke tooling such as eclipse, visual code, intellij, atom, GoLand
* An integrating environment that integrates tools from outside inwards such as plan9 and acme.

#### Vague Installation Guide

* [acme](https://ilanpillemer.github.io/subjects/acme.html)

### f. Atom

Atom supports Go development with the
[go-plus](https://atom.io/packages/go-plus) package.

To use Delve inside Atom, install the
[go-debug](https://atom.io/packages/go-debug) package.

To run your Go code in Atom, install the
[atom-runner](https://atom.io/packages/atom-runner) package.

## 6. Links

[A Tour of Go](https://tour.golang.org/welcome/1)

[How to Write Go Code](https://golang.org/doc/code.html)

[Effective Go](https://golang.org/doc/effective_go.html)

[Source code: The Go Programming Language](https://github.com/adonovan/gopl.io)

[YouTube: Concurrency is not Parallelism by Rob Pike](https://www.youtube.com/watch?v=oV9rvDllKEg)

[YouTube: Go Proverbs](https://www.youtube.com/watch?v=PAAkCSZUG1c)

## 7. Rights

All exercises from The Go Programming Language are copyright 2016 Alan A. A. Donovan & Brian W. Kernighan and included with permission from the authors.

All submitted code is covered under [Apache License 2.0](LICENSE).
