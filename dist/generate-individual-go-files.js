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
Object.defineProperty(exports, "__esModule", { value: true });
exports.writeFile = void 0;
const routes_1 = require("./routes");
const utils_1 = require("./utils");
const main = () => __awaiter(void 0, void 0, void 0, function* () {
    const structsRoutes = yield (0, utils_1.getRawFiles)("deso-protocol/backend/main/routes/", (0, routes_1.getRoutes)());
    const structsLib = yield (0, utils_1.getRawFiles)("deso-protocol/core/main/lib/", (0, routes_1.getLib)());
    const countryTypes = yield (0, utils_1.getRawFiles)("deso-protocol/backend/main/countries/", ["iso-3166-1-alpha-3-codes.go"]);
    [...new Set([...structsLib, ...structsRoutes, ...countryTypes])].map((file, index) => __awaiter(void 0, void 0, void 0, function* () {
        const [route, fileContents] = yield file;
        const structs = (0, utils_1.getStructs)(fileContents);
        if (structs.length) {
            if (route === "mempool.go") {
                structs.push("type MempoolTxFeeMinHeap []*MempoolTx");
            }
            (0, exports.writeFile)(structs.sort().join(" "), route, index);
        }
    }));
});
main();
const writeFile = (file, fileName, index) => {
    // "
    // fs.writeFile(
    //   `${__dirname}/generated/${fileName}`,
    //   ["package types", file].join("\n\n"),
    //   (err) => {}
    // );
};
exports.writeFile = writeFile;
//# sourceMappingURL=generate-individual-go-files.js.map