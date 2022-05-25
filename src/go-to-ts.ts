import { exec } from "child_process";

import fs from "fs";
const main = () => {
  exec(`go run ./generated/main/go-to-ts.go`, (err, stdout, stderr) => {
    console.log({ err });
    fs.writeFile("./generated/types/types.ts", stdout, (err) => {
      console.log(err);
    });
  });
};
main();
