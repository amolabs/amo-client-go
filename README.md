# amo-client-go
Reference implementation of AMO client for golang

## Introduction
The purposes of this software are:
* To demonstrate how to interact with the AMO blockchain nodes and AMO storage
  services.
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

