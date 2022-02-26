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
const GET_STRUCT_OBJ = /\ntype[a-zA-Z\n\s]+ {.*?\n}/gs;
const GET_STRUCT_INT = /\ntype[a-zA-Z\s]+?int[0-9]*/gs;
const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?string/gs;
const GET_STRUCT_UTXO = /\ntype [a-zA-Z\s]+? UtxoKey\n/gs;
// const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?[a-zA-Z0-9\.]+? \n/gs;
const GET_STRUCT_BYTE = /\ntype[a-zA-Z\s]+?(\[HashSizeBytes\]byte|\[[0-9a-zA-Z\.]*\]byte)/gs;
const getTypes = (base, routes) => __awaiter(void 0, void 0, void 0, function* () {
    let response = routes.map((route) => __awaiter(void 0, void 0, void 0, function* () {
        return (yield axios_1.default.get(`https://raw.githubusercontent.com/${base}${route}`)).data;
    }));
    let files = (yield Promise.all(response)).join(" ");
    files = files.replace(/lib\./gs, "");
    let structsObjects = [...files.matchAll(GET_STRUCT_OBJ)].map((x) => x[0]);
    let structsInt = [...files.matchAll(GET_STRUCT_INT)].map((x) => x[0]);
    let structsString = [...files.matchAll(GET_STRUCT_STRING)].map((x) => x[0]);
    let structsByte = [...files.matchAll(GET_STRUCT_BYTE)].map((x) => x[0]);
    let structsUtxo = [...files.matchAll(GET_STRUCT_UTXO)].map((x) => x[0]);
    // let structsString = [...files.matchAll(GET_STRUCT_STRING)].map((x) => x[0]);
    return [
        ...structsObjects,
        ...structsInt,
        ...structsString,
        ...structsByte,
        ...structsUtxo,
    ];
});
// const getCountries = async () => {
//   let countries: any = getCountryPackage().map(async (route) => {
//     return (await axios.get(route)).data;
//   });
//   countries = await Promise.all(countries);
//   console.log(countries);
//   countries.forEach((c, i) => {
//     fs.writeFile(`${__dirname}/generated/countries${i}.go`, c, (err) => {
//       // console.log(err);
//     });
//   });
// };
const main = () => __awaiter(void 0, void 0, void 0, function* () {
    let structsRoutes = yield getTypes("deso-protocol/backend/main/routes/", (0, routes_1.getRoutes)());
    let structsLib = yield getTypes("deso-protocol/core/main/lib/", (0, routes_1.getLib)());
    fs_1.default.writeFile(`${__dirname}/generated/types.go`, [...["package types\n"], ...[...structsRoutes, ...structsLib].sort()].join("\n\n"), (err) => {
        // console.log(err);
    });
    // stucts;
});
main();
const flatten = () => { };
//# sourceMappingURL=main.js.map