# LSP

Language server protocol library/framework for Golang 1.21+

> :warning: **This project is in early development.** The API is subject to
> change.

This package simplifies the process of creating a language server. It's as easy
as implementing the language-specific handlers.

I created this package as a library for myself, so that I can quickly implement
custom IDE features whenever I have some need. I hope you find it useful too!

The root `lsp` package provides a high-level framework for creating a language
server. It includes a `Server` that handles communication with the client and
delegates the actual language server implementation using handlers.

For users who need more flexibility, the low-level `jsonrpc` and `types`
packages provide the core types needed to fully customize the server
implementation to your heart's content.

## What's a language server?

A language server is a program that provides language-specific features like
auto-completion, go-to-definition, and hover information to an editor or IDE.
Read more about the [Language Server
Protocol](https://microsoft.github.io/language-server-protocol/).

## Feature roadmap

- [ ] Feature-complete implementation of the [Language Server Protocol Speficication - 3.17](https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/)
- [ ] A server framework that handles communication with the client and delegates the actual language server implementation using handlers.
- [ ] Test utilities for testing language servers.
- [ ] Examples included in all packages.
