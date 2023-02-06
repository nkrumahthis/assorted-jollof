import express, {Request, Response} from "express"
import userRouter from "./routers/user.router"
import customerRouter from "./routers/customer.router"
import orderRouter from "./routers/order.router"
import paymentRouter from "./routers/payment.router"

const app = express();
app.use(express.json())
const port = 8000;

app.get('/', (req: Request, res:Response) => {
  res.send('Welcome to Assorted Jollof');
});

app.use("/users", userRouter)
app.use("/customers", customerRouter)
app.use("/orders", orderRouter)
app.use("/payments", paymentRouter)

app.listen(port, () => {
  console.log(`[server]: Server is running at http://localhost:${port}`);
});