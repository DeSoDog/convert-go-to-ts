import { exec } from "child_process";

import fs from "fs";
const main = () => {
  exec(
    `go run ./generated/blob/blob-output/go-to-ts.go`,
    (err, stdout, stderr) => {
      fs.writeFile("./generated/blob/blob-output/blob.ts", stdout, (err) => {
        console.log(err);
      });
    }
  );
};
main();
