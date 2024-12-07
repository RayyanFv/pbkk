// controllers/reportController.go
package controllers

import (
	"net/http"
	"pos-system/services"

	"github.com/gin-gonic/gin"
)

// GetSalesReport untuk mengambil laporan penjualan berdasarkan tanggal
func GetSalesReport(c *gin.Context) {
	// Parsing query parameters
	startDate := c.DefaultQuery("start_date", "2023-01-01")
	endDate := c.DefaultQuery("end_date", "2023-12-31")

	// Panggil service untuk menghasilkan laporan
	report, err := services.GenerateSalesReport(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"report": report})
}
