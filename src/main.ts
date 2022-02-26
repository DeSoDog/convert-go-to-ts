import axios from "axios";
import fs from "fs";
import { getLib, getRoutes } from "./routes";
const GET_STRUCT_OBJ = /\ntype[a-zA-Z\n\s]+ {.*?\n}/gs;
const GET_STRUCT_INT = /\ntype[a-zA-Z\s]+?int[0-9]*/gs;
const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?string/gs;
const GET_STRUCT_UTXO = /\ntype [a-zA-Z\s]+? UtxoKey\n/gs;
// const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?[a-zA-Z0-9\.]+? \n/gs;
const GET_STRUCT_BYTE =
  /\ntype[a-zA-Z\s]+?(\[HashSizeBytes\]byte|\[[0-9a-zA-Z\.]*\]byte)/gs;
const getTypes = async (base, routes: string[]) => {
  let response: any = routes.map(async (route) => {
    return (
      await axios.get(`https://raw.githubusercontent.com/${base}${route}`)
    ).data;
  });

  let files = (await Promise.all(response)).join(" ");
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
};

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
const main = async () => {
  let structsRoutes = await getTypes(
    "deso-protocol/backend/main/routes/",
    getRoutes()
  );
  let structsLib = await getTypes("deso-protocol/core/main/lib/", getLib());

  fs.writeFile(
    `${__dirname}/generated/types.go`,
    [...["package types\n"], ...[...structsRoutes, ...structsLib].sort()].join(
      "\n\n"
    ),
    (err) => {
      // console.log(err);
    }
  );
  // stucts;
};
main();
const flatten = () => {};
