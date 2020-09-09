# Exercise 4.11

## Description

Build a tool that lets users create, read, update and close GitHub issues from command line, invoking their preferred text editor when substantial text input is required.


## Usage

`go build -o issuer` // compile

### Create an issue

Export Github credentials as `GITHUB_USERNAME` and `GITHUB_PASSWORD`

Run `./issuer create -owner OWNER -title TITLE -repo REPO [-assignees ASSIGNEES] [-milestone MILESTONE] [-labels LABELS] [-editor vi|vim|nano]`

### Update an issue

Export Github credentials as `GITHUB_USERNAME` and `GITHUB_PASSWORD`

Run `./issuer update -owner OWNER -id ID -repo REPO [-title TITLE] [-assignees ASSIGNEES] [-milestone MILESTONE] [-labels LABELS] [-editor vi|vim|nano]`

### Close an issue

Export Github credentials as `GITHUB_USERNAME` and `GITHUB_PASSWORD`

Run `./issuer close -id ID -owner OWNER -repo REPO`

### Read an issue

Run `./issuer read -owner OWNER -repo REPO -id ID`








