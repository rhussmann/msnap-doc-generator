import * as functions from 'firebase-functions';
import { exec } from "child_process";
import * as path from "path";
import * as fs from 'fs';

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
    response.sendStatus(404);
    return;
  }

  // Return mocked docx
  fs.readFile(path.join(__dirname, "..", "etc", "mock-intake.docx"), (err, data) => {
    if (err) {
      return response.status(500).send(err);
    } 
    return response
      .status(200)
      .contentType("application/vnd.openxmlformats-officedocument.wordprocessingml.document")
      .send(data);
  });
  return;
});


function executeBinary(arch: string = "linux.amd64"): Promise<String> {
  const goBin = path.join(__dirname, "..", "bin", `hello.${arch}`);
  return new Promise<String>((resolve, reject) => {
    exec(goBin, (error, stdout) => error ? reject(error) : resolve(stdout));
  });
}
