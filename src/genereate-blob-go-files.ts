import { getLib, getRoutes } from "./routes";
import fs from "fs";
import { getRawFiles, getStructs } from "./utils";
const main = async (): Promise<void> => {
  const structsRoutes = await getRawFiles(
    "deso-protocol/backend/main/routes/",
    getRoutes()
  );
  const structsLib = await getRawFiles(
    "deso-protocol/core/main/lib/",
    getLib()
  );
  const countryTypes = await getRawFiles(
    "deso-protocol/backend/main/countries/",
    ["iso-3166-1-alpha-3-codes.go"]
  );

  const blobPromise = [
    ...new Set([...structsLib, ...structsRoutes, ...countryTypes]),
  ].map(async (filePromise) => {
    const file = await filePromise;
    return getStructs(file[1], false).join(" ");
    // return getStructs(await file);
  });
  const blob = await Promise.all(blobPromise);

  writeFile(blob.sort().join(" "));
};
main();
export const writeFile = (file: string) => {
  fs.writeFile(
    `${__dirname}/generated/blob/blob.go`,
    ["package types", getImports(), file].join("\n\n"),
    (err) => {}
  );
};
const getImports = (): string => {
  return `import (
	"container/list"
	"math/big"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
	chainlib "github.com/btcsuite/btcd/blockchain"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/btcsuite/btcd/addrmgr"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/deso-protocol/core/lib"
	"github.com/deso-protocol/go-deadlock"
	merkletree "github.com/deso-protocol/go-merkle-tree"
	"github.com/dgraph-io/badger"
	"github.com/go-pg/pg"
	lru "github.com/hashicorp/golang-lru"
	"github.com/holiman/uint256"
	"github.com/kevinburke/twilio-go"
	"honnef.co/go/tools/config"
)`;
};
