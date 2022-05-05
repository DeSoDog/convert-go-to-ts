import fs from "fs";
import axios from "axios";

const GET_STRUCT_OBJ = /\ntype[a-zA-Z\n\s0-9]+ {.*?\n}/gs;
const GET_STRUCT_INT = /\ntype[a-zA-Z\s]+?(u|\s)int[0-9]*\n/gs;
const GET_STRUCT_STRING = /\ntype[a-zA-Z\s]+?string/gs;
const GET_STRUCT_FUNC = /\ntype[a-zA-Z\s]+?[^\n]func.+?\)/gs;
const GET_STRUCT_UTXO = /\ntype [a-zA-Z\s]+? UtxoKey\n/gs;
const GET_IMPORTS = /\nimport \(.+?\)\n/gs;
const GET_STRUCT_BYTE =
  /\ntype[a-zA-Z\s]+?(\[HashSizeBytes\]byte|\[[0-9a-zA-Z\.]*\]byte)/gs;
export const getStructs = (
  rawFile: string,
  includeImports = true
): string[] => {
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
  const imports = includeImports
    ? [...file.matchAll(GET_IMPORTS)]
        .map((x) => x[0])
        .map((x) => x.replace(/\t"/gs, '_"'))
    : [];
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

export const flattenImportedObjs = (file: string) => {
  file = file.replace(/\*lib\./gs, "");
  file = file.replace(/countries\./gs, "");
  return file;
};

export const getRawFiles = (base, routes: string[]): Promise<string[]>[] => {
  return routes.map((route) => {
    return axios
      .get(`https://raw.githubusercontent.com/${base}${route}`)
      .then((r) => [route, r.data as string]);
  });
};
