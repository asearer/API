from fastapi import FastAPI, HTTPException
from datetime import datetime
from typing import List

from models import Note
from schemas import NoteCreate, NoteUpdate

# Initialize FastAPI app
app = FastAPI(title="Notes API", description="A simple API to manage notes", version="1.0.0")

# In-memory storage for notes (list of Note objects)
notes: List[Note] = []
next_id = 1  # Auto-incrementing ID for each new note

@app.post("/notes/", response_model=Note)
def create_note(note: NoteCreate):
    """
    Create a new note.

    Args:
        note (NoteCreate): The note data from the request body.

    Returns:
        Note: The newly created note with ID and timestamp.
    """
    global next_id
    new_note = Note(
        id=next_id,
        title=note.title,
        content=note.content,
        created_at=datetime.utcnow()
    )
    notes.append(new_note)
    next_id += 1
    return new_note

@app.get("/notes/", response_model=List[Note])
def get_notes():
    """
    Retrieve all notes.

    Returns:
        List[Note]: A list of all notes.
    """
    return notes

@app.get("/notes/{note_id}", response_model=Note)
def get_note(note_id: int):
    """
    Retrieve a single note by its ID.

    Args:
        note_id (int): The ID of the note to retrieve.

    Raises:
        HTTPException: 404 if note not found.

    Returns:
        Note: The note with the specified ID.
    """
    for note in notes:
        if note.id == note_id:
            return note
    raise HTTPException(status_code=404, detail="Note not found")

@app.put("/notes/{note_id}", response_model=Note)
def update_note(note_id: int, note_update: NoteUpdate):
    """
    Update an existing note.

    Args:
        note_id (int): ID of the note to update.
        note_update (NoteUpdate): Data to update (title/content).

    Raises:
        HTTPException: 404 if note not found.

    Returns:
        Note: The updated note.
    """
    for note in notes:
        if note.id == note_id:
            if note_update.title is not None:
                note.title = note_update.title
            if note_update.content is not None:
                note.content = note_update.content
            return note
    raise HTTPException(status_code=404, detail="Note not found")

@app.delete("/notes/{note_id}", response_model=dict)
def delete_note(note_id: int):
    """
    Delete a note by its ID.

    Args:
        note_id (int): ID of the note to delete.

    Returns:
        dict: Confirmation message.

    Raises:
        HTTPException: 404 if note not found.
    """
    global notes
    for note in notes:
        if note.id == note_id:
            notes = [n for n in notes if n.id != note_id]
            return {"message": "Note deleted"}
    raise HTTPException(status_code=404, detail="Note not found")
