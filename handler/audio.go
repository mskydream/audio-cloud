package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	storage "github.com/mskydream/audio-cloud/storage"

	"strconv"
)

const (
	MaxUploadSize = 10 << 20
)

func (h *Handler) getAllAudio(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var input storage.AudioListParam
	if err := c.BindQuery(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid query")
		return
	}

	result, err := h.services.GetAudioList(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) uploadAudio(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxUploadSize)

	file, _, err := c.Request.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	defer file.Close()

	fileId := uuid.New()

	err = h.services.StoreFile(fileId, file)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	audioId, err := h.services.UploadFile(userId, fileId.String())
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, idResponse{
		ID: audioId,
	})
}

func (h *Handler) addDescription(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	audioId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid audio id param")
		return
	}

	var input storage.UpdateAudio
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.AddDescription(userId, audioId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) downloadAudio(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	audioId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid audio id param")
		return
	}

	audio, err := h.services.DownloadFile(userId, audioId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	fileId, err := uuid.Parse(audio.FilePath)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, fileSize, err := h.services.GetFile(fileId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer file.Close()

	c.DataFromReader(http.StatusOK, fileSize, "application/octet-stream", file, map[string]string{"Content-Disposition": audio.Title + storage.FileExt})
}
