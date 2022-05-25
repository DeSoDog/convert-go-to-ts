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
  });
  let blob = await Promise.all(blobPromise);
  const structNames = blob
    .join(" ")
    .matchAll(/(?<=type )[A-Za-z]*(?= struct)/gs);
  const goCommand = [...structNames]
    .filter(
      (s) =>
        s.toString().includes("Response") || s.toString().includes("Request")
    )
    .map((s) => {
      return `s.Add(types.${s}{})\n`;
    });
  writeFile(
    [
      "package main",
      getImports(),

      getBodyStart(),
      goCommand.join("  "),
      getBodyEnd(),
    ],
    "go-to-ts.go",
    "/../generated/main/"
  );
  writeFile(
    [
      "package types",
      getTypeImports(),
      "type MempoolTxFeeMinHeap []*MempoolTx",
      blob.join(" "),
    ],
    "types.go",
    "/../generated/types/"
  );
};

main();
const writeFile = (file: string[], name: string, path: string) => {
  fs.writeFile(`${__dirname}${path}${name}`, file.join("\n\n"), (err) => {});
};

// types "types/blob"
const getTypeImports = () => {
  return `import (
	"bytes"
	"container/list"
	"context"
	"math/big"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
	"github.com/DataDog/datadog-go/statsd"
	"github.com/btcsuite/btcd/addrmgr"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/deso-protocol/core/lib"
    "github.com/btcsuite/btcutil/hdkeychain"
	"github.com/deso-protocol/go-deadlock"
	merkletree "github.com/deso-protocol/go-merkle-tree"
	"github.com/dgraph-io/badger"
	"github.com/go-pg/pg"
	"github.com/holiman/uint256"
	"github.com/kevinburke/twilio-go"
	"golang.org/x/sync/semaphore"
	"github.com/decred/dcrd/lru"
	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
  	"github.com/cloudflare/circl/group"
    	chainlib "github.com/btcsuite/btcd/blockchain"
  	"github.com/oleiade/lane"
	"honnef.co/go/tools/config")`;
};

const getImports = (): string => {
  return `import (
	types "types/generated/types"
	"flag"
	"io"
	"log"
	"os"

	"github.com/OneOfOne/struct2ts"
)`;
};

const getBodyStart = (): string => {
  return `func main() {
	log.SetFlags(log.Lshortfile)

	var (
		f   = os.Stdout
		err error
	)

	flag.Parse()
	if err = runStruct2TS(f); err != nil {
		panic(err)
	}
}

func runStruct2TS(w io.Writer) error {
	s := struct2ts.New(&struct2ts.Options{
		Indent: "	",

		NoAssignDefaults: false,
		InterfaceOnly:    true,

		NoConstructor: false,
		NoCapitalize:  false,
		MarkOptional:  false,
		NoToObject:    false,
		NoExports:     false,
		NoHelpers:     false,
		NoDate:        false,

		ES6: false,
	})
`;
};
const getBodyEnd = () => {
  return `io.WriteString(w, "// this file was automatically generated, DO NOT EDIT")
	return s.RenderTo(w) }`;
};
