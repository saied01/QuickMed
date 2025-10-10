package reservation

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReservationHandler struct {
	service *ReservationService
}

func NewReservationHandler(s *ReservationService) *ReservationHandler {
	return &ReservationHandler{service: s}
}

func (h *ReservationHandler) CreateReservation(c *gin.Context) {
	userID := uint(1) // TODO: replace with actual logged-in user from session/auth middleware

	resourceIDStr := c.PostForm("resource_id")
	startStr := c.PostForm("start")
	endStr := c.PostForm("end")

	resourceID64, err := strconv.ParseUint(resourceIDStr, 10, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid resource ID")
		return
	}

	layout := "2006-01-02T15:04"
	start, err := time.Parse(layout, startStr)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid start time")
		return
	}
	end, err := time.Parse(layout, endStr)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid end time")
		return
	}

	res, err := h.service.Create(userID, uint(resourceID64), start, end)
	if err != nil {
		c.String(http.StatusBadRequest, "error creating reservation")
		return
	}

	c.HTML(http.StatusOK, "reservation_success.html", gin.H{
		"Reservation": res,
	})

}
