package repositories

import (
	"context"

	"github.com/riz254/Go-Ticketing-System.git/models"
	"gorm.io/gorm"
)

type TicketRepository struct {
	db *gorm.DB
}

func (r *TicketRepository) GetMany(ctx context.Context) ([]*models.Ticket, error) {
	tickets := []*models.Ticket{}

	res := r.db.Model(&models.Ticket{}).Preload("Event").Order("updated_at desc").Find(&tickets)

	if res.Error != nil {
		return nil, res.Error
	}

	return tickets, nil
}

func (r *TicketRepository) GetOne(ctx context.Context, ticketId uint) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	res := r.db.Model(ticket).Preload("Event").Where("id = ?", ticketId).First(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return ticket, nil
}

func (r *TicketRepository) CreateOne(ctx context.Context, ticket *models.Ticket) (*models.Ticket, error) {
	res := r.db.Model(ticket).Create(ticket)

	if res.Error != nil {
		return nil, res.Error
	}

	return r.GetOne(ctx, ticket.ID)
}

func (r *TicketRepository) UpdateOne(ctx context.Context, ticketId uint, updateData map[string]interface{}) (*models.Ticket, error) {
	ticket := &models.Ticket{}

	// Perform the update operation
	updateRes := r.db.Model(ticket).Where("id = ?", ticketId).Updates(updateData)
	if updateRes.Error != nil {
		return nil, updateRes.Error
	}

	// Fetch the updated ticket using ticketId instead of ticket.ID
	updatedTicket := &models.Ticket{}
	getRes := r.db.Where("id = ?", ticketId).First(updatedTicket)
	if getRes.Error != nil {
		return nil, getRes.Error
	}

	return updatedTicket, nil
}

func NewTicketRepository(db *gorm.DB) models.TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
