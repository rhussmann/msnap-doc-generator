import * as functions from 'firebase-functions';
import { exec } from "child_process";
import * as path from "path";

// // Start writing Firebase Functions
// // https://firebase.google.com/docs/functions/typescript
//
export const generateVoucherIntake = functions.https.onRequest(async (request, response) => {
  console.log("Got request = ", { 
    headers: request.rawHeaders,
    body: JSON.stringify(request.body),
  });
  const result = await executeBinary();
  console.log(`Exec result: ${result}`);

  if (request.method !== "POST" || request.header("content-type") !== "application/json") {
    return response.sendStatus(404);
  }
  return response.status(200).send(request.body);
});

function executeBinary(arch: string = "linux.amd64"): Promise<String> {
  const goBin = path.join(__dirname, "..", "bin", `hello.${arch}`);
  return new Promise<String>((resolve, reject) => {
    exec(goBin, (error, stdout) => error ? reject(error) : resolve(stdout));
  });
}
