from datetime import datetime
from pydantic import BaseModel

class Note(BaseModel):
    """
    Note model representing a single note.

    Attributes:
        id (int): Unique identifier for the note.
        title (str): Title of the note.
        content (str): Content/body of the note.
        created_at (datetime): Timestamp when the note was created.
    """
    id: int
    title: str
    content: str
    created_at: datetime

