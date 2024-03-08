import express from "express";
import jwt from "jsonwebtoken";
import bodyParser from "body-parser";

const port = 3000;
let solution = "";
const app = express();

app.use(bodyParser.text());
app.use(bodyParser.urlencoded({ extended: true }));

let secret = "";

app.get("/", (req, res) => {
  res.send("Hang on!.. I am jotting jwts!!");
});

app.post("/setjwt", (req, res) => {
  try {
    secret = req.body;
    res.send("JWT SECRET set successfully!");
  } catch (error) {
    console.log(error);
  }
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
