/**
 * Task Model - In-memory representation of a Task.
 * In a real-world app, this would connect to a database.
 */
class Task {
    constructor(id, title, description, completed = false) {
        this.id = id;               // Unique identifier
        this.title = title;         // Short task title
        this.description = description; // Longer description
        this.completed = completed; // Completion status (true/false)
    }
}

module.exports = Task;
