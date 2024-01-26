type packed_data = { bytes: string };

const response = await fetch(
  `https://hackattic.com/challenges/help_me_unpack/problem?access_token=${process.env.ACCESS_TOKEN}`
);

const packed_data = await(response.json() as Promise<packed_data>);

console.log("Data to unpack: ", packed_data);

const buffer = Buffer.from(packed_data.bytes, "base64");

// Using node buffer api - https://nodejs.org/api/buffer.html
const unpacked_data = {
  int: buffer.readInt32LE(0),
  uint: buffer.readUint32LE(4),
  short: buffer.readInt16LE(8),
  float: buffer.readFloatLE(12),
  double: buffer.readDoubleLE(16),
  big_endian_double: buffer.readDoubleBE(24),
};

console.log("Unpacked data: ", unpacked_data);

const result = await fetch(
  `https://hackattic.com/challenges/help_me_unpack/solve?access_token=${process.env.ACCESS_TOKEN}`,
  {
    method: "POST",
    body: JSON.stringify(unpacked_data),
    headers: { "Content-Type": "application/json" },
  }
);

console.log("Response to unpack submission: ", await result.json());
