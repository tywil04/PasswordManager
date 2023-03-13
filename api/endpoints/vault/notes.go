package vaults

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"PasswordManager/api/lib/db"
	"PasswordManager/api/lib/exceptions"
	"PasswordManager/ent"
)

const (
	NoteDescription string = ""
)

type GetNoteInput struct{}

type PostNoteInput struct {
	VaultId   string `form:"vaultId" json:"vaultId" xml:"vaultId" pmParseType:"uuid"`
	Name      string `form:"name" json:"name" xml:"name" pmParseType:"base64"`
	NameIv    string `form:"nameIv" json:"nameIv" xml:"nameIv" pmParseType:"base64"`
	Title     string `form:"title" json:"title" xml:"title" pmParseType:"base64"`
	TitleIv   string `form:"titleIv" json:"titleIv" xml:"titleIv" pmParseType:"base64"`
	Content   string `form:"content" json:"content" xml:"content" pmParseType:"base64"`
	ContentIv string `form:"contentIv" json:"contentIv" xml:"contentIv" pmParseType:"base64"`
	Colour    string `form:"colour" json:"colour" xml:"colour" pmParseType:"base64"`
	ColourIv  string `form:"colourIv" json:"colourIv" xml:"colourIv" pmParseType:"base64"`
}

type DeleteNoteInput struct {
	VaultId string `form:"vaultId" json:"vaultId" xml:"vaultId" pmParseType:"uuid"`
	NoteId  string `form:"noteId" json:"noteId" xml:"noteId" pmParseType:"uuid"`
}

func GetNote(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)

	notes, notesErr := db.GetUserNotes(authedUser)
	if notesErr != nil {
		c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
		return
	}

	jsonNotes := make([]gin.H, len(notes))
	for index, note := range notes {
		vault, vaultErr := db.GetNoteVault(note)
		if vaultErr != nil {
			c.JSON(500, exceptions.Builder("", exceptions.Unknown, exceptions.TryAgain))
			return
		}

		jsonNotes[index] = gin.H{
			"id":        note.ID.String(),
			"vaultId":   vault.ID.String(),
			"name":      note.Name,
			"nameIv":    note.NameIv,
			"title":     note.Title,
			"titleIv":   note.TitleIv,
			"content":   note.Content,
			"contentIv": note.ContentIv,
			"colour":    note.Colour,
			"colourIv":  note.ColourIv,
		}
	}

	c.JSON(200, gin.H{"notes": jsonNotes})
}

func PostNote(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	vault, vaultErr := db.GetUserVault(authedUser, params["vaultId"].(uuid.UUID))
	if vaultErr != nil {
		c.JSON(400, exceptions.Builder("vaultId", exceptions.InvalidParam, exceptions.Uuid, exceptions.Owns))
		return
	}

	newNote, newNoteErr := db.Client.Note.Create().
		SetVault(vault).
		SetName(params["name"].([]byte)).
		SetNameIv(params["nameIv"].([]byte)).
		SetTitle(params["title"].([]byte)).
		SetTitleIv(params["titleIv"].([]byte)).
		SetContent(params["content"].([]byte)).
		SetContentIv(params["contentIv"].([]byte)).
		SetColour(params["colour"].([]byte)).
		SetColourIv(params["colourIv"].([]byte)).
		Save(db.Context)

	if newNoteErr != nil {
		c.JSON(500, exceptions.Builder("note", exceptions.Creating, exceptions.TryAgain))
		return
	}

	c.JSON(200, gin.H{"noteId": newNote.ID.String()})
}

func DeleteNote(c *gin.Context) {
	authedUser := c.MustGet("authedUser").(*ent.User)
	params := c.GetStringMap("params")

	vault, vaultErr := db.GetUserVault(authedUser, params["vaultId"].(uuid.UUID))
	if vaultErr != nil {
		c.JSON(400, exceptions.Builder("vaultId", exceptions.InvalidParam, exceptions.Uuid, exceptions.Owns))
		return
	}

	dpErr := db.DeleteVaultNoteViaId(vault, params["passwordId"].(uuid.UUID))
	if dpErr != nil {
		c.JSON(400, exceptions.Builder("note", exceptions.Deleting))
		return
	}

	c.JSON(200, gin.H{})
}
