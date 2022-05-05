import { getLib, getRoutes } from "./routes";
import { getRawFiles, getStructs } from "./utils";
import fs from "fs";
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

  [...new Set([...structsLib, ...structsRoutes, ...countryTypes])].map(
    async (file, index) => {
      const [route, fileContents] = await file;
      const structs = getStructs(fileContents, true);
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
export const writeFile = (file: string, fileName: string, index: number) => {
  fs.writeFile(
    `${__dirname}/../generated/individual/${fileName}`,
    ["package types", file].join("\n\n"),
    (err) => {}
  );
};
