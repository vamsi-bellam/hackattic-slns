const APP_URL = "https://jotting-jwts-ayyk.onrender.com";
type jwtSec = { jwt_secret: string };
const response = await fetch(
  `https://hackattic.com/challenges/jotting_jwts/problem?access_token=${process.env.ACCESS_TOKEN}`
);

const responsePayload = await(response.json() as Promise<jwtSec>);

console.log("Response Payload ", responsePayload);

const setJwtResponse = await fetch(`${APP_URL}/setjwt`, {
  method: "POST",
  body: responsePayload.jwt_secret,
  headers: { "Content-Type": "application/json" },
});

console.log("Set JWT Response: ", setJwtResponse);

const result = await fetch(
  `https://hackattic.com/challenges/jotting_jwts/solve?access_token=${process.env.ACCESS_TOKEN}`,
  {
    method: "POST",
    body: JSON.stringify({
      app_url: APP_URL,
    }),
    headers: { "Content-Type": "application/json" },
  }
);

console.log("Response to jotting_jwts submission: ", await result.json());
