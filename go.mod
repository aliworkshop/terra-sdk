go 1.16

module github.com/aliworkshop/terra-sdk

require (
	github.com/99designs/keyring v1.1.6
	github.com/armon/go-metrics v0.3.10
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/cockroachdb/apd/v2 v2.0.2
	github.com/coinbase/rosetta-sdk-go v0.7.2
	github.com/confio/ics23/go v0.6.6
	github.com/cosmos/btcutil v1.0.4
	github.com/cosmos/cosmos-proto v1.0.0-alpha6
	github.com/cosmos/cosmos-sdk/db v1.0.0-beta.1
	github.com/cosmos/cosmos-sdk/errors v1.0.0-beta.2
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/iavl v0.17.3
	github.com/cosmos/ledger-cosmos-go v0.11.1
	github.com/gogo/gateway v1.1.0
	github.com/gogo/protobuf v1.3.3
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/go-getter v1.5.11
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/hdevalence/ed25519consensus v0.0.0-20210204194344-59a8610d2b87
	github.com/improbable-eng/grpc-web v0.15.0
	github.com/jhump/protoreflect v1.10.1
	github.com/lazyledger/smt v0.2.1-0.20210709230900-03ea40719554
	github.com/magiconair/properties v1.8.5
	github.com/mattn/go-isatty v0.0.14
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.1
	github.com/prometheus/common v0.32.1
	github.com/rakyll/statik v0.1.7
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/rs/zerolog v1.26.1
	github.com/spf13/cast v1.4.1
	github.com/spf13/cobra v1.4.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.10.1
	github.com/stretchr/testify v1.7.1
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.35.3
	github.com/tendermint/tm-db v0.6.6
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
	pgregory.net/rapid v0.4.7
	sigs.k8s.io/yaml v1.3.0
)

// latest grpc doesn't work with with our modified proto compiler, so we need to enforce
// the following version across all dependencies.
replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.1.7-0.20210622111912-ef00f8ac3d76
	github.com/cosmos/cosmos-sdk/db => ./db
	github.com/gin-gonic/gin => github.com/gin-gonic/gin v1.7.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
