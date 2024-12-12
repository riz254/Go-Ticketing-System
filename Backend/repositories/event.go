package repositories

import (
	"context"
	"log"

	"github.com/riz254/Go-Ticketing-System.git/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	var events []*models.Event

	// Find events and order by updated_at field
	res := r.db.Order("updated_at desc").Find(&events)

	if res.Error != nil {
		return nil, res.Error
	}

	return events, nil
}

func (r *EventRepository) GetOne(ctx context.Context, eventId uint) (*models.Event, error) {
	var event models.Event

	// Find the event by ID
	res := r.db.Where("id = ?", eventId).First(&event)

	if res.Error != nil {
		return nil, res.Error
	}

	return &event, nil
}

func (r *EventRepository) CreateOne(ctx context.Context, event *models.Event) (*models.Event, error) {
	// Create the event, auto-generating the ID
	res := r.db.Create(event)

	if res.Error != nil {
		return nil, res.Error
	}

	// Log the raw SQL that was executed
	log.Printf("Created event with query: %s", res.Statement.SQL.String())

	return event, nil
}

func (r *EventRepository) UpdateOne(ctx context.Context, eventId uint, updateData map[string]interface{}) (*models.Event, error) {
	event := &models.Event{}

	// Update the event with the provided ID
	updateRes := r.db.Model(event).Where("id = ?", eventId).Updates(updateData)

	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	// Fetch the updated event
	getRes := r.db.Where("id = ?", eventId).First(&event)

	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return event, nil
}

func (r *EventRepository) DeleteOne(ctx context.Context, eventId uint) error {
	// Delete the event by ID
	res := r.db.Delete(&models.Event{}, eventId)
	return res.Error
}

// NewEventRepository returns a new instance of EventRepository
func NewEventRepository(db *gorm.DB) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
