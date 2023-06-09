package equip

import (
	"jd_workout_golang/app/middleware"
	repo "jd_workout_golang/app/repositories/equip"
	"strconv"
	"github.com/gin-gonic/gin"
)

type updateFrom struct {
	Name string `json:"name" form:"name" binding:"required"`
	Note string `json:"note" form:"note"`
}

// update personal equip
// @Summary update equip
// @Description update equip for personal user
// @Tags Equip
// @Accept json
// @Produce json
// @Param id path integer true "equip id"
// @Param weights body updateFrom false "note for equip"
// @Success 200 {string} string "{'message': 'create success'}"
// @Failure 422 {string} string "{'message': '缺少必要欄位', 'error': 'error message'}"
// @Failure 403 {string} string "{'message': 'jwt token error', 'error': 'error message'}"
// @Router /equip/{id} [patch]
// @Security Bearer
func UpdateEquip(c *gin.Context) {

	id := c.Param("id")
	
	weightId , err:= strconv.ParseUint(id, 10, 32)
	
	if err != nil {
		c.JSON(422, gin.H{
			"message": "uri id error",
		})

		c.Abort()

		return
	}

	updateFrom := updateFrom{}
	if err := c.ShouldBind(&updateFrom); err != nil {
		c.JSON(422, gin.H{
			"message": "缺少必要欄位",
			"error":   err.Error(),
		})

		c.Abort()

		return
	}
	
	equip, err := repo.GetEquip(weightId, middleware.Uid)

	if err != nil {
		c.JSON(422, gin.H{
			"message": "equip not found",
			"error":   err.Error(),
		})

		c.Abort()

		return
	}

	equip.Name = updateFrom.Name
	equip.Note = updateFrom.Note

	repo.Update(equip)

	c.JSON(200, gin.H{
		"message": "equip updated",
	})
}