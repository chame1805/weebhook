package handlers

import (
	"github_wb/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryID := ctx.GetHeader("X-GitHub-Delivery")

	log.Printf("📩 Webhook recibido: Evento=%s, DeliveryID=%s", eventType, deliveryID)

	payload, err := ctx.GetRawData()
	if err != nil {
		log.Printf("❌ Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al leer el cuerpo de la solicitud"})
		return
	}

	log.Println("✅ Payload recibido correctamente. Enviando a ProcessPullRequest...")

	var statusCode int
	switch eventType {
	case "pull_request":
		statusCode = application.ProcessPullRequest(payload)
	default:
		log.Printf("⚠️ Evento ignorado: %s", eventType)
	}

	log.Printf("📢 Finalizando request con código: %d", statusCode)

	ctx.JSON(http.StatusOK, gin.H{"status": "Evento procesado"})
}
