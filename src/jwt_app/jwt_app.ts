import express from "express";
import jwt from "jsonwebtoken";

const port = 3000;
let solution = "";
const app = express();

//app.use(express.text());

app.use(function (req, res, next) {
  req.rawBody = "";
  req.setEncoding("utf8");

  req.on("data", function (chunk) {
    req.rawBody += chunk;
  });

  req.on("end", function () {
    next();
  });
});

let secret = "";

app.get("/", (req, res) => {
  res.send("Hang on!.. I am jotting jwts!!");
});

app.post("/setjwt", (req, res) => {
  try {
    console.log(req.rawBody);
    secret = req.rawBody;
    res.send("JWT SECRET set successfully!");
  } catch (error) {
    console.log(error);
  }
});

app.post("/", (req, res) => {
  console.log(req);
  try {
    const decodedData = jwt.verify(req.rawBody, secret) as jwt.JwtPayload;
    if (decodedData?.append) {
      solution = solution + decodedData.append;
      console.log("solution: ", solution);
      res.json({});
    } else {
      res.json({ solution });
    }
  } catch (error) {
    console.log("IN ERROR: Request Body ", req.rawBody);
    console.log(error);
    res.send("Didn't found token in body");
  }
});

app.listen(port, () => {
  console.log(`App running at port - ${port}`);
});
