package application

import (
	"bytes"
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
	"net/http"
)

const discordWebhookURL = "https://discordapp.com/api/webhooks/1343777777291231292/zphpoJrtodMhS_jE9aDvFe3NOBnfUqQK96na-a4XrRxXLfFUu0xdKC4vu8jcfBtkz2Hp"

func ProcessPullRequest(payload []byte) int {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(payload, &eventPayload); err != nil {
		return 500
	}

	if eventPayload.Action == "closed" {
		base := eventPayload.PullRequest.Base.Ref
		branch := eventPayload.PullRequest.Head.Ref
		user := eventPayload.PullRequest.User.Login
		pRID := eventPayload.PullRequest.ID
		title := eventPayload.PullRequest.Title
		repo := eventPayload.Repository.FullName

		log.Printf("Pull Request Recibido:\nID:%d\nBase:%s\nHead:%s\nUser:%s", pRID, base, branch, user)

		// Enviar notificación a Discord
		message := formatDiscordMessage(repo, user, base, branch, title, pRID)
		sendDiscordNotification(message)
	} else {
		log.Printf("Pull Request Action no es Closed: %s", eventPayload.Action)
	}

	return 200
}

// Formatea el mensaje para enviarlo a Discord
func formatDiscordMessage(repo, user, base, branch, title string, pRID int) string {
	return "📢 **Pull Request Cerrado**\n" +
		"📌 **Repositorio:** " + repo + "\n" +
		"👤 **Autor:** " + user + "\n" +
		"🔀 **Base:** " + base + "\n" +
		"🌿 **Head:** " + branch + "\n" +
		"📝 **Título:** " + title + "\n" +
		"🔗 **PR ID:** " + string(pRID)
}

// Envía un mensaje al Webhook de Discord
func sendDiscordNotification(message string) {
	payload := map[string]string{"content": message}
	jsonPayload, _ := json.Marshal(payload)

	resp, err := http.Post(discordWebhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		log.Printf("❌ Error enviando mensaje a Discord: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Println("✅ Notificación enviada a Discord correctamente")
}
