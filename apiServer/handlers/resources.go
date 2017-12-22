package handlers

import (
	"apiServer/entities"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

type TransportResource struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Available bool      `json:"available"`
}

func mapStruct(r *entities.Resource) *TransportResource {
	return &TransportResource{
		Id:        r.ID,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
		Available: r.Availability,
	}
}

func mapArray(resources []*entities.Resource) []*TransportResource {
	t := []*TransportResource{}
	for _, r := range resources {
		t = append(t, mapStruct(r))
	}
	return t
}

func (h *Handlers) GetResources(c echo.Context) error {
	resources := []*entities.Resource{}
	if err := h.Env.DB.Find(&resources).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mapArray(resources))
}

func (h *Handlers) BookResource(c echo.Context) error {
	resourceId := c.Param("id")

	tx := h.Env.DB.Begin()
	resource := &entities.Resource{}
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(resource, resourceId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if !resource.Availability {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Resource have already been book",
		})
	}

	resource.Availability = false
	if err := tx.Save(resource).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return c.JSON(http.StatusOK, mapStruct(resource))
}

func (h *Handlers) RestoreAllResources(c echo.Context) error {
	tx := h.Env.DB.Begin()
	if err := tx.Model(entities.Resource{}).Update(entities.Resource{Availability: true}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}
