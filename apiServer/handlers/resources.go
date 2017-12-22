package handlers

import (
	"apiServer/entities"
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
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
	query := h.Env.DB.Order("created_at").Find(&resources)
	if err := handleQueryError(query, c); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, mapArray(resources))
}

func (h *Handlers) BookResource(c echo.Context) error {
	resourceID := c.Param("id")

	tx := h.Env.DB.Begin()
	resource := &entities.Resource{}
	query := tx.Set("gorm:query_option", "FOR UPDATE").First(resource, resourceID)
	if err := handleQueryError(query, c); err != nil {
		tx.Rollback()
		return err
	}

	if !resource.Availability {
		tx.Rollback()
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Resource have already been book",
		})
	}

	resource.Availability = false
	query = tx.Save(resource)
	if err := handleQueryError(query, c); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return c.JSON(http.StatusOK, mapStruct(resource))
}

func (h *Handlers) RestoreAllResources(c echo.Context) error {
	tx := h.Env.DB.Begin()
	query := tx.Model(entities.Resource{}).Update(entities.Resource{Availability: true})
	if err := handleQueryError(query, c); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return c.JSON(http.StatusOK, map[string]bool{
		"success": true,
	})
}

func handleQueryError(query *gorm.DB, c echo.Context) error {
	if query.Error != nil {
		if query.RecordNotFound() {
			return echo.NewHTTPError(http.StatusNotFound, "Resource not found")
		}
		return query.Error
	}
	return nil
}
