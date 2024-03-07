import express from "express";
import jwt from "jsonwebtoken";
import bodyParser from "body-parser";
import multer from "multer";

const secret = "1234";
let solution = "";
const app = express();

app.use(bodyParser.text());
app.use(bodyParser.urlencoded({ extended: true }));

const port = 3000;

app.get("/", (req, res) => {
  res.send("Hang on!.. I am jotting jwts!!");
});

app.post("/", (req, res) => {
  try {
    const decoded_data = jwt.verify(req.body, secret) as jwt.JwtPayload;
    if (decoded_data.append) {
      solution = solution + decoded_data.append;
      res.send("Done! Send More!");
    } else {
      res.send({ solution });
    }
  } catch (error) {
    console.log("request body ", req.body);
    console.log(error);
    res.send("Didn't found token in body");
  }
});

app.listen(port, () => {
  console.log(`App running at port - ${port}`);
});
