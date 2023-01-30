import express, {Request, Response} from "express"
import userRouter from "./user.router"
import orderRouter from "./order.router"

const app = express();
app.use(express.json())
const port = 8000;

app.get('/', (req: Request, res:Response) => {
  res.send('Welcome to Assorted Jollof');
});

app.use("/users", userRouter)
app.use("/orders", orderRouter)

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});