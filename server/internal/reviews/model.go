package reviews

import "time"

type Review struct {
	ID        int       `json:"id"`
	UserID    string    `json:"user_id"`
	PlantID   int       `json:"plant_id"`
	Rating    int       `json:"rating"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

type ReviewWithUser struct {
	Review
	UserName string `json:"user_name"`
}

type ReviewWithDetails struct {
	Review
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	PlantName string `json:"plant_name"`
	PlantSlug string `json:"plant_slug"`
}

type CreateReviewRequest struct {
	PlantID int    `json:"plant_id" binding:"required"`
	Rating  int    `json:"rating"   binding:"required,min=1,max=5"`
	Text    string `json:"text"     binding:"required,min=10"`
}
