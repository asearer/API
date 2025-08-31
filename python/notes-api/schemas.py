from pydantic import BaseModel

class NoteCreate(BaseModel):
    """
    Schema for creating a new note.

    Attributes:
        title (str): Title of the note.
        content (str): Content/body of the note.
    """
    title: str
    content: str

class NoteUpdate(BaseModel):
    """
    Schema for updating an existing note.

    Attributes:
        title (str, optional): New title for the note.
        content (str, optional): New content for the note.
    """
    title: str | None = None
    content: str | None = None
