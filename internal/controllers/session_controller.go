package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"UserService/internal/models"
	"UserService/internal/services"
)

type SessionController interface {
	CreateSession(c *gin.Context)
	DeleteSession(c *gin.Context)
}

type SessionControllerImpl struct {
	sessionService services.SessionService
}

func NewSessionControllerImpl(sessionService services.SessionService) SessionController {
	return &SessionControllerImpl{
		sessionService: sessionService,
	}
}

func (s *SessionControllerImpl) CreateSession(c *gin.Context) {
	var userToAuthenticate models.User

	if err := c.ShouldBindJSON(&userToAuthenticate); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	sessionId, err := s.sessionService.CreateSession(c, userToAuthenticate)
	if err != nil {
		c.JSON(errorToCode(err), err.Error())
		return
	}

	sessionIdStr, ok := sessionId.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, "Error creating session")
		return
	}

	c.SetCookie("session_id", sessionIdStr, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "Session created")
}

func (s *SessionControllerImpl) DeleteSession(c *gin.Context) {
	sessionId, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = s.sessionService.DeleteSession(c, sessionId)
	if err != nil {
		c.JSON(errorToCode(err), err.Error())
		return
	}
	c.SetCookie("session_id", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "Session deleted")
}
