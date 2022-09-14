package mock

import "github.com/Budi721/alterra-agmc/v2/models"

// Books mocking static data for endpoint book
var Books = []models.Book{
	{ID: uint(1), Title: "Anak Singkong", Author: "Chairil Tanjung", Price: uint(50000)},
	{ID: uint(2), Title: "Garis Waktu", Author: "Fiersa Besari", Price: uint(35000)},
}
