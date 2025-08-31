/**
 * Task Controller - contains all business logic for handling tasks.
 */
const Task = require("../models/task");

let tasks = [];   // In-memory store
let nextId = 1;   // Auto-increment task ID

// Create a new task
exports.createTask = (req, res) => {
    const { title, description } = req.body;
    const task = new Task(nextId++, title, description);
    tasks.push(task);
    res.status(201).json(task);
};

// Get all tasks
exports.getTasks = (req, res) => {
    res.json(tasks);
};

// Get a task by ID
exports.getTask = (req, res) => {
    const task = tasks.find(t => t.id === parseInt(req.params.id));
    if (!task) {
        return res.status(404).json({ error: "Task not found" });
    }
    res.json(task);
};

// Update a task
exports.updateTask = (req, res) => {
    const task = tasks.find(t => t.id === parseInt(req.params.id));
    if (!task) {
        return res.status(404).json({ error: "Task not found" });
    }

    const { title, description, completed } = req.body;
    if (title !== undefined) task.title = title;
    if (description !== undefined) task.description = description;
    if (completed !== undefined) task.completed = completed;

    res.json(task);
};

// Delete a task
exports.deleteTask = (req, res) => {
    const id = parseInt(req.params.id);
    tasks = tasks.filter(t => t.id !== id);
    res.json({ message: "Task deleted" });
};
