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
const axios_1 = __importDefault(require("axios"));
const fs_1 = __importDefault(require("fs"));
const routes_1 = require("./routes");
const GET_STRUCT_OBJ = /\ntype[a-zA-Z\n\s0-9]+ {.*?\n}/gs;
const GET_STRUCT_INT = /\ntype[a-zA-Z\s]+?(u|\s)int[0-9]*\n/gs;
const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?string/gs;
const GET_STRUCT_FUNC = /\ntype[a-zA-Z\s]+?[^\n]func.+?\)/gs;
const GET_STRUCT_UTXO = /\ntype [a-zA-Z\s]+? UtxoKey\n/gs;
const GET_IMPORTS = /\nimport \(.+?\)\n/gs;
const GET_STRUCT_BYTE = /\ntype[a-zA-Z\s]+?(\[HashSizeBytes\]byte|\[[0-9a-zA-Z\.]*\]byte)/gs;
const getTypes = (base, routes) => {
    return routes.map((route) => {
        return axios_1.default
            .get(`https://raw.githubusercontent.com/${base}${route}`)
            .then((r) => [route, r.data]);
    });
};
const getStructs = (rawFile) => {
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
const getImports = () => {
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
const flattenImportedObjs = (file) => {
    file = file.replace(/\*lib\./gs, "");
    file = file.replace(/countries\./gs, "");
    return file;
};
const writeFile = (file, fileName, index) => {
    // "
    fs_1.default.writeFile(`${__dirname}/generated/${fileName}`, ["package types", file].join("\n\n"), (err) => { });
};
const main = () => __awaiter(void 0, void 0, void 0, function* () {
    const structsRoutes = yield getTypes("deso-protocol/backend/main/routes/", (0, routes_1.getRoutes)());
    const structsLib = yield getTypes("deso-protocol/core/main/lib/", (0, routes_1.getLib)());
    const countryTypes = yield getTypes("deso-protocol/backend/main/countries/", [
        "iso-3166-1-alpha-3-codes.go",
    ]);
    [...new Set([...structsLib, ...structsRoutes, ...countryTypes])].map((file, index) => __awaiter(void 0, void 0, void 0, function* () {
        const [route, fileContents] = yield file;
        const structs = getStructs(fileContents);
        if (structs.length) {
            // "type MempoolTxFeeMinHeap []*MempoolTx
            writeFile(structs.sort().join(" "), route, index);
        }
    }));
});
main();
//# sourceMappingURL=main.js.map