import express from "express"

const app = express();
const port = 8000;

app.get('/', (req, res) => {
  res.send('Welcome to Assorted Jollof');
});

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});