package constructor

import (
	"github.com/hudayberdipolat/golang-api-crud/internal/app"
	postConstructor "github.com/hudayberdipolat/golang-api-crud/internal/domain/post/constructor"
)

func Build(dependencies app.AppDependencies) {
	postConstructor.PostRequirementsCreator(dependencies.DB, *dependencies.AppConfig)
}
