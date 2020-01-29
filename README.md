# London Gophers Study Group

<p align="center">
  <img width="239" height="300" src="https://raw.githubusercontent.com/LondonGophers/StudyGroup/master/resources/study-gopher.png">
</p>

``` go
func main() {
    fmt.Println("Hello, London Gophers!")
}
```

## Code of Conduct

We follow the [Go Code of Conduct](https://golang.org/conduct) adapted from the
[Contributor Covenant], version 1.4; in summary:

- Treat everyone with respect and kindness.
- Be thoughtful in how you communicate.
- Don’t be destructive or inflammatory.

From [GopherCon 2018: Natalie Pistunovich - The Importance of Beginners][yt-1]:

_"No question is stupid. When you're [...] answering beginners' questions, do_
_it well, be nice, be informative, and mainly be patient. If you see a_
_question for the hundredth time, you don't have to answer. Answers like 'just_
_use the standard library' are not helpful and they're not friendly. Somebody_
_else can do this who is maybe a bit more fresh."_

## The Workshop

Visit the [London Gophers Study Group Meetup][meetup] for details of our IRL
study group meetings.

Find us on Slack in the `#london-study-group` channel of the
`gophers.slack.com` workspace. To join this Slack workspace, start
[here][slack].

Our study material is:
[The Go Programming Language by Alan Donovan and Brian W. Kernighan][gopl]

Work through the material at a speed that suits you. We also have _open space_
meetups to meet to chat about Go and the book etc. So just join in any time! :)

## 1. Begin

### Just tell me how to write Go code

There's a great how to write Go code document here -> [golang.org/doc/code.html](https://golang.org/doc/code.html) :)

### Ok, the workshop...

To get access to the repo ask an organiser at one of the meetups to add you. 
Then to begin, create your own directory in the [workspaces](workspaces) 
directory, work through the book exercises, and add your code there. When 
committing to the repo just commit directly to `master` but before you do 
so make sure you run `git pull -r` (`-r` to rebase) before you push, to 
avoid _merge commits_.

It may be easier to instead run:

``` shell
git config branch.autosetuprebase always
```

...in the repository directory to ensure that git pull always performs a
rebase, rather than a merge, on `git pull`.

To do this for all repositories (a global setting) use:

``` shell
git config --global branch.autosetuprebase always
```

You can add your code in separate directories if you're using
[Go Modules](#a-running-go-with-go-modules) or if you're
[using a GOPATH](#b-running-go-with-a-gopath) you could structure your code as
in the [GOPATH Project Structure](#gopath-project-structure) section.

If you need the exercises from the book, they're available in
[exercises.tar.gz](exercises.tar.gz) zipped archive but these are already in
the book so you might not need them.

## 2. Installing Go

### a. Home/Linuxbrew

It may be convenient to install the latest version of Go through the
[Homebrew](https://brew.sh/) and [Linuxbrew](http://linuxbrew.sh/) package
managers.

``` shell
brew install go
```

### b. Install with Binary Distributions

The [https://golang.org/dl/](https://golang.org/dl/) page contains distros for
Windows, MacOS, Linux, and source.

The [installation instructions](https://golang.org/doc/install) explain how to
install them.

## 3. Getting the Workshop Code

To obtain all of the source code examples used in the book (many of which are
code examples to be amended) you can get this by running in your workspace:

``` shell
git clone https://github.com/adonovan/gopl.io.git
```

(This will get the `helloworld` code, plus all the other examples).

Remember to add a `.gitignore` file containing:

```
gopl.io
```

...so you don't check in that example code to the repo.

## 4. Building and Running Go

There are two main options, by using _Go Modules_ or by setting a `GOPATH`. You
may wish to use a GOPATH pointing to all of your workspace code but you might
find modules more appropriate and this is now the **recommended** approach.

### a. Running Go with Go Modules

_Go modules_ are available in Go 1.11 onwards which removes the need to have a
`$GOPATH` set.

To use modules in your project directory, run:

``` shell
go mod init MODULE
```

...where `MODULE` is the name of your module. For these exercises this could
be something simple such as `studygroup` but could be a domain name and path
such as `github.com/somebody/component`. This will create a file called
`go.mod`.

To include an external dependency to your project just add the import, such as:

``` go
package main

import "github.com/some/dependency"

func main() {
    dependency.f()
}
```

When running your code this module should be automatically downloaded and added
to your `go.mod` file which could look like:

``` go
module github.com/somebody/component

require (
    github.com/some/dependency v1.2.3
)
```

Because we have a number of separate components in one place we probably want
to run things individually rather than all samples together. Something like
this will run one exercise:

``` shell
go run ./$MODULE/ch1/ex1_1
```

...assuming your code is in the path `./$MODULE/ch1/ex1_1` underneath your
`go.mod` file.

### b. Running Go with a GOPATH

The traditional way to run/build Go code (prior to _Go Modules_) is using a
GOPATH. You should set your `$GOPATH` to your current directory, such as:

``` shell
export GOPATH=/home/gopherg/eng-golang-workshop/workspaces/gogopher
```

To run some code you can then use:

``` shell
go run gopl.io/ch1/helloworld
```

_(which actually builds then executes_
_`/home/gopherg/eng-golang-workshop/workspaces/gogopher/src/gopl.io/ch1/helloworld/main.go`)_

To build it and output in your `$GOPATH\bin` directory:

``` shell
go build -o $GOPATH/bin/helloworld gopl.io/ch1/helloworld
```

To get another module (such as the imaginary `some/dependency`):

``` shell
go get github.com/some/dependency
```

...and this will then be downloaded to `$GOPATH/src/github.com/some/dependency`
and imported with `import "github.com/some/dependency"`

#### GOPATH Project Structure

When using GOPATH your project structure may be something like:

``` text
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

[Delve](https://github.com/derekparker/delve) is is a debugger for Go. To
install, run:

``` shell
go get -u github.com/go-delve/delve/cmd/dlv
```

To see the available commands, run `dlv` then `help` at the `(dlv)` prompt.

### b. GoLand

GoLand is a new commercial [Go IDE by JetBrains][GoLand] aimed at providing an
ergonomic environment for Go development.

The new IDE extends the IntelliJ platform with coding assistance and tool
integrations specific for the Go language.

### c. Emacs

If you follow similar instructions to get Go support for emacs (OS X) from
[here][go-emacs] and you run into the following error when trying to get
auto-complete to work:

``` shell
Error running timer ‘ac-update-greedy’: (file-missing "Searching for program" "No such file or directory" "gocode")
Error running timer ‘ac-show-menu’: (file-missing "Searching for program" "No such file or directory" "gocode")
Error running timer ‘ac-update-greedy’: (file-missing "Searching for program" "No such file or directory" "gocode"
```

...then the problem is probably down to `gocode` not being
[available in your path][emacs-path].

So if you edit `/etc/paths.d/go` and add the path to the bin directory of your
project it should fix the problem.

### d. Visual Studio Code

Visual Studio Code is a lightweight but powerful source code editor with
support for debugging, embedded Git control, syntax highlighting, intelligent
code completion, snippets, and code refactoring. Visual Studio Code is based on
Electron and uses the Blink layout engine.

VSCode uses the same editor component as [Atom](#f-atom) (codenamed "Monaco").

The [Go extension for Visual Studio Code][vscode-go] provides language features
such as IntelliSense, code navigation, symbol search, bracket matching,
snippets etc.

### e. acme

There are three kinds of IDEs:

- A character driven IDE such as unix, emacs or vi
- A closed environment with its own bespoke tooling such as Eclipse, Visual
  Studio Code, IntelliJ, Atom, GoLand
- An integrating environment that integrates tools from outside inwards such as
  plan9 and acme.

#### Vague Installation Guide

- [acme](https://ilanpillemer.github.io/subjects/acme.html)

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

[YouTube: Concurrency is not Parallelism by Rob Pike][yt-2]

[YouTube: Go Proverbs](https://www.youtube.com/watch?v=PAAkCSZUG1c)

## 7. Rights

All exercises from The Go Programming Language are copyright 2016 Alan A. A.
Donovan & Brian W. Kernighan and included with permission from the authors.

All submitted code is covered under [Apache License 2.0](LICENSE).

[Contributor Covenant]: https://github.com/LondonGophers/StudyGroup/blob/master/CONTRIBUTING.md
[emacs-path]: https://emacs.stackexchange.com/questions/10722/emacs-and-command-line-path-disagreements-on-osx
[go-emacs]: http://tleyden.github.io/blog/2014/05/22/configure-emacs-as-a-go-editor-from-scratch/
[GoLand]: https://www.jetbrains.com/go/
[gopl]: https://www.gopl.io
[meetup]: https://www.meetup.com/LondonGophersStudyGroup/
[slack]: https://invite.slack.golangbridge.org
[vscode-go]: https://code.visualstudio.com/docs/languages/go
[yt-1]: https://www.youtube.com/watch?v=7yMXs9TRvVI
[yt-2]: https://www.youtube.com/watch?v=oV9rvDllKEg
