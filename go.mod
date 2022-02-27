module types

go 1.16

replace github.com/deso-protocol/core => ../core/

require (
	cloud.google.com/go/storage v1.15.0
	github.com/DataDog/datadog-go v4.8.3+incompatible
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/bxcodec/faker v2.0.1+incompatible
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/lru v1.0.0
	github.com/deso-protocol/backend v1.2.9
	github.com/deso-protocol/core v0.0.0-00010101000000-000000000000
	github.com/deso-protocol/go-deadlock v1.0.0
	github.com/deso-protocol/go-merkle-tree v1.0.0
	github.com/dgraph-io/badger/v3 v3.2103.0
	github.com/ethereum/go-ethereum v1.9.25
	github.com/fatih/structs v1.1.0
	github.com/gernest/mention v2.0.0+incompatible
	github.com/go-pg/pg/v10 v10.10.0
	github.com/golang/glog v1.0.0
	github.com/gorilla/mux v1.8.0
	github.com/h2non/bimg v1.1.7
	github.com/holiman/uint256 v1.1.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/montanaflynn/stats v0.6.6
	github.com/nyaruka/phonenumbers v1.0.74
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/sendgrid/sendgrid-go v3.11.0+incompatible
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0
	github.com/stretchr/testify v1.7.0
	github.com/tyler-smith/go-bip39 v1.1.0
	github.com/unrolled/secure v1.0.8
	google.golang.org/api v0.46.0
)
