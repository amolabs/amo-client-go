// Package rpc is a Tendermint RPC wrapper package to broadcast transaction
// and query ABCI query.
//
// Simply create, sign and broadcast transaction via functions in tx.go
//
//    // Get KeyEntry
//    key := keys.ImportKey(...)
//
//    rpc.Transfer(0, "<TO ADDRESS>", "<AMOUNT>", key, fee, lastHeight) // Send AMO Coin
//    rpc.Transfer(144, ...) // Send UDC 144
//
//    rpc.Stake("<VALIDATOR ADDRESS>", "<AMOUNT>", key, fee, lastHeight)
//    rpc.Propose("<DRAFT ID>", "<CONFIG>", "<DESCRIPTION>")
//
// Query AMO Blockchain ABCI query via functions in query.go
// Return types of ABCI Query are defined in https://github.com/amolabs/docs/blob/master/rpc.md
//
//    rpc.Balance(0, "<ADDRESS>") // Query AMO coin balance of address
//    rpc.Balance(144, "<ADDRESS>") // Query UDC 144 balance of address
//    rpc.Parcel("<PARCEL ID>") // Query parcel meta data
package rpc
