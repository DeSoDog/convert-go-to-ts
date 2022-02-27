import axios from "axios";
import fs from "fs";
import { getCountryPackage, getLib, getRoutes } from "./routes";
const GET_STRUCT_OBJ = /\ntype[a-zA-Z\n\s0-9]+ {.*?\n}/gs;
const GET_STRUCT_INT = /\ntype[a-zA-Z\s]+?(u|\s)int[0-9]*\n/gs;

const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?string/gs;
const GET_STRUCT_FUNC = /\ntype[a-zA-Z\s]+?[^\n]func.+?\)/gs;
const GET_STRUCT_UTXO = /\ntype [a-zA-Z\s]+? UtxoKey\n/gs;
// const GET_STRUCT_MEMPOOL_TX = /\ntype Mem.+?\*MempoolTx/gs;
const GET_STRUCT_BYTE =
  /\ntype[a-zA-Z\s]+?(\[HashSizeBytes\]byte|\[[0-9a-zA-Z\.]*\]byte)/gs;
const getTypes = async (base, routes: string[]) => {
  let response: any = routes.map(async (route) => {
    return (
      await axios.get(`https://raw.githubusercontent.com/${base}${route}`)
    ).data;
  });

  let files = (await Promise.all(response)).join(" ");
  files = files.replace(/\*lib\./gs, "");
  files = files.replace(/countries\./gs, "");
  let structsObjects = [...files.matchAll(GET_STRUCT_OBJ)].map((x) => {
    return x[0];
  });
  let structsInt = [...files.matchAll(GET_STRUCT_INT)].map((x) => x[0]);
  let structsString = [...files.matchAll(GET_STRUCT_STRING)].map((x) => x[0]);
  let structsByte = [...files.matchAll(GET_STRUCT_BYTE)].map((x) => {
    return x[0].replace(/\[(.*?)\]/gs, "[100]");
  });

  let structsUtxo = [...files.matchAll(GET_STRUCT_UTXO)].map((x) => x[0]);
  // let structsMem = [...files.matchAll(GET_STRUCT_MEMPOOL_TX)].map((x) => x[0]);
  let structsFunc = [...files.matchAll(GET_STRUCT_FUNC)].map((x, i) => {
    return x[0];
  });
  return [
    ...structsObjects,
    ...structsInt,
    ...structsString,
    ...structsByte,
    ...structsUtxo,
    ...structsFunc,
    // ...structsMem,
  ];
};
const main = async () => {
  let structsRoutes = await getTypes(
    "deso-protocol/backend/main/routes/",
    getRoutes()
  );
  let structsLib = await getTypes("deso-protocol/core/main/lib/", getLib());
  let countryTypes = await getTypes("deso-protocol/backend/main/countries/", [
    "iso-3166-1-alpha-3-codes.go",
  ]);
  fs.writeFile(
    `${__dirname}/generated/types.go`,
    [
      ...["package types\n"],
      ...[
        "type MempoolTxFeeMinHeap []*MempoolTx",
        ...new Set([...structsLib, ...structsRoutes, ...countryTypes]),
      ].sort(),
    ].join("\n\n"),
    (err) => {
      // console.log(err);
    }
  );
  // stucts;
};
main();
const flatten = () => {};

// muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
// chainlib "github.com/btcsuite/btcd/blockchain"
