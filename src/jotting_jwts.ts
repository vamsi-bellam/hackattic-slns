type jwt = { jwt_secret: string };

const response = await fetch(
  `https://hackattic.com/challenges/jotting_jwts/problem?access_token=${process.env.ACCESS_TOKEN}`
);

const jwt = await(response.json() as Promise<jwt>);

console.log("Data to unpack: ", jwt);
