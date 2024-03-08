import express from "express";
import jwt from "jsonwebtoken";
import bodyParser from "body-parser";

const port = 3000;
let solution = "";
const app = express();

app.use(bodyParser.text());
app.use(bodyParser.urlencoded({ extended: true }));

type jwtSec = { jwt_secret: string };
const response = await fetch(
  `https://hackattic.com/challenges/jotting_jwts/problem?access_token=${process.env.ACCESS_TOKEN}`
);

const response_payload = await (response.json() as Promise<jwtSec>);

console.log("Response Payload ", response_payload);
const secret = response_payload.jwt_secret;

const result = await fetch(
  `https://hackattic.com/challenges/jotting_jwts/solve?access_token=${process.env.ACCESS_TOKEN}`,
  {
    method: "POST",
    body: JSON.stringify({
      app_url: "https://jotting-jwts-ayyk.onrender.com/",
    }),
    headers: { "Content-Type": "application/json" },
  }
);

console.log("Response to jotting_jwts submission: ", await result.json());

app.get("/", (req, res) => {
  res.send("Hang on!.. I am jotting jwts!!");
});

app.post("/", (req, res) => {
  try {
    const decodedData = jwt.verify(req.body, secret) as jwt.JwtPayload;
    if (decodedData.append) {
      solution = solution + decodedData.append;
      console.log("solution: ", solution);
      res.send("Done! Send More!");
    } else {
      res.send({ solution });
    }
  } catch (error) {
    console.log("IN ERROR: Request Body ", req.body);
    console.log(error);
    res.send("Didn't found token in body");
  }
});

app.listen(port, () => {
  console.log(`App running at port - ${port}`);
});
