// package lsp simplifies the process of creating a language server. It
// provides a high-level framework for creating a language server. The included
// [Server] handles communication with the client and delegates the actual
// language server implementation to the [Handler].
//
// For users who need more flexibility, see the low-level [jsonrpc] and [types]
// packages. These provide the core types needed to fully customize the server
// implementation to your heart's content.
package lsp
