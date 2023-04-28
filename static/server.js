const express = require("express");
const app = express();

app.use(express.static("./"));
app.listen(9988, () => {
  console.log("Listen on the port 9988...");
});
