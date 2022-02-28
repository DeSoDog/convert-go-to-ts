import { exec } from "child_process";
const main = () => {
  exec(
    `go run github.com/OneOfOne/struct2ts/cmd/struct2ts types/types.AdminGetNFTDropRequest -i`,
    (err, stdout, stderr) => {
      console.log("adf");
      console.log(stdout);
      console.log(err);
      console.log(stderr);
    }
  );
};
main();
