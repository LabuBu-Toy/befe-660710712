package week5_assignment

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Table struct {
	ID    string `json:"id"`
	Total int    `json:"total"`
}

type Component struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Stock    int    `json:"stock"`
	Category string `json:"category"`
}

var TableInShop = []Table{
	{ID: "1", Total: 4},
	{ID: "2", Total: 2},
	{ID: "3", Total: 1},
	{ID: "4", Total: 5},
	{ID: "5", Total: 3},
}

var ComponentsInStock = []Component{
	{ID: "1", Name: "Ryzen 5 5600X", Stock: 12, Category: "CPU"},
	{ID: "2", Name: "Core i5-12400F", Stock: 18, Category: "CPU"},
	{ID: "3", Name: "GeForce RTX 4060", Stock: 9, Category: "GPU"},
	{ID: "4", Name: "Radeon RX 7600", Stock: 7, Category: "GPU"},
	{ID: "5", Name: "16GB DDR4 3200", Stock: 40, Category: "RAM"},
	{ID: "6", Name: "32GB DDR5 5600", Stock: 25, Category: "RAM"},
	{ID: "7", Name: "1TB NVMe SSD", Stock: 33, Category: "Storage"},
	{ID: "8", Name: "2TB SATA HDD", Stock: 20, Category: "Storage"},
	{ID: "9", Name: "B550 ATX Motherboard", Stock: 14, Category: "Motherboard"},
	{ID: "10", Name: "Z690 mATX Motherboard", Stock: 10, Category: "Motherboard"},
	{ID: "11", Name: "650W 80+ Gold PSU", Stock: 22, Category: "Power"},
	{ID: "12", Name: "ATX Mid Tower Case", Stock: 16, Category: "Case"},
}

func getTable(c *gin.Context) {
	tableID := c.Query("ID")
	if tableID != "" {
		var found *Table
		for _, table := range TableInShop {
			if table.ID == tableID {
				t := table
				found = &t
				break
			}
		}
		if found == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "table not found"})
			return
		}
		size := "Big"
		if found.Total > 0 && found.Total < 3 {
			size = "Small"
		}
		c.JSON(http.StatusOK, gin.H{
			"table": found,
			"size":  size,
		})
		return
	}
	c.JSON(http.StatusOK, TableInShop)
}

func getComponents(c *gin.Context) {
	componentID := c.Query("ID")
	if componentID != "" {
		for _, comp := range ComponentsInStock {
			if comp.ID == componentID {
				c.JSON(http.StatusOK, comp)
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"error": "component not found"})
		return
	}
	c.JSON(http.StatusOK, ComponentsInStock)
}

func main() {
	r := gin.Default()

	// Single payload (rather than writing multiple JSONs)
	r.GET("/Soup", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"soups": []string{"Mala", "Shabu", "Tomyum"}})
	})

	api := r.Group("/api/v1")
	{
		api.GET("/MalaTable", getTable)
		// Replaced ingredient endpoint with computer components
		api.GET("/PCComponents", getComponents)
	}

	r.Run(":8080")
}
