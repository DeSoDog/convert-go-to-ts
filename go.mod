module types

go 1.16

replace github.com/deso-protocol/core => ../core/

require (
	cloud.google.com/go/storage v1.15.0
	github.com/DataDog/datadog-go v4.8.3+incompatible
	github.com/Microsoft/go-winio v0.5.2 // indirect
	github.com/OneOfOne/struct2ts v1.0.6
	github.com/btcsuite/btcd v0.22.1
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1 // indirect
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cloudflare/circl v1.1.0
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/lru v1.1.1
	github.com/deso-protocol/backend v1.2.9
	github.com/deso-protocol/core v0.0.0-00010101000000-000000000000 // indirect
	// github.com/deso-protocol/core v1.2.9
	// github.com/deso-protocol/core v0.0.0-00010101000000-000000000000
	github.com/deso-protocol/go-deadlock v1.0.0
	github.com/deso-protocol/go-merkle-tree v1.0.0
	github.com/dgraph-io/badger v1.6.2
	github.com/dgraph-io/badger/v3 v3.2103.2
	github.com/ethereum/go-ethereum v1.9.25
	github.com/fatih/structs v1.1.0
	github.com/gcash/bchutil v0.0.0-20210113190856-6ea28dff4000
	github.com/gernest/mention v2.0.0+incompatible
	github.com/go-pg/pg v8.0.7+incompatible
	github.com/go-pg/pg/v10 v10.10.0
	github.com/golang/glog v1.0.0
	github.com/gorilla/mux v1.8.0
	github.com/h2non/bimg v1.1.9
	github.com/hashicorp/golang-lru v0.5.4
	github.com/holiman/uint256 v1.1.1
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac // indirect
	github.com/kevinburke/go-types v0.0.0-20210723172823-2deba1f80ba7 // indirect
	github.com/kevinburke/rest v0.0.0-20210506044642-5611499aa33c // indirect
	github.com/kevinburke/twilio-go v0.0.0-20210327194925-1623146bcf73
	github.com/mitchellh/mapstructure v1.4.2
	github.com/montanaflynn/stats v0.6.6
	github.com/nyaruka/phonenumbers v1.0.75
	github.com/oleiade/lane v1.0.1
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pkg/errors v0.9.1
	github.com/sendgrid/sendgrid-go v3.11.1+incompatible
	github.com/shibukawa/configdir v0.0.0-20170330084843-e180dbdc8da0
	github.com/stretchr/testify v1.7.1
	github.com/ttacon/builder v0.0.0-20170518171403-c099f663e1c2 // indirect
	github.com/ttacon/libphonenumber v1.2.1 // indirect
	github.com/tyler-smith/go-bip39 v1.1.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	google.golang.org/api v0.46.0
	gopkg.in/DataDog/dd-trace-go.v1 v1.38.1
	honnef.co/go/tools v0.2.2
)
