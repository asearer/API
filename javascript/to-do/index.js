/**
 * To-Do API using Express
 * Provides CRUD operations for tasks.
 */
const express = require("express");
const taskRoutes = require("./routes/tasks");

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware to parse JSON requests
app.use(express.json());

// Task routes
app.use("/tasks", taskRoutes);

// Root route
app.get("/", (req, res) => {
    res.send("Welcome to the To-Do API 🚀");
});

// Start server
app.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}`);
});

