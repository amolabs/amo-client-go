# amo-client-go
Golang을 위한 AMO 클라이언트의 참조 구현. This document is available in
[English](README.md) also.

## Introduction
AMO 클라이언트란 [AMO client RPC
specification](https://github.com/amolabs/docs/blob/master/rpc.md)과 [AMO
storage
specification](https://github.com/amolabs/docs/blob/master/storage.md)를
준수하는 모든 소프트웨어 프로그램, 혹은 하드웨어 장치를 의미한다. AMO
Labs에서는 참조 구현 형태로 CLI 프로그램을 제공하지만, 일상적인 작업에도 그대로
사용할 수 있다. 이 프로그램은 AMO 블록체인 노드와 AMO 스토리지 서버라는 두 종류의 원격 서버와 통신한다. AMO 블록체인 노드와의 통신은 필수이지만, data 거래를 직접 하는 경우가 아니라면 AMO 스토리지 서버와의 통신은 필요하지 않다.

## 설치
### 컴파일된 바이너리 설치
TBA
### 소스코드로부터 설치
`amo-client-go`를 컴파일하여 설치하기 전에 먼저 다음을 설치해야 한다.
* [git](https://git-scm.com)
* [make](https://www.gnu.org/software/make/)
  * Debian이나 Ubuntu linux에서는 `build-essential` package를 설치하면 함께
	설치된다.
  * MacOS에서는 Xcode에 포함된 make를 사용할 수도 있고,
	[Homebrew](https://brew.sh)를 통하여 GNU Make를 설치할 수도 있다.
* [golang](https://golang.org/dl/) v1.13 혹은 이후 버전 (Go Modules 이용)
  * 경우에 따라서 `GOPATH`와 `GOBIN` 환경변수를 직접 설정해야 할 수 있다. 더
	진행하기 전에 이 환경변수들을 확인해야 한다.

터미널에서 `go get github.com/amolabs/amo-client-go/cmd/amocli`라고 입력하거나,
다음과 같이 각 과정을 직접 수행할 수도 있다:
* github에서 소스코드 다운로드:
```bash
mkdir -p $GOPATH/src/github.com/amolabs
cd $GOPATH/src/github.com/amolabs
git clone https://github.com/amolabs/amo-client-go
```
* 컴파일 및 설치:
```bash
cd amo-client-go
make install
```

컴파일된 바이너리는 `$GOPATH/bin/amocli`에 설치된다. `PATH` 환경변수에
`$GOBIN`이 포함돼 있다면 터미널에 `amocli`라고만 입력하여 `amo-client-go`
프로그램을 구동할 수 있다.

### 클라이언트 라이브러리 설치
별도의 설치 과정이 필요하지는 않다.
`https://github.com/amolabs/amo-client-go/lib` 패키지를 golang library로서 바로
사용할 수 있다.

## 원격 서버
`amocli`는 클라이언트 프로그램이기 때문에 사용자 요청을 처리하기 위해서는 원격
서버의 주소가 필요하다. 다음의 두가지 원격 서버가 필요하다:
* AMO 블록체인 RPC 노드
* AMO 스토리지 API 서버

**AMO 블록체인 RPC 노드**는 AMO 블록체인 네트워크에 연결돼 있고 RPC 서비스를
제공하기만 하면 어떠한 노드여도 상관 없다. 공개돼 있는 노드에 연결할 수도 있고
자신의 클라이언트를 위한 전용 노드를 구축하여 사용할 수도 있다. RPC 노드의
주소는 IP 주소와 포트 번호로 이루어진다. 일반적으로 포트 번호는 26657이지만
노드에 따라 다른 포트 번호를 사용할 수 있다. 이 포트 번호가 자신의 네트워크의
방화벽에 의해 차단되고 있지는 않는지 확인한다.

TBA: AMO Labs에서 제공하는 공개 RPC 노드들

**AMO 스토리지 API 서버**는 자신이 선택한 스토리지 서비스의 API 종단 주소이다.
여러 개의 AMO 스토리지 서비스가 있을 수 있으며 사용자는 이 중에서 자신이 사용할
서비스를 선택한다. 가용성의 문제가 있을 수 있으므로, AMO Labs에서는 기본
스토리지 서비스를 제공하며 이를 AMO 기본 스토리지 서비스라 한다. 

TBA: 기본 스토리지 서비스의 API 단말 주소

## Keyring 보호
블록체인에 전송되는 모든 transaction은 사용자 키에 의해 서명되어야 한다.
`amocli`는 `$HOME/.amocli/keys/keys.json`에 위치한 keyring의 파일에 저장된
키들을 사용한다. 이 키들은 사용자 옵션에 따라 암호화되지 않은 형태로 저장될 수
있기 때문에 잘 보호할 필요가 있다. 이 keyring 파일은 파일 소유주만 읽고 쓸 수
있도록 할 것을 권고한다(Linux나 MacOS에서 `chmod` 명령을 사용할 경우 모드
0600으로 설정).

## 사용법

### 구조
`amocli`는 `amocli [flags] <command> [args...]`의 형태나 `amocli [flags]
<command> <subcommand> {args...}`의 형태로 사용되며 `command`에 따라 다르다.

단말에서만 동작하는 명령:
* `amocli [flags] version`: 현재 `amocli`의 버전 표시
* `amocli [flags] key <subcommand>`: keyring의 키를 관리

원격 서버와 함께 동작하는 명령:
* `amocli [flags] query <subcommand>`: 블록체인 데이터 조회
* `amocli [flags] tx <subcommand>`: 블록체인에 서명된 거래 전송
* `amocli [flags] parcel <subcommand>`: 스토리지 서비스의 데이터 parcel 관리

### 전역 플래그
원격 서버와 함께 동작하는 명령들은 `--rpc`나 `--sto` 플래그를 필요로 할 수
있다. (대부분의 경우 둘 중 하나가 필요하다.)
* `--rpc <ip>:<port>`: 원격 AMO 블록체인 RPC 노드 지정
  * `status`, `query`, `tx` 명령
* `--sto <ip>:<port>`: 원격 AMO 스토리지 서비스 지정
  * `parcel` 명령

`--json` 플래그는 명령의 결과가 사용자 친화적인 문구 대신 JSON object로
표시되도록 만든다.

사용자 키를 필요로 하는 명령들을 실행할 때는 어떤 사용자 키를 사용할 것인지
지정해야 한다. `--user` 플래그를 입력하면 `amocli`가 keyring에서 해당하는
사용자명을 검색한다. 그렇지 않은 경우 화면에 키 목록을 표시하고 사용자 이름의
입력을 요구하고 멈춘다. 사용자 키가 암호화된 채로 keyring에 저장돼 있는 경우
`amocli`는 프롬프트를 출력하고 암호를 입력할 것을 요구하고 멈춘다. `--pass`
플래그를 통해 암호를 입력한 경우는 사용자 입력을 따로 요구하지 않는다.

### Version 명령
```bash
amocli version
```
현재 `amocli`의 버전을 표시한다.

### Key 명령
```bash
amocli key <subcommand>
```
Keyring의 키를 관리한다. `amocli`는 키들을 keyring 파일에 저장된 형태로
관리한다. 각각의 키에는 *사용자명*이 할당된다. 이 *사용자명*은 키 사이의 식별을
쉽게 해서 사용성을 높이기 위한 것으로, 실제 블록체인 프로토콜 차원에서는 특별한
의미를 갖지 않는다. 이는 단지 프로그램이 아닌 사람을 위한 편의기능이다.

```bash
amocli key list
```
Keyring에 저장된 키들의 목록을 출력한다. 출력은 다음의 열들로 이루어진다:
* `#`: keyring 내의 위치
* `username`: 해당 키의 사용자명
* `enc`: 암호화 여부(`o` or `x`)
* `address`: 해당 키의 계정 주소

```bash
amocli key import <private_key> --useranme <username> [flags]
```
평문 비밀키를 외부로부터 읽어들여 keyring에 저장한다. `<private_key>`는
base64로 인코딩된 문자열이라고 가정한다. 이 명령에서 사용 가능한 플래그들은
다음과 같다:
* `--encrypt[=false]`: 키를 암호화된 형태로 저장할지 여부 (기본은 암호화함)

```bash
amocli key export <username>
```
Keyring으로부터 비밀키 하나를 외부로 출력한다.

**주의: 현재 구현상 비밀키를 평문의 형태로 터미널에 출력한다. 따라서 이 비밀키가 의도치 않게 노출되지 않도록 주의해야 한다.**

```bash
amocli key generate <username> [flags]
```
새로운 키를 생성하여 keyring에 추가한다. 이 명령에서 사용 가능한 플래그들은
다음과 같다:
* `--seed <seed_string>`: 비밀키를 생성하는 데 seed로 사용되는 임의의 문자열
* `--encrypt[=false]`: 키를 암호화된 형태로 저장할지 여부 (기본은 암호화함)

**주의: seed 문자열이 노출되면 누구나 같은 키를 생성할 수 있다. 이 경우
사용자의 자산이 탈취될 위험이 있다. 따라서 이 seed 문자열이 노출되지 않도록
주의해야 한다.**

```bash
amocli key remove <username>
```
Keyring에서 키 하나를 제거한다.

**주의: 현재 구현에는 백업이 없다. 따라서 제거된 키는 완전히 상실된다. 이 키가
디지털 자산을 보유하고 있는 계정에 연결된 키라면 해당 자산에 대한 통제권 역시
상실하게 된다.**

### Query 명령
```bash
amocli query <subcommand>
```
블록체인 상의 데이터를 조회한다. 모든 `query` 하부명령은 사용자 키를 필요로 하지
않는다.

```bash
amocli query node [flags] 
```
블록체인 노드의 상태를 표시한다.

```bash
amocli query config [flags]
```
블록체인 노드의 설정을 표시한다.

```bash
amocli query balance <address> [flags]
```
계정의 AMO 잔고를 출력한다. `<address>`는 HEX 인코딩된 바이트 열이다. 계정의
잔고는 AMO 단위와 *mote* 단위 두가지로 표시된다. 1 AMO는 1000000000000000000
mote과 같다. `--json` 플래그가 주어졌을 경우는 *mote* 단위만 문자열의 형태로
출력된다. `--udc <udc_id>` 플래그가 주어졌을 경우는 계정의 UDC 잔고를 출력한다. 

```bash
amocli query udc <udc_id> [flags]
```
블록체인 상에서 발행된 UDC의 전반적인 정보를 표시한다. `<udc_id>`는 십진법
숫자이다.

```bash
amocli query lock <udc_id> <address> [flags]
```
계정의 잠긴 UDC 잔고를 출력한다. `<address>`는 HEX 인코딩된 바이트 열이다.
잠긴 잔고는 AMO 단위와 *mote* 단위 두가지로 표시된다. 1 AMO는
1000000000000000000 mote과 같다. `--json` 플래그가 주어졌을 경우는 *mote*
단위만 문자열의 형태로 출력된다.

```bash
amocli query stake <address> [flags]
```
계정의 stake 정보를 출력한다. `<address>`는 HEX 인코딩된 바이트 열이다. 이
명령은 stake와 연결된 validator 공개키도 함께 표시한다. `--json` 플래그가
주어졌을 경우는 출력은 다음과 같은 형태가 된다:
```json
{"amount":"100000000000000000000","validator":[2,159,24,22,130,8,178,58,184,144,63,228,30,59,242,78,67,4,214,169,251,33,154,132,147,202,252,180,160,43,19,241]}
```
Validator 공개키는 바이트열로 표시된다.

```bash
amocli query delegate <address> [flags]
```
계정의 위임된 stake 정보를 출력한다. `<address>`는 HEX 인코딩된 바이트 열이다.

```bash
amocli query incentive <block_height | address> [flags]
```
`<block_height>` 혹은 `<address>`의 incentive 정보를 출력한다. 두 가지 모두
동시에 주어질 수 있다. `<block_height>`는 십진법 숫자이다. `<address>`는 HEX
인코딩된 바이트 열이다.

```bash
amocli query draft <draft_id> [flags]
```
제안된 draft 상태를 출력한다. `<draft_id>`는 십진법 숫자이다.

```bash
amocli query vote <draft_id> <address> [flags]
```
주어진 투표자 주소의 투표 현황을 출력한다. `<draft_id>`는 십진법 숫자이다.
`<address>`는 HEX 인코딩된 바이트 열이다.

```bash
amocli query parcel <parcelID> [flags]
```
등록된 parcel의 정보를 출력한다. `<parcelID>`는 HEX 인코딩된 바이트 열이다.

```bash
amocli query request <buyer_address> <parcel_id> [flags]
```
Parcel에 대한 구매요청 정보를 출력한다. `<buyer_address>`와 `<parcel_id>`는 HEX
인코딩된 바이트 열이다.

```bash
amocli query usage <buyer_address> <parcel_id> [flags]
```
Parcel에 대한 사용허가 정보를 출력한다. `<buyer_address>`와 `<parcel_id>`는 HEX
인코딩된 바이트 열이다.

### Tx 명령
```bash
amocli tx <subcommand>
```
서명된 거래를 블록체인에 전송한다. 모든 `tx` 하부명령은 사용자 키를 필요로
한다. `--user`와 `--pass` 플래그를 사용하여 선택한 키를 사용하게 지정한다. 이
사용자 계정이 전송되는 거래의 송신자로 간주된다.

```bash
amocli tx transfer <address> <amount> [flags]
```
송신자는 `<amount>` 만큼의 AMO coin(mote 단위)을 `<address>`로 식별되는 계정에
송금한다.

```bash
amocli tx stake <validator_pubkey> <amount> [flags]
```
송신자는 `<validator_pubkey>`에 연결된 stake를 새로 생성하거나 기존의 stake를
`<amount>` 만큼 증액한다. `<validator_pubkey>`는 HEX 인코딩된 바이트 열이다.
사용자는 서로 다른 validator 공개키를 갖는 복수의 stake를 보유할 수 없다.

AMO coin을 stake로 지정한다는 것은 사용자가 블록 생성 과정에 validator로서
참여하겠다는 의미이다. 모든 validator 노드(AMO 블록체인 네트워크 문서 참조)는
validator 키쌍을 가지고 있어야 한다. 어떤 validator key와 연결해서 coin을
stake로 지정하려면 해당 validator key가 설치된 validator 노드가 존재해야 한다.
비록 제대로 된 validator 공개키가 아닌 무작위의 바이트 열이나 아직 실행되지
않는 validator 노드의 validator 공개키를 지정하여 stake를 생성할 수 있지만,
실제 동작하는 validator 노드가 없이는 블록 생성 과정으로부터 보상을 받지 못하게
된다. 보다 안전한 방법은 validator 노드를 먼저 실행한 후 해당 validator에
연결된 stake를 생성하는 것이다.

```bash
amocli tx withdraw <amount> [flags]
```
송신자는 자신의 계정에 stake된 AMO coin을 `<amount>`만큼 인출한다. 계정에
stake된 AMO coin의 양이 0이 되면 해당 stake가 완전히 제거된다.

```bash
amocli tx delegate <address> <amount> [flags]
```
송신자는 자신의 AMO coin을 다른 계정에 위임한다. 이 계정은 stake를 보유한
계정이어야 한다. 사용자는 복수의 계정에게 coin을 위임할 수 없다.

```bash
amocli tx retract <amount> [flags]
```
송신자는 `<amount>` 만큼의 AMO coin 위임을 철회한다.

```bash
amocli tx register <parcel_id> <key_custody> [flags]
```
송신자는 `<parcel_id>`를 갖는 데이터 parcel을 소유주의 키 보관값
`<key_custody>`와 함께 블록체인에 등록한다. `<parcel_id>`와 `<key_custody>`는
HEX 인코딩된 바이트 열이다. `<parcel_id>`는 구매자들이 저장 위치를 식별할 수
있도록 AMO 스토리지 서비스로부터 획득한 것이어야 한다. 기술적으로는
`<key_custody>`는 특별한 의미가 없는 임의의 값을 사용해도 되지만, 이를 저장하는
목적은 소유주의 암호화키를 안전하게 보관하려는 것이다. 따라서 이 값은 데이터
암호화 키(DEK: data encryption key)를 소유주의 공개키로 암호화한 것이어야 한다.
이 DEK는 소유주가 AMO 스토리지 서비스에 데이터를 업로드할 때 데이터 parcel의
본체를 암호화할 때 쓰인 키이다.

```bash
amocli tx request <parcel_id> <amount> [flags]
```
송신자는 소유주에게 데이터 parcel `<parcel_id>`의 사용권을 요청하며
`<amount>`의 AMO coin을 댓가로 지불하겠다고 약속한다. `<amount>`의 AMO coin이
송신자의 계정에서 차감되고 블록체인 내에 데이터 parcel 요청과 함께 잠긴다.

```bash
amocli tx grant <parcel_id> <address> <key_custody> [flags]
```
송신자는 데이터 parcel `<parcel_id>`의 사용권을 계정 `<address>`에게 허가하며
구매자의 키 보관값 `<key_custody>`를 전달한다. 구매자는 나중에 블록체인을
조회하여 자신에게 전달된 암호화키를 획득할 수 있다. 기술적으로는
`<key_custody>`는 특별한 의미가 없는 임의의 값을 사용해도 되지만, 이를 저장하는
목적은 구매자의 암호화키를 안전하게 보관하려는 것이다. 따라서 이 값은 데이터
암호화키(DEK: data encryption key)를 구매자의 공개키로 암호화한 것이어야 한다.

```bash
amocli tx discard <parcel_id> [flags]
```
송신자는 블록체인으로부터 데이터 parcel `<parcel_id>`를 폐기한다. 이 이후에는
어떤 구매자도 AMO 스토리지 서비스로부터 해당 데이터 parcel을 다운로드하지
못한다.

```bash
amocli tx cancel <parcel_id> [flags]
```
송신자는 자신이 이전에 전송했던 데이터 parcel 요청을 취소하고 블록체인으로부터
삭제한다. 지불하기로 약속되어 잠겼던 금액은 송신자에게 환불된다.

```bash
amocli tx revoke <parcel_id> <address> [flags]
```
송신자는 자신이 이전에 계정 `<address>`에게 허가했던 데이터 parcel
`<parcel_id>`에 대한 사용 허가를 파기한다. 이 이후에 계정 `<address>`는 AMO
스토리지 서비스로부터 해당 데이터 parcel을 다운로드하지 못한다.

### Parcel 명령
```bash
amocli parcel <subcommand>
```
스토리지 서비스의 데이터 parcel들을 관리한다. `upload`, `download`와 `remove`
하부명령은 사용자 키를 필요로 한다. `--user`와 `--pass` 플래그를 사용하여
선택한 키를 사용하게 지정한다. 이 사용자 계정이 전송되는 AMO 스토리지 API
요청의 송신자로 간주된다.

```bash
amocli parcel upload {<hex> | --file <filename>} [flags]
```
송신자는 AMO 스토리지 서비스에 새로운 데이터 parcel을 업로드한다. 송신자는
업로드된 데이터 parcel의 소유주로 지정된다. AMO 스토리지 서비스는 업로드된
데이터로부터 생성된 parcel의 ID를 응답한다. 업로드할 데이터를 지정하는 데에는
두 가지 방식이 있다: HEX 인코딩된 바이트 열을 직접 입력하거나, 읽어들일 파일
이름을 지정한다. 이는 클라이언트 쪽의 기능이고 서버는 언제나 HEX 인코딩된
바이트 열을 수신하게 된다.

```bash
amocli parcel download <parcelID> [flags]
```
송신자는 AMO 스토리지 서비스로부터 데이터 parcel `<parcelID>`를 다운로드한다.
데이터 parcel의 소유주가 해당 데이터 parcel에 대한 사용권한을 송신자에게 허가해
놓은 상태라면 서버는 데이터 parcel의 본체와 metadata를 응답한다. 그렇지 않으면
에러를 응답한다. 데이터 parcel의 소유주는 `amocli tx grant` 명령으로 특정
데이터 parcel에 대한 사용권한을 허가할 수 있다.

```bash
amocli parcel inspect <parcelID> [flags]
```
송신자는 AMO 스토리지 서비스로부터 데이터 parcel `<parcelID>`의 metadata를
다운로드한다. 이 명령을 소유주로부터의 허가를 필요로 하지 않는다. 따라서 데이터
parcel이 스토리지 서버에 존재하기만 하면 언제나 성공한다.

```bash
amocli parcel remove <parcelID> [flags]
```
송신자는 AMO 스토리지 서비스로부터 데이터 parcel `<parcelID>`를 삭제한다.
송신자는 해당 데이터 parcel의 소유주여야 한다.
