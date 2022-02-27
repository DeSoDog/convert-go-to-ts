import axios from "axios";
import fs from "fs";
import { getLib, getRoutes } from "./routes";
const GET_STRUCT_OBJ = /\ntype[a-zA-Z\n\s0-9]+ {.*?\n}/gs;
const GET_STRUCT_INT = /\ntype[a-zA-Z\s]+?(u|\s)int[0-9]*\n/gs;
const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?string/gs;
const GET_STRUCT_FUNC = /\ntype[a-zA-Z\s]+?[^\n]func.+?\)/gs;
const GET_STRUCT_UTXO = /\ntype [a-zA-Z\s]+? UtxoKey\n/gs;
const GET_IMPORTS = /\nimport \(.+?\)\n/gs;
const GET_STRUCT_BYTE =
  /\ntype[a-zA-Z\s]+?(\[HashSizeBytes\]byte|\[[0-9a-zA-Z\.]*\]byte)/gs;

const getTypes = (base, routes: string[]): Promise<string[]>[] => {
  return routes.map((route) => {
    return axios
      .get(`https://raw.githubusercontent.com/${base}${route}`)
      .then((r) => [route, r.data as string]);
  });
};

const getStructs = (rawFile: string): string[] => {
  let file = flattenImportedObjs(rawFile);
  const structsObjects = [...file.matchAll(GET_STRUCT_OBJ)].map((x) => {
    return x[0];
  });

  const structsInt = [...file.matchAll(GET_STRUCT_INT)].map((x) => x[0]);
  const structsString = [...file.matchAll(GET_STRUCT_STRING)].map((x) => x[0]);
  const structsByte = [...file.matchAll(GET_STRUCT_BYTE)].map((x) => {
    return x[0].replace(/\[(.*?)\]/gs, "[100]");
  });

  const structsUtxo = [...file.matchAll(GET_STRUCT_UTXO)].map((x) => x[0]);
  const structsFunc = [...file.matchAll(GET_STRUCT_FUNC)].map((x) => x[0]);
  const imports = [...file.matchAll(GET_IMPORTS)].map((x) => x[0]);
  const allStructs = [
    ...structsObjects,
    ...structsInt,
    ...structsString,
    ...structsByte,
    ...structsUtxo,
    ...structsFunc,
  ];
  // has a length of zero? that means there were no structs in the file
  return allStructs.length !== 0 ? [...imports, ...allStructs] : [];
};
// Might eventually write a regex for this but the command go mod tidy handles auto importing everything
// except for imports with an alias.
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
// flatten obj in the struct so *lib.requireObject becomes requiredObject
const flattenImportedObjs = (file: string) => {
  file = file.replace(/\*lib\./gs, "");
  file = file.replace(/countries\./gs, "");
  return file;
};
const writeFile = (file: string, fileName: string, index: number) => {
  // "
  fs.writeFile(
    `${__dirname}/generated/${fileName}`,
    ["package types", file].join("\n\n"),
    (err) => {}
  );
};
const main = async (): Promise<void> => {
  const structsRoutes = await getTypes(
    "deso-protocol/backend/main/routes/",
    getRoutes()
  );
  const structsLib = await getTypes("deso-protocol/core/main/lib/", getLib());
  const countryTypes = await getTypes("deso-protocol/backend/main/countries/", [
    "iso-3166-1-alpha-3-codes.go",
  ]);

  [...new Set([...structsLib, ...structsRoutes, ...countryTypes])].map(
    async (file, index) => {
      const [route, fileContents] = await file;
      const structs = getStructs(fileContents);
      if (structs.length) {
        if (route === "mempool.go") {
          structs.push("type MempoolTxFeeMinHeap []*MempoolTx");
        }
        writeFile(structs.sort().join(" "), route, index);
      }
    }
  );
};
main();
