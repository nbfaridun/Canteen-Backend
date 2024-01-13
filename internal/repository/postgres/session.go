package postgres

import (
	"Canteen-Backend/internal/constants"
	"Canteen-Backend/internal/models"
	"gorm.io/gorm"
)

type SessionPostgres struct {
	db *gorm.DB
}

func NewSessionPostgres(db *gorm.DB) *SessionPostgres {
	return &SessionPostgres{db: db}
}

func (r *SessionPostgres) CreateSession(session *models.Session) error {
	result := r.db.Table(constants.SessionTableName).Create(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SessionPostgres) DeleteSession(session *models.Session) error {
	result := r.db.Table(constants.SessionTableName).Delete(&models.Session{}, "session_id = ?", session.ID)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *SessionPostgres) GetSessionByRefreshToken(refreshToken string) (*models.Session, error) {
	var session models.Session
	result := r.db.Table(constants.SessionTableName).First(&session, "refresh_token = ?", refreshToken)
	if result.Error != nil {
		return nil, result.Error
	}

	return &session, nil
}
