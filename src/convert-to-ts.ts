import { exec } from "child_process";

import fs from "fs";
const main = () => {
  exec(`go run ..\\..\\to-ts.go`, (err, stdout, stderr) => {
    // console.log("oy", stdout);
    fs.writeFile("blob.ts", stdout, (err) => {
      console.log(err);
    });
  });
};
main();
