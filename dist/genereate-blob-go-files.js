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
exports.writeFile = void 0;
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
        // return getStructs(await file);
    }));
    const blob = yield Promise.all(blobPromise);
    // writeFile(blob.sort().join(" "));
    const structNames = blob
        .join(" ")
        .matchAll(/(?<=type )[A-Za-z]*(?= struct)/gs);
    const goCommand = [...structNames].map((s) => {
        return `s.Add(types.${s}{})\n`;
    });
    (0, exports.writeFile)(goCommand.join(" "), "commands.go");
});
main();
const writeFile = (file, name) => {
    fs_1.default.writeFile(`${__dirname}/generated/blob/${name}`, ["package types", getImports(), file].join("\n\n"), (err) => { });
};
exports.writeFile = writeFile;
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
//# sourceMappingURL=genereate-blob-go-files.js.map