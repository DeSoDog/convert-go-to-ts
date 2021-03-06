"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const child_process_1 = require("child_process");
const fs_1 = __importDefault(require("fs"));
const main = () => {
    (0, child_process_1.exec)(`go run ./generated/main/go-to-ts.go`, (err, stdout, stderr) => {
        console.log({ err });
        fs_1.default.writeFile("./generated/types/types.ts", stdout, (err) => {
            console.log(err);
        });
    });
};
main();
//# sourceMappingURL=go-to-ts.js.map