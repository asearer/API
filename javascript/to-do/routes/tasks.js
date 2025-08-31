/**
 * Routes for Task API
 */
const express = require("express");
const router = express.Router();
const taskController = require("../controllers/taskController");

// CRUD routes
router.post("/", taskController.createTask);   // Create
router.get("/", taskController.getTasks);      // Read all
router.get("/:id", taskController.getTask);    // Read one
router.put("/:id", taskController.updateTask); // Update
router.delete("/:id", taskController.deleteTask); // Delete

module.exports = router;
