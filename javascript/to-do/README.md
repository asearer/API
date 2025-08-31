#  To-Do API  

A simple **To-Do API** built with **Node.js** and **Express**, demonstrating CRUD operations with in-memory storage. This project is part of the **Awesome APIs** collection, showcasing APIs across **Python, JavaScript, and Go**.  


##  Features
- Create, read, update, and delete tasks  
- Each task has:
  - `id` → unique identifier  
  - `title` → short description  
  - `description` → detailed text  
  - `completed` → boolean (default: `false`)  
- In-memory storage (no database required)  
- Clean project structure (controllers, routes, models)  


##  Installation & Setup

1. Clone the repo (or navigate to this folder if part of a larger repo):  
 
   git clone https://github.com/your-username/awesome-apis.git
   cd awesome-apis/javascript/todo-api


2. Install dependencies:

   ```bash
   npm install
   ```

3. Start the server in development mode (with auto-reload via nodemon):

   ```bash
   npm run dev
   ```

4. The API will be running at:

   ```
   http://localhost:3000
   ```

---

## 📡 API Endpoints

### Base URL

```
http://localhost:3000
```

### Endpoints

| Method | Endpoint     | Description           |
| ------ | ------------ | --------------------- |
| POST   | `/tasks`     | Create a new task     |
| GET    | `/tasks`     | Retrieve all tasks    |
| GET    | `/tasks/:id` | Retrieve a task by ID |
| PUT    | `/tasks/:id` | Update a task         |
| DELETE | `/tasks/:id` | Delete a task         |


## 🛠️ Example Usage

### Create a Task

```http
POST /tasks
Content-Type: application/json

{
  "title": "Buy groceries",
  "description": "Milk, eggs, bread"
}

```

### Get All Tasks

```http
GET /tasks
```

### Update a Task

```http
PUT /tasks/1
Content-Type: application/json

{
  "completed": true
}
```

### Delete a Task

```http
DELETE /tasks/1
```

---

## 🔮 Next Steps

* Add persistence with MongoDB or PostgreSQL
* Implement user authentication (JWT)
* Add filtering & search (e.g., completed vs incomplete tasks)
* Deploy on Render, Railway, or Heroku for demo access

