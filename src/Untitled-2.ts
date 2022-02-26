import axios from "axios";
import fs from "fs";
const GET_STRUCT_REGEX = /(type)(.+?)(struct {)(.+?)(})/gs;
const getStructures = async (): Promise<string> => {
  const response = (
    await axios.get(
      "https://raw.githubusercontent.com/deso-protocol/backend/main/routes/transaction.go"
    )
  ).data;
  return response;
};

const main = async () => {
  const response = await getStructures();

  let oy = [...response.matchAll(GET_STRUCT_REGEX)][0].join(" ");

  oy;

  fs.writeFile(`${__dirname}/types.go`, oy, (err) => {
    console.log(err);
  });
};
main();
