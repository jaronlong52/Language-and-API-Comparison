import express, { Request, Response } from "express";

// Create an Express application
const app = express();

// Define a port number
const port = 3000;

// Middleware to parse JSON request bodies
app.use(express.json());

// Define a route
app.get("/", (req: Request, res: Response) => {
	res.send("Welcome to the TypeScript Express API!");
});

// Create a simple "users" endpoint that returns a list of users
app.get("/users", (req: Request, res: Response) => {
	const users = [
		{ id: 1, name: "John Doe" },
		{ id: 2, name: "Jane Doe" },
	];
	res.json(users);
});

// Start the server
app.listen(port, () => {
	console.log(`Server running at http://localhost:${port}`);
});
