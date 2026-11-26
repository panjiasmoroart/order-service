package dto

import (
	"order-service/constants"
	"time"

	"github.com/google/uuid"
)

type PaymentData struct {
	OrderID   uuid.UUID                     `json:"orderID"`
	PaymentID uuid.UUID                     `json:"paymentID"`
	Status    constants.PaymentStatusString `json:"status"`
	ExpiredAt *time.Time                    `json:"expiredAt"`
	PaidAt    *time.Time                    `json:"paidAt"`
}

type PaymentContent struct {
	Event    KafkaEvent             `json:"event"`
	Metadata KafkaMetaData          `json:"metadata"`
	Body     KafkaBody[PaymentData] `json:"body"`
}

// sample response
// {
//   "event": {
//     "name": "PAYMENT_UPDATED",
//     "version": "1.0"
//   },
//   "metadata": {
//     "service": "payment-service",
//     "timestamp": "2026-11-26T13:00:00Z"
//   },
//   "body": {
//     "type": "PAYMENT",
//     "data": {
//       "orderID": "a1b2c3d4-1111-4444-8888-123456789abc",
//       "paymentID": "b2c3d4e5-2222-4444-8888-123456789abc",
//       "status": "SETTLEMENT",
//       "expiredAt": "2026-11-26T14:00:00Z",
//       "paidAt": "2026-11-26T13:05:00Z"
//     }
//   }
// }
