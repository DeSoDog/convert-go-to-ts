"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const child_process_1 = require("child_process");
const main = () => {
    (0, child_process_1.exec)(`go run github.com/OneOfOne/struct2ts/cmd/struct2ts types/types.AdminGetNFTDropRequest -i`, (err, stdout, stderr) => {
        console.log("adf");
        console.log(stdout);
        console.log(err);
        console.log(stderr);
    });
};
main();
//# sourceMappingURL=convert-to-ts.js.map