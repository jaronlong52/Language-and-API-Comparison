import express, { Request, Response } from "express";
import mysql2 from "mysql2";
import dotenv from "dotenv";

// Load environment variables from .env file
// used to privately store password
dotenv.config();

// Create an Express application
const app = express();

// Define a port number
const port = process.env.DB_PORT || 3000;

// Middleware to parse JSON request bodies
app.use(express.json());

// mySQL database connection
const db = mysql2.createConnection({
	host: process.env.DB_HOST,
	user: process.env.DB_USER,
	password: process.env.DB_PASSWORD,
	database: process.env.DB_NAME,
});

// Define a route
app.get("/", (req: Request, res: Response) => {
	res.send("Welcome to the TypeScript Express API!");
});

// get request for "users" endpoint that returns a list of users
app.get("/users", (req: Request, res: Response) => {
	db.query<mysql2.RowDataPacket[]>(
		"SELECT * FROM users",
		(err: mysql2.QueryError | null, data: mysql2.RowDataPacket[]) => {
			if (err) {
				console.error("Error during progress update:", err);
				return res.status(500).json({ error: "Database error" });
			}
			return res.json(data);
		}
	);
});

// post request for "addUser" endpoint that adds a new user
app.post("/addUser", (req: Request, res: Response) => {
	interface AddUserRequestBody {
		username: string;
		name: string;
	}

	const username: string = (req.body as AddUserRequestBody).username;
	const name: string = (req.body as AddUserRequestBody).name;

	db.query<mysql2.ResultSetHeader>(
		"INSERT INTO users (username, name) VALUES (?,?)",
		[username, name],
		(err, data) => {
			if (err) {
				return res.status(500).json({ error: "Database error" });
			}
			return res
				.status(201)
				.json({ message: "User added", insertId: data.insertId });
		}
	);
});

// Start the server
app.listen(port, () => {
	console.log(`Server running at http://localhost:${port}`);
});
