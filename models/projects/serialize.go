package projects

import (
	"github.com/pinem/server/models"
)

func GetSimpleProjects(projects []models.Project) []models.SimpleProject {
	simpleProjects := []models.SimpleProject{}
	for _, project := range projects {
		simpleProjects = append(simpleProjects, GetSimpleProject(project))
	}
	return simpleProjects
}

func GetSimpleProject(project models.Project) models.SimpleProject {
	return models.SimpleProject{Project: &project}
}
