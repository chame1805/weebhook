package application

import (
	"encoding/json"
	domain "github_wb/domain/value_objects"
	"log"
)

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

		log.Printf("Pull Request Recibido:\nID:%d\nBase:%s\nHead:%s\nUser:%s", pRID, base, branch, user)
	} else {
		log.Printf("Pull Request Action no es Closed: %s", eventPayload.Action)
	}

	/*log.Printf(
	"Evento Pull Request recibido: \nAcción=%s, \nPR Título='%s', \nRama Base='%s', \nRepositorio='%s'",
	eventPayload.Action, eventPayload.PullRequest.Title, eventPayload.PullRequest.Base.Ref, eventPayload.Repository.FullName)]*/

	/*mainBranch := "develop"

	if eventPayload.PullRequest.Base.Ref == mainBranch {
		log.Printf("¡Pull Request a la rama '%s' detectado en el repositorio '%s'!", mainBranch, eventPayload.Repository.FullName)
		fmt.Printf("Pull Request detectado en la rama %s!\n", mainBranch)
	} else {
		log.Printf(
			"Pull Request detectado, pero no dirigido a la rama '%s'. Rama base: '%s'",
			mainBranch, eventPayload.PullRequest.Base.Ref)
	}*/

	return 200
}
