import * as functions from 'firebase-functions';

// // Start writing Firebase Functions
// // https://firebase.google.com/docs/functions/typescript
//
export const generateVoucherIntake = functions.https.onRequest((request, response) => {
  console.log("Got request = ", { 
    headers: request.rawHeaders,
    body: JSON.stringify(request.body),
  });
  if (request.method !== "POST" || request.header("content-type") !== "application/json") {
    return response.sendStatus(404);
  }
  return response.status(200).send(request.body);
});
