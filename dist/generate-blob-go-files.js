"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const routes_1 = require("./routes");
const fs_1 = __importDefault(require("fs"));
const utils_1 = require("./utils");
const main = () => __awaiter(void 0, void 0, void 0, function* () {
    const structsRoutes = yield (0, utils_1.getRawFiles)("deso-protocol/backend/main/routes/", (0, routes_1.getRoutes)());
    const structsLib = yield (0, utils_1.getRawFiles)("deso-protocol/core/main/lib/", (0, routes_1.getLib)());
    const countryTypes = yield (0, utils_1.getRawFiles)("deso-protocol/backend/main/countries/", ["iso-3166-1-alpha-3-codes.go"]);
    const blobPromise = [
        ...new Set([...structsLib, ...structsRoutes, ...countryTypes]),
    ].map((filePromise) => __awaiter(void 0, void 0, void 0, function* () {
        const file = yield filePromise;
        return (0, utils_1.getStructs)(file[1], false).join(" ");
    }));
    let blob = yield Promise.all(blobPromise);
    const structNames = blob
        .join(" ")
        .matchAll(/(?<=type )[A-Za-z]*(?= struct)/gs);
    const goCommand = [...structNames].map((s) => {
        return `s.Add(${s}{})\n`;
    });
    writeFile([
        "package types",
        getImports(),
        getBodyStart(),
        goCommand.join("  "),
        getBodyEnd(),
    ], "go-to-ts.go");
    writeFile([
        "package types",
        getTypeImports(),
        "type MempoolTxFeeMinHeap []*MempoolTx",
        blob.join(" "),
    ], "blobby.go");
});
main();
const writeFile = (file, name) => {
    fs_1.default.writeFile(`${__dirname}/../generated/blob/${name}`, file.join("\n\n"), (err) => { });
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
const getImports = () => {
    return `import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/OneOfOne/struct2ts"
)`;
};
const getBodyStart = () => {
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
//# sourceMappingURL=generate-blob-go-files.js.map