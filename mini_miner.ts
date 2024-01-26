import { createHash } from "node:crypto";
type block = { nonce: null | number; data: Array<any> };
type minerData = {
  difficulty: number;
  block: block;
};

const response = await fetch(
  `https://hackattic.com/challenges/mini_miner/problem?access_token=${process.env.ACCESS_TOKEN}`
);

const minerData = await (response.json() as Promise<minerData>);
console.log("Data to unpack: ", minerData);

const isEnoughZeroes = (block: block, numofZeroes: number) => {
  console.log("Block : ", JSON.stringify(block));
  const sha = createHash("sha256").update(JSON.stringify(block)).digest("hex");
  console.log("Sha : ", sha);
  let zeroesCount = 0;
  for (let i = 0; i < sha.length; i++) {
    const char = sha.charAt(i);

    switch (char) {
      case "0":
        zeroesCount = zeroesCount + 4;
        break;
      case "1":
        zeroesCount = zeroesCount + 3;
        break;
      case "2":
      case "3":
        zeroesCount = zeroesCount + 2;
        break;
      case "4":
      case "5":
      case "6":
      case "7":
        zeroesCount = zeroesCount + 1;
        break;
      default:
        break;
    }

    if (char !== "0") {
      if (zeroesCount < numofZeroes) {
        return false;
      }
    }
  }
  return true;
};

let nonce = 0;
while (
  !isEnoughZeroes(
    { data: minerData.block.data, nonce: nonce },
    minerData.difficulty
  )
) {
  nonce++;
}

console.log("Nonce: ", nonce);

const result = await fetch(
  `https://hackattic.com/challenges/mini_miner/solve?access_token=${process.env.ACCESS_TOKEN}`,
  {
    method: "POST",
    body: JSON.stringify({ nonce }),
    headers: { "Content-Type": "application/json" },
  }
);

console.log("Response to miniminer submission: ", await result.json());
