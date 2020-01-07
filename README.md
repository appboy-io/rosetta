# Rosetta: Alias Context Switcher

Alias manager for the command line. Separated by context/workspace. Often working in full stack development, developers end up with large alias files and finding, editing, and adding new commands can be a pain. Rosetta is here to fix that.

## Pre-Requisites

- Golang

## Features

## Links

## Usage

On intial install of Rosetta, you can have it set up default context alias files using:

```bash
rose init
```

This will make the default context alias files. For example,

- .rose_devops
- .rose_frontend
- .rose_backend

### Switching Context

Switching context is done with the following command:

```bash
rose -c devops
```

This will switch your context to devops. Your current terminal session will have the devops aliases.
