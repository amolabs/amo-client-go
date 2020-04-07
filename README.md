# amo-client-go [![GoDoc](https://godoc.org/github.com/amolabs/amo-client-go/lib?status.svg)](https://godoc.org/github.com/amolabs/amo-client-go/lib)
Reference implementation of AMO client for golang. This document is available
in [Korean](README.ko.md) also.

## Introduction
An AMO client is any software or hardware application conforming to
[AMO client RPC specification](https://github.com/amolabs/docs/blob/master/rpc.md),
and optionally,
[AMO storage specification](https://github.com/amolabs/docs/blob/master/storage.md).
AMO Labs provides a CLI software program as a reference implementation, but it
can be usable for day-to-day workflow. This program interacts with two types of
remote servers: AMO blockchain nodes and AMO storage servers. It is essential
to interact with AMO blockchain nodes, but it is unnecessary to interact with
AMO storage servers unless you intend to trade parcels on your own.

The purposes of this software are:
* To provide a reference implementation for demonstrating how to interact with
  AMO blockchain nodes, and optionally, AMO storage services.
* To provide a ready-to-use library for those who don't have much resources to
  write their own client implementation.
* To provide a out-of-the-stock CLI program to quickly interact with AMO
  infrastructure services.

with the following conditions:
* As little dependency to the [AMO ABCI](https://github.com/amolabs/amoabci)
  codes as possible (hopefully, no dependency)
* No direct dependency to the
  [Tendermint](https://github.com/tendermint/tendermint) codes
* No direct dependency to the [Ceph](https://github.com/ceph) or librgw codes

## Installation
### Install pre-compiled binary
TBA
### Install from source code
Before compile and install `amo-client-go`, you need to install the followings:
* [git](https://git-scm.com)
* [make](https://www.gnu.org/software/make/)
  * For Debian or Ubuntu linux, you can install `build-essential` package.
  * For MacOS, you can use `make` from Xcode, or install GNU Make via
	[Homebrew](https://brew.sh).
* [golang](https://golang.org/dl/) v1.13 or later (for Go Modules)
  * In some cases, you need to set `GOPATH` and `GOBIN` environment variable
	manually. Check these variables are set before you proceed.

You can just type `go get github.com/amolabs/amo-client-go/cmd/amocli` in the
terminal, or do every step on your own as the following:
* download source code from github:
```bash
mkdir -p $GOPATH/src/github.com/amolabs
cd $GOPATH/src/github.com/amolabs
git clone https://github.com/amolabs/amo-client-go
```
* complie and install:
```bash
cd amo-client-go
make install
```

Compiled bianry will be installed in `$GOPATH/bin/amocli`. If your `PATH`
environment variable contains `$GOBIN`, then you can just type `amocli` in a
termianl to invoke a `amo-client-go` program.

### Install client library
No prior installation is needed. You may use
`https://github.com/amolabs/amo-client-go/lib` package as a golang library for
your client program.

## Remote servers
Since `amocli` is a client program, it needs remote server addresses to perform
user requests. There are two types of remote servers:
* AMO blockchain RPC node
* AMO storage API server

**AMO blockchain RPC node** is any AMO blockchain node which is connected to
AMO blockchain network and provides an RPC service. You can connect to any
publicly available nodes or run your own dedicated node for your clients. An
RPC node address is composed of an IP address and a port number. The default
port number is 26657. Make sure the firewall of your network does not block
this port number.

TBA: Public RPC node addresses provided by AMO Labs

**AMO storage API server** is an API endpoint of a storage service of your
choice. There can be numerous AMO storage services and you can select your
storage from them. Since there might be an availability issue, AMO Labs
provides a default storage service, called AMO default storage service.

TBA: Address of default storage serivce API endpoint

## Keyring protection
Any transaction sent to the blockchain should be signed with a user key.
`amocli` utilizes keys saved in a keyring file located at
`$HOME/.amocli/keys/keys.json`. There may be an unencrypted key in a keyring if
appropriate option is given. So you need to protect this keyring. It is
recommended to set a `read` and `write` permissions of this keyring file to
owner-only (use mode `0600` on Linux and MacOS when using `chmod` command).

## Usage

### Structure
`amocli` can be used as `amocli [flags] <command> [args...]` or `amocli
[flags] <command> <subcommand> {args...}` depending on `command`.

Commands working locally:
* `amocli [flags] version`: print current `amocli` version
* `amocli [flags] key <subcommand>`: manage keys in the local keyring

Commands working with a remote server:
* `amocli [flags] query <subcommand>`: query blockchain data
* `amocli [flags] tx <subcommand>`: send a signed transaction to the blockchain
* `amocli [flags] parcel <subcommand>`: manage data parcels in the storage service

### Global flags
Commands working with a remote server may need either `--rpc` or `--sto` flag.
(In most cases, they need one of them.)
* `--rpc <ip>:<port>`: used to specify remote AMO blockchain RPC node
  * `status`, `query`, `tx` commands
* `--sto <ip>:<port>`: used to specify remote AMO storage service endpoint
  * `parcel` command

`--json` flag is used to indicate that command result should be displayed as a
JSON object rather than a human-friendly format.

When you execute commands requiring a user key, you need to specify which user
key should be used. If you supply a `--user` flag, `amocli` will search in the
keyring for the specified user key. Otherwise, `amocli` will display the list
of stored keys and pause for username. In case the user key is encrypted in the
keyring, `amocli` will display a prompt message and pause for a passphrase
input. This action can be suppressed if you supply a `--pass` flag to the
command.

### Version command
```bash
amocli version
```
Print current `amocli` version.

### Key command
```bash
amocli key <subcommand>
```
Manage keys in a local keyring. `amocli` manages keys saved in a local keyring
file. Each key is associated with a *username*. This *username* is used to
differentiate between keys and increase usability, but does not have special
meaning in the AMO blockchain protocol. It is just a convenience feature for a
*human* user.

```bash
amocli key list
```
Print a list of stored keys in the local keyring. Output comprises of the
following columns:
* `#`: index in the keyring
* `username`: username for this key
* `enc`: encrypted(`o`) or not(`x`)
* `address`: account address for this key

```bash
amocli key import <private_key> --useranme <username> [flags]
```
Import a private key into the local keyring. `<private_key>` is assumed to be a
base64-encoded string. Available flags are:
* `--encrypt[=false]`: store a key in an encrypted form or not (default true)

```bash
amocli key export <username>
```
Export a private key from the local keyring.

**CAUTION: current implementation just prints a private key in a plain-text
form to the terminal. So, be careful not to expose this plain-text key.**

```bash
amocli key generate <username> [flags]
```
Generate a new key and add to the local keyring. Available flags are:
* `--seed <seed_string>`: arbitrary string used as a seed when generating a
  private key
* `--encrypt[=false]`: store a key in an encrypted form or not (default true)

**CAUTION: if the seed string is exposed, anyone can generate the same key as
yours. Your assets are in danger in that case. So, be careful not to expose
this seed string.**

```bash
amocli key remove <username>
```
Remove a key from the local keyring.

**CAUTION: there is no backup in current implementation. So the key is lost
permanently. If this is the key for the address holding some digital assets,
you will lose control over these assets.**

### Query command
```bash
amocli query <subcommand>
```
Query blockchain data. All of `query` subcommands do not require user key.

```bash
amocli query node [flags]
```
Print blockchain node status.

```bash
amocli query config [flags]
```
Print blockchain node config.

```bash
amocli query balance <address> [flags]
```
Print AMO balance of an account. `<address>` is a HEX-encoded byte array.
Account balance is displayed in two ways: in AMO unit and in *mote* unit. 1 AMO
is equivalent to 1000000000000000000 mote. When supplied with `--json` flag,
only *mote* unit is displayed as a string form.

```bash
amocli query udc <udc_id> [flags]
```
Print general information of UDC issued on blockchain. `udc_id` is a decimal
number.

```bash
amocli query stake <address> [flags]
```
Print stake info of an account. `<address>` is a HEX-encoded byte array. This
subcommand displays a validator public key associated with a stake. When
supplied with `--json` flag, output has the following form:
```json
{"amount":"100000000000000000000","validator":[2,159,24,22,130,8,178,58,184,144,63,228,30,59,242,78,67,4,214,169,251,33,154,132,147,202,252,180,160,43,19,241]}
```
Validator public key is displayed as a byte array.

```bash
amocli query delegate <address> [flags]
```
Print delegated stake info of an account. `<address>` is a HEX-encoded byte
array.

```bash
amocli query parcel <parcelID> [flags]
```
Print registered parcel info. `<parcelID>` is a HEX-encoded byte array.

```bash
amocli query request <buyer_address> <parcel_id> [flags]
```
Print buyer request info on a parcel. `<buyer_address>` and `<parcel_id>` are
HEX-encoded byte arrays.

```bash
amocli query usage <buyer_address> <parcel_id> [flags]
```
Print granted usage info on a parcel. `<buyer_address>` and `<parcel_id>` are
HEX-encoded byte arrays.

### Tx command
```bash
amocli tx <subcommand>
```
Send a signed transaction to the blockchain. All of `tx` subcommands require a
user key, so you may supply `--user` and `--pass` flags to indicate which key
to use when signing. And this user account is assumed to be the sender of the
transaction.

```bash
amocli tx transfer <address> <amount> [flags]
```
The sender transfers `<amount>` of AMO coin (in mote unit) to the account
associated with `<address>`.

```bash
amocli tx stake <validator_pubkey> <amount> [flags]
```
The sender creates a new stake or increase an existing stake associated with
the `<validator_pubkey>` by `<amount>` of AMO coin. `<validator_pubkey>` is a
HEX-encoded byte array. A user cannot have multiple stakes with different
validator public keys.

Staking coins means the user wants to participate in a block production process
as a validator. Every validator node(see AMO blockchain network document) has a
validator key pair. Staking coins associated with a validator key requires
there exists a validator node loaded with the associated validator public key.
Although it is possible to stake coins with a random byte array rather than a
valid validator public key or a public key of a validator node not yet running,
without actual running validator node the staker will not get any reward from
the block production process. It is safer to run a validator node first and
stake coins for the validator.

```bash
amocli tx withdraw <amount> [flags]
```
The sender withdraws `<amount>` of AMO coin staked for the account. If the
staked AMO coin reaches zero, then the stake is removed completely.

```bash
amocli tx delegate <address> <amount> [flags]
```
The sender delegates his/her AMO coins to other account. This account must have
staked coins. A user cannot delegate coins to multiple accounts.

```bash
amocli tx retract <amount> [flags]
```
The sender retract coin delegation by `<amount>` of AMO coins.

```bash
amocli tx register <parcel_id> <key_custody> [flags]
```
The sender registers a data parcel having `<parcel_id>` to the blockchain along
with owner's key custody `<key_custody>`. `<parcel_id>` and `<key_custody>` are
HEX-encoded byte arrays. `<parcel_id>` must be a value acquired from AMO
storage service so that buyers can locate it later. Technically,
`<key_custody>` can be an arbitrary value, but its purpose is to store owner's
encryption key in a secure manner. So, it should be a data decryption key(DEK)
encrypted by the owner's public key. The DEK is the key used to encrypt a data
parcel body when uploading to one of AMO storage service.

```bash
amocli tx request <parcel_id> <amount> [flags]
```
The sender requests a data parcel `<parcel_id>` to the owner pledging he/she
would pay `<amount>` of AMO coins as a data price. `<amount>` of AMO coins will
be subtracted from the sender's account and it will be locked with the data
parcel request in the blockchain.

```bash
amocli tx grant <parcel_id> <address> <key_custody> [flags]
```
The sender grant a usage on a data parcel `<parcel_id>` to an account
`<address>` along with buyer's key custody `<key_custody>`. The buyer can query
blockchain to acquire his/her encryption key later. Technically,
`<key_custody>` can be an arbitrary value, but its purpose is to store buyer's
encryption key in a secure manner. So, it should be a data encryption key(DEK)
encrypted by the buyer's public key.

```bash
amocli tx discard <parcel_id> [flags]
```
The sender discard a data parcel `<parcel_id>` from the blockchain. After this,
any buyer cannot download the data parcel from an AMO storage service.

```bash
amocli tx cancel <parcel_id> [flags]
```
The sender cancel a data parcel request sent from the sender and remove it from
the blockchain. The pledged AMO coins will be refunded to the sender.

```bash
amocli tx revoke <parcel_id> <address> [flags]
```
The sender revoke a granted usage on a data parcel `<parcel_id>` to an account
`<address>`. After this, the account `<address>` cannot download the data
parcel from an AMO storage service.

### Parcel command
```bash
amocli parcel <subcommand>
```
Manage data parcels in the storage service. Subcommands `upload`, `download`
and `remove` require a user key, so you may supply `--user` and `--pass` flags
to indicate which key to use when signing. And this user account is assumed to
be the sender of the AMO storage API request.

```bash
amocli parcel upload {<hex> | --file <filename>} [flags]
```
The sender uploads a new data parcel to an AMO storage service. The sender
will be designated as the owner of the uploaded data parcel. An AMO storage
service will respond with a parcel id newly generated by the uploaded data.
There are two ways to specify data to be uploaded: specify HEX-encoded byte
array directly, or specify filename to be read. This is a client-side feature,
and a server always receives HEX-encoded byte array as the data body.

```bash
amocli parcel download <parcelID> [flags]
```
The sender downloads a data parcel `<parcelID>` form an AMO storage service. If
the owner of the data parcel have granted a usage on the data parcel to the
sender, the server will respond with a data parcel body and metadata.
Otherwise, the server will respond with an error. An owner of a data parcel can
grant a data parcel usage using `amocli tx grant` command.

```bash
amocli parcel inspect <parcelID> [flags]
```
The sender downlaods a metadata of a data parcel `<parcelID>` from an AMO
storage service. This command does not require any permission from the owner of
the data parcel. It will always succeed if the parcel exists on the storage
server.

```bash
amocli parcel remove <parcelID> [flags]
```
The sender removes a data parcel `<parcelID>` from an AMO storage service. The
sender must be the owner the data parcel.

## Lib

### AMO blockchain RPC

https://godoc.org/github.com/amolabs/amo-client-go/lib/rpc

AMO blockchain is based on [Tendermint](https://github.com/tendermint/tendermint).
Client supports basic Tendermint [RPCs](https://docs.tendermint.com/master/rpc/).

```text
rpc.NodeStatus()
rpc.BroadcastTx(tx)
rpc.ABCIQuery(path, queryData)
```

Client also provides wrapper of above functions for AMO blockchain specific
[RPCs](https://github.com/amolabs/docs/blob/master/rpc.md)
```text
// Key is needed for signing transaction
// Transaction definition 
// https://github.com/amolabs/docs/blob/master/protocol.md

// Bank
rpc.Transfer(...)
// Stake
rpc.Stake(...)
// Governance
rpc.Propose(...)
// Parcel
rpc.Register(...)
// UDC
rpc.Issue(...)

// Return types of ABCI Query are defined in AMO blockchain RPC document
// https://github.com/amolabs/docs/blob/master/rpc.md#abci-query
// ABCI Query
rpc.QueryBalance(0, address)   // Query AMO of address
rpc.QueryBalance(udc, address) // Query UDC of address
```

### AMO Storage

- TODO
