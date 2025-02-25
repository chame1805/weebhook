package application

import (
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
)

func ProcessPullRequest(payload []byte) int {
	log.Println("📩 Recibiendo evento de Pull Request...")

	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		log.Printf("❌ Error al deserializar el payload: %v", err)
		return 500
	}

	log.Printf("✅ Evento recibido: Acción = %s", eventPayload.Action)

	if eventPayload.Action == "closed" {
		base := eventPayload.PullRequest.Base.Ref
		branch := eventPayload.PullRequest.Head.Ref
		user := eventPayload.PullRequest.User.Login
		pRID := eventPayload.PullRequest.ID

		log.Printf("🎉 Pull Request Cerrado:\n🔗 ID: %d\n📌 Base: %s\n🌿 Head: %s\n👤 Usuario: %s",
			pRID, base, branch, user)
	} else {
		log.Printf("⚠️ El evento recibido NO es 'closed', es: %s", eventPayload.Action)
	}

	return 200
}
